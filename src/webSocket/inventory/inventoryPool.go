package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/player"
	"../../mechanics/players"
	"../utils"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersInventoryWs = make(map[*websocket.Conn]*player.Player)

func AddNewUser(ws *websocket.Conn, login string, id int) {

	mutex.Lock()

	utils.CheckDoubleLogin(login, &usersInventoryWs)

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	usersInventoryWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS inventory Сессия: ")                          // просто смотрим новое подключение
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

		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			println(err.Error())
			utils.DelConn(ws, &usersInventoryWs, err)
			break
		}

		if msg.Event == "openInventory" {
			Open(ws, msg)
		}

		if msg.Event == "SetMotherShipBody" {
			SetMotherShipBody(ws, msg)
		}

		if msg.Event == "SetMotherShipWeapon" {
			SetMotherShipWeapon(ws, msg)
		}
	}
}
