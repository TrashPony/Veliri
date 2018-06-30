package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/player"
	"../utils"
	"strconv"
)


var usersInventoryWs = make(map[*websocket.Conn]*player.Player)

func AddNewUser(ws *websocket.Conn, login string, id int) {
	utils.CheckDoubleLogin(login, &usersInventoryWs)

	newPlayer := &player.Player{}
	newPlayer.SetLogin(login)
	newPlayer.SetID(id)

	print("WS inventory Сессия: ")                          // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))

	usersInventoryWs[ws] = newPlayer // Регистрируем нового Клиента
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция завершается

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	for {
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message

		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			utils.DelConn(ws, &usersInventoryWs, err)
			break
		}

		if msg.Event == "openInventory" {
			Open(ws, msg)
		}

		if msg.Event == "AddNewSquad" || msg.Event == "SelectSquad" || msg.Event == "SelectMatherShip" || msg.Event == "DeleteSquad" {
			SquadSettings(ws, msg)
		}

		if msg.Event == "AddUnit" || msg.Event == "ReplaceUnit" || msg.Event == "RemoveUnit" {
			UnitSquad(ws, msg)
		}

		if msg.Event == "AddEquipment" || msg.Event == "ReplaceEquipment" || msg.Event == "RemoveEquipment" {
			EquipSquad(ws, msg)
		}

		if msg.Event == "UnitConstructor" {
			UnitConstructor(ws, msg)
		}
	}
}
