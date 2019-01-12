package inventory

import (
	"../../mechanics/squadInventory"
	"../storage"
	"github.com/gorilla/websocket"
)

func RemoveAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	var err error

	if msg.Event == "RemoveMotherShipAmmo" {
		err = squadInventory.RemoveAmmo(user, msg.EquipSlot, user.GetSquad().MatherShip, msg.Destination)
	}

	if msg.Event == "RemoveUnitAmmo" {
		err = squadInventory.RemoveUnitAmmo(user, msg.EquipSlot, msg.UnitSlot, msg.Destination)
	}

	if err != nil {
		ws.WriteJSON(Response{Event: msg.Event, Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		storage.Updater(user.GetID())
	}
}
