package lobby

import (
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
	"github.com/TrashPony/Veliri/src/webSocket/utils"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersLobbyWs = make(map[*websocket.Conn]*player.Player) // тут будут храниться наши подключения
var lobbyPipe = make(chan Message)

func AddNewUser(ws *websocket.Conn, login string, id int) {

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	if newPlayer.InBaseID == 0 {
		// если игрок находиться не на базе то говорим ему что он загружал глобальную игру или бой

		if newPlayer.GetSquad().InGame {
			ws.WriteJSON(Message{Event: "LocalGame"})
		} else {
			ws.WriteJSON(Message{Event: "OutBase"})
		}

		return
	} else {

		usersLobbyWs[ws] = newPlayer // Регистрируем нового Клиента
		newPlayer.LastBaseID = newPlayer.InBaseID

		// отправляем текущие состояние базы
		BaseStatus(newPlayer)
		checkNoobs(newPlayer)

		// убираем скорость у игрока если у него есть отряд
		if newPlayer.GetSquad() != nil {
			newPlayer.GetSquad().MatherShip.X = 0
			newPlayer.GetSquad().MatherShip.Y = 0
		}
	}

	print("WS lobby Сессия: ") // просто смотрим новое подключение
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

	Reader(ws)
}

func checkNoobs(newPlayer *player.Player) {
	if newPlayer.Fraction == "" {
		//новый игрок без фракции должен сделать выбор
		lobbyPipe <- Message{Event: "choiceFraction", UserID: newPlayer.GetID()}
	} else {
		if newPlayer.Training == 0 {

			// если игрок не прошел обучение то кидаем ему первую страницу диалога введения
			trainingDialog := gameTypes.Dialogs.GetTrainingStartDialog(newPlayer.Fraction)
			if trainingDialog != nil {
				base, _ := bases.Bases.Get(newPlayer.InBaseID)
				trainingDialog.ProcessingDialogText(newPlayer.GetLogin(), base.Name, "", "", newPlayer.Fraction)
				lobbyPipe <- Message{Event: "dialog", UserID: newPlayer.GetID(), DialogPage: trainingDialog.Pages[1]}
				newPlayer.SetOpenDialog(trainingDialog)
			}
		} else {
			if newPlayer.Training < 999 {
				lobbyPipe <- Message{Event: "training", UserID: newPlayer.GetID(), Count: newPlayer.Training}
			}
		}
	}
}

func Reader(ws *websocket.Conn) {
	// TODO проверять при крафте и произвосдстве на то что игрок находится на нужной базе
	// source, slot_number, slot
	var recycleItems map[string]map[int]*lobby.RecycleItem

	for {
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil {          // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			println(err.Error())
			utils.DelConn(ws, &usersLobbyWs, err)
			break
		}

		user := usersLobbyWs[ws]

		if msg.Event == "choiceFraction" {
			if (msg.Fraction == "Replics" || msg.Fraction == "Explores" || msg.Fraction == "Reverses") && user.Fraction == "" {
				user.Fraction = msg.Fraction

				base := bases.Bases.GetCapital(user.Fraction)

				user.InBaseID = base.ID
				user.LastBaseID = base.ID

				dbPlayer.UpdateUser(user)
				lobbyPipe <- Message{Event: "choiceFractionComplete", UserID: user.GetID()}

				trainingDialog := gameTypes.Dialogs.GetByID(1)
				lobbyPipe <- Message{Event: "dialog", UserID: user.GetID(), DialogPage: trainingDialog.Pages[1]}
				user.SetOpenDialog(trainingDialog)
			}
		}

		if user != nil && user.InBaseID > 0 && (user.Fraction == "Replics" || user.Fraction == "Explores" || user.Fraction == "Reverses") {

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

			if msg.Event == "GetBaseStatus" {
				BaseStatus(user)
			}

			if msg.Event == "GetDetails" {
				GetDetails(user)
			}

			if msg.Event == "SellDetail" {
				SellDetail(user, msg)
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
					utils.DelConn(ws, &usersLobbyWs, err)
				}
			}
		}
		mutex.Unlock()
	}
}
