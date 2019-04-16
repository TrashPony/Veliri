package lobby

import (
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/TrashPony/Veliri/src/webSocket/utils"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersLobbyWs = make(map[*websocket.Conn]*player.Player) // тут будут храниться наши подключения
var lobbyPipe = make(chan Message)

func AddNewUser(ws *websocket.Conn, login string, id int) {

	utils.CheckDoubleLogin(login, &usersLobbyWs)

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	if newPlayer.InBaseID == 0 {
		// если игрок находиться не на базе то говорим ему что он загружал глобальную игру
		ws.WriteJSON(Message{Event: "OutBase"})
		return
	} else {
		newPlayer.LastBaseID = newPlayer.InBaseID

		if newPlayer.Fraction == "" {
			lobbyPipe <- Message{Event: "choiceFraction", UserID: newPlayer.GetID()}
		} else {
			// убираем у него скорость)
			if newPlayer.GetSquad() != nil {
				newPlayer.GetSquad().GlobalX = 0
				newPlayer.GetSquad().GlobalY = 0
			}

			if newPlayer.Training == 0 {
				// если игрок не прогшел обучение то кидаем ему первую страницу диалога введения
				trainingDialog := gameTypes.Dialogs.GetByID(1)
				lobbyPipe <- Message{Event: "dialog", UserID: newPlayer.GetID(), DialogPage: trainingDialog.Pages[1]}
				newPlayer.SetOpenDialog(&trainingDialog)
			} else {
				if newPlayer.Training < 999 {
					lobbyPipe <- Message{Event: "training", UserID: newPlayer.GetID(), Count: newPlayer.Training}
				}
			}
		}
	}

	usersLobbyWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS lobby Сессия: ") // просто смотрим новое подключение
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	// TODO проверять при крафте и произвосдстве на то что игрок находится на нужной базе
	var recycleItems map[int]*lobby.RecycleItem

	for {
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			utils.DelConn(ws, &usersLobbyWs, err)
			break
		}

		user := usersLobbyWs[ws]

		if msg.Event == "choiceFraction" {
			if msg.Fraction == "Replicas" || msg.Fraction == "Explores" || msg.Fraction == "Reverses" {
				user.Fraction = msg.Fraction
				// TODO у каждый фракции своя база для респауна
				//			выбераем базу, обновляем юзера, выставляем LastBaseID, запускаем туториал
				dbPlayer.UpdateUser(user)
			}
		}

		if user != nil && user.Fraction == "Replicas" || user.Fraction == "Explores" || user.Fraction == "Reverses" {

			if msg.Event == "Logout" {
				ws.Close()
			}

			if msg.Event == "OutBase" {
				outBase(user, msg)
			}

			if msg.Event == "PlaceItemsToProcessor" || msg.Event == "PlaceItemToProcessor" {
				placeItemToProcessor(user, msg, &recycleItems)
			}

			if msg.Event == "RemoveItemFromProcessor" || msg.Event == "RemoveItemsFromProcessor" {
				removeItemToProcessor(user, msg, &recycleItems)
			}

			if msg.Event == "ClearProcessor" {
				recycleItems = nil
			}

			if msg.Event == "recycle" {
				recycle(user, msg, &recycleItems)
			}

			if msg.Event == "OpenWorkbench" {
				openWorkbench(user, msg)
			}

			if msg.Event == "SelectBP" {
				selectBP(user, msg)
			}

			if msg.Event == "Craft" {
				craft(user, msg)
			}

			if msg.Event == "SelectWork" {
				selectWork(user, msg)
			}

			if msg.Event == "CancelCraft" {
				cancelCraft(user, msg)
			}

			if msg.Event == "OpenDialog" {

			}

			if msg.Event == "Ask" {

				page, err, action := dialog.Ask(user, user.GetOpenDialog(), "base", msg.ToPage, msg.AskID)

				if usersLobbyWs[ws].InBaseID > 0 && err == nil {
					lobbyPipe <- Message{Event: "dialog", UserID: user.GetID(), DialogPage: page, DialogAction: action}
				} else {
					lobbyPipe <- Message{Event: "Error", UserID: user.GetID(), Error: err.Error()}
				}
			}

			if msg.Event == "training" {
				user.Training = msg.Count
				dbPlayer.UpdateUser(user)
			}
		}
	}
}

func ReposeSender() {
	for {
		resp := <-lobbyPipe
		mutex.Lock()
		for ws, client := range usersLobbyWs {
			if client.GetID() == resp.UserID {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Fatal("lobby sender " + err.Error())
					utils.DelConn(ws, &usersLobbyWs, err)
				}
			}
		}
		mutex.Unlock()
	}
}
