package inventory

import (
	"../../mechanics/factories/players"
	"../../mechanics/player"
	"../utils"
	"github.com/gorilla/websocket"
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

	print("WS inventory Сессия: ") // просто смотрим новое подключение
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

		if msg.Event == "changeSquad" {
			changeSquad(ws, msg)
		}

		if msg.Event == "SetMotherShipBody" || msg.Event == "SetUnitBody" {
			SetBody(ws, msg)
		}

		if msg.Event == "SetMotherShipWeapon" || msg.Event == "SetUnitWeapon" {
			SetWeapon(ws, msg)
		}

		if msg.Event == "SetMotherShipEquip" || msg.Event == "SetUnitEquip" {
			SetEquip(ws, msg)
		}

		if msg.Event == "SetMotherShipAmmo" || msg.Event == "SetUnitAmmo" {
			SetAmmo(ws, msg)
		}

		if msg.Event == "RemoveMotherShipBody" || msg.Event == "RemoveUnitBody" {
			RemoveBody(ws, msg)
		}

		if msg.Event == "RemoveMotherShipAmmo" || msg.Event == "RemoveUnitAmmo" {
			RemoveAmmo(ws, msg)
		}

		if msg.Event == "RemoveMotherShipWeapon" || msg.Event == "RemoveUnitWeapon" {
			RemoveWeapon(ws, msg)
		}

		if msg.Event == "RemoveMotherShipEquip" || msg.Event == "RemoveUnitEquip" {
			RemoveEquip(ws, msg)
		}

		if msg.Event == "itemToStorage" || msg.Event == "itemsToStorage" { // из инвентаря в хранилище
			itemToStorage(ws, msg)
		}

		if msg.Event == "itemToInventory" || msg.Event == "itemsToInventory" { // из хранилища в инвентарь
			itemToInventory(ws, msg)
		}

		if msg.Event == "SetThorium" {
			setThorium(ws, msg)
		}

		if msg.Event == "RemoveThorium" {
			removeThoriumThorium(ws, msg)
		}

		if msg.Event == "InventoryRepair" || msg.Event == "EquipsRepair" || msg.Event == "AllRepair" {
			Repair(ws, msg)
		}
	}
}
