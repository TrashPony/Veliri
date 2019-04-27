package storage

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/webSocket/utils"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersStorageWs = make(map[*websocket.Conn]*player.Player)

type Message struct {
	Event string `json:"event"`
}

func AddNewUser(ws *websocket.Conn, login string, id int) {

	mutex.Lock()

	utils.CheckDoubleLogin(login, &usersStorageWs)

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	usersStorageWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS storage Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))

	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция завершается

	mutex.Unlock()

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	for {
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil {
			println(err.Error())
			utils.DelConn(ws, &usersStorageWs, err)
			break
		}

		if msg.Event == "viewAllStorage" {
			// todo выводить информацию о всех итемах на всех базах
		}

		if msg.Event == "openStorage" {
			userStorage, ok := storages.Storages.Get(usersStorageWs[ws].GetID(), usersStorageWs[ws].InBaseID)
			if ok {
				ws.WriteJSON(userStorage)
			}
		}
	}
}

func Updater(userID int) {

	// тут происходит обновление инвентаря склада когда произошло обновление, а у пользователя он открыт
	// например произошла покупка итема и он упал в склад, надо обновить информацию у пользователя

	mutex.Lock()
	for ws, user := range usersStorageWs {
		if user.GetID() == userID {
			userStorage, ok := storages.Storages.Get(usersStorageWs[ws].GetID(), usersStorageWs[ws].InBaseID)
			if ok {
				err := ws.WriteJSON(userStorage)
				if err != nil {
					log.Fatal("Storage" + err.Error())
				}
			}
		}
	}
	mutex.Unlock()
}
