package inventory

import (
	"github.com/gorilla/websocket"
	"../../inventory"
)


var usersInventoryWs = make(map[*websocket.Conn]*inventory.User)

func AddNewUser(ws *websocket.Conn, login string, id int) {
	CheckDoubleLogin(login, &usersInventoryWs)
	usersInventoryWs[ws] = &inventory.User{Name: login, Id: id} // Регистрируем нового Клиента
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция завершается

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	for {
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message

		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersInventoryWs, err)
			break
		}

		if msg.Event == "openInventory" {

		}

		if msg.Event == "AddNewSquad" || msg.Event == "SelectSquad" || msg.Event == "SelectMatherShip" || msg.Event == "DeleteSquad" {
			SquadSettings(ws, msg)
		}

		if msg.Event == "GetMatherShips" || msg.Event == "GetListSquad" || msg.Event == "GetDetailOfUnits" || msg.Event == "GetEquipping" {
			GetDetailSquad(ws, msg)
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
