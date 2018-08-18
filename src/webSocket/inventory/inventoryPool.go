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

		if msg.Event == "SetMotherShipEquip" {
			SetMotherShipEquip(ws, msg)
		}

		if msg.Event == "SetMotherShipAmmo" {
			SetMotherShipAmmo(ws, msg)
		}

		if msg.Event == "RemoveMotherShipBody" {
			RemoveMotherShipBody(ws)
		}

		if msg.Event == "RemoveMotherShipAmmo" {
			RemoveMotherShipAmmo(ws, msg)
		}

		if msg.Event == "RemoveMotherShipWeapon" {
			RemoveMotherShipWeapon(ws, msg)
		}

		if msg.Event == "RemoveMotherShipEquip" {
			RemoveMotherShipEquip(ws, msg)
		}

		if msg.Event == "SetUnitBody" {
			SetUnitBody(ws, msg)
		}

		if msg.Event == "RemoveUnitBody" {
			RemoveUnitBody(ws, msg)
		}
	}
}
