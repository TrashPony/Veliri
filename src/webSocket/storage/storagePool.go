package storage

import (
	"../../mechanics/player"
	"../../mechanics/players"
	"../../mechanics/storage"
	"../utils"
	"github.com/gorilla/websocket"
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
			ws.WriteJSON(storage.GetUserBaseStorage(usersStorageWs[ws]))
		}
	}
}

func Updater(userID int) {
	// TODO когда происходит обновление инвентаря склада, а у пользователя он открыт
	// TODO например произошла покупка итема и он упал в склад, надо обновить информацию у пользователя
	// TODO найти его в соеденения и отправить ему на сокет новый инвентарь ws.WriteJSON(storage.GetUserBaseStorage(usersStorageWs[ws]))
}
