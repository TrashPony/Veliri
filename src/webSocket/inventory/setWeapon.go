package inventory

import (
	"../../mechanics/squadInventory"
	"../storage"
	"github.com/gorilla/websocket"
)

func SetWeapon(ws *websocket.Conn, msg Message) {

	user := usersInventoryWs[ws]

	var err error

	if msg.Event == "SetMotherShipWeapon" {
		err = squadInventory.SetWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot, user.GetSquad().MatherShip, msg.Source)
	}

	if msg.Event == "SetUnitWeapon" {
		err = squadInventory.SetUnitWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot, msg.Source)
	}

	if err != nil {
		ws.WriteJSON(Response{Event: msg.Event, Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		storage.Updater(user.GetID())
	}
}
