package lobby

import (
	"../../mechanics/factories/players"
	"../../mechanics/player"
	"../utils"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
	"../../mechanics/lobby"
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
	for {
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			utils.DelConn(ws, &usersLobbyWs, err)
			break
		}

		if msg.Event == "Logout" {
			ws.Close()
		}

		if msg.Event == "OutBase" {
			err := lobby.OutBase(usersLobbyWs[ws])

			// todo запускать метод в отдельной горутине
			// todo флаг выхода с базы, т.к. пока освобождается респаун игрок может передумать

			if err != nil {
				ws.WriteJSON(Message{Event: "Error", Error: err.Error()})
			} else {
				ws.WriteJSON(Message{Event: msg.Event})
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
					log.Fatal(err)
					utils.DelConn(ws, &usersLobbyWs, err)
				}
			}
		}
		mutex.Unlock()
	}
}
