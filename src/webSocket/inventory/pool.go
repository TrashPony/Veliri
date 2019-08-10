package inventory

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/webSocket/utils"
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
			mutex.Lock()
			utils.DelConn(ws, &usersInventoryWs, err)
			mutex.Unlock()
			break
		}

		user := usersInventoryWs[ws]

		if user != nil {
			if msg.Event == "openInventory" {
				Open(ws, msg)
			}

			if msg.Event == "RenameSquad" {
				renameSquad(ws, msg)
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

			if msg.Event == "divideItems" {
				divideItems(ws, msg)
			}

			if msg.Event == "combineItems" {
				combineItems(ws, msg)
			}

			if msg.Event == "changeColor" {
				changeColor(ws, msg)
			}
		}
	}
}

func UpdateSquad(event string, user *player.Player, err error, ws *websocket.Conn, msg Message) {

	if user.GetSquad() != nil {
		go update.Squad(user.GetSquad(), true)
	}

	if err != nil {
		ws.WriteJSON(Response{Event: msg.Event, Error: err.Error(), UnitSlot: msg.UnitSlot})
	} else {
		UpdateInventory(user.GetID())
	}

	if user.InBaseID > 0 {
		UpdateStorage(user.GetID())
	}
}

func UpdateStorage(userID int) {
	mutex.Lock()
	defer mutex.Unlock()

	// тут происходит обновление инвентаря склада когда произошло обновление, а у пользователя он открыт
	// например произошла покупка итема и он упал в склад, надо обновить информацию у пользователя

	for ws, user := range usersInventoryWs {
		if user.GetID() == userID {
			userStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)
			err := ws.WriteJSON(Response{Event: "UpdateStorage", Storage: userStorage})
			if err != nil {
				ws.Close()
			}
		}
	}
}

func UpdateInventory(userID int) {
	mutex.Lock()
	defer mutex.Unlock()

	for ws, user := range usersInventoryWs {
		if user.GetID() == userID {
			if user.GetSquad() != nil {
				err := ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), BaseSquads: user.GetSquadsByBaseID(user.InBaseID),
					InventorySize: user.GetSquad().Inventory.GetSize(), InBase: user.InBaseID > 0})
				if err != nil {
					ws.Close()
				}
			} else {
				err := ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), BaseSquads: user.GetSquadsByBaseID(user.InBaseID),
					InventorySize: 0, InBase: user.InBaseID > 0})
				if err != nil {
					ws.Close()
				}
			}
		}
	}
}
