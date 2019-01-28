package lobby

import (
	"../../mechanics/factories/players"
	"../../mechanics/gameObjects/inventory"
	"../../mechanics/player"
	"../utils"
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

	if newPlayer.InBaseID == 0 { // если игрок находиться не на базе то говорим ему что он загружал глобальную игру
		ws.WriteJSON(Message{Event: "OutBase"})
		return
	} else { // иначе убираем у него скорость)
		newPlayer.GetSquad().GlobalX = 0
		newPlayer.GetSquad().GlobalY = 0
	}

	usersLobbyWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS lobby Сессия: ") // просто смотрим новое подключение
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	var recycleItems map[int]*inventory.Slot

	for {
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil {          // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			utils.DelConn(ws, &usersLobbyWs, err)
			break
		}

		if msg.Event == "Logout" {
			ws.Close()
		}

		if msg.Event == "OutBase" {
			outBase(ws, msg)
		}

		if msg.Event == "PlaceItemsToProcessor" || msg.Event == "PlaceItemToProcessor" {
			placeItemToProcessor(ws, msg, &recycleItems)
		}

		if msg.Event == "RemoveItemFromProcessor" {
			removeItemToProcessor(ws, msg, &recycleItems)
		}

		if msg.Event == "ClearProcessor" {
			recycleItems = nil
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
					log.Fatal(err)
					utils.DelConn(ws, &usersLobbyWs, err)
				}
			}
		}
		mutex.Unlock()
	}
}
