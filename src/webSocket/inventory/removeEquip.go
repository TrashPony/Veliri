package inventory

import (
	"../../mechanics/squadInventory"
	"../storage"
	"github.com/gorilla/websocket"
)

func RemoveEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	var err error

	if msg.Event == "RemoveMotherShipEquip" {
		err = squadInventory.RemoveEquip(user, msg.EquipSlot, msg.EquipSlotType, user.GetSquad().MatherShip, msg.Destination)
	}

	if msg.Event == "RemoveUnitEquip" {
		err = squadInventory.RemoveUnitEquip(user, msg.EquipSlot, msg.EquipSlotType, msg.UnitSlot, msg.Destination)
	}

	if err != nil {
		ws.WriteJSON(Response{Event: msg.Event, Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		storage.Updater(user.GetID())
	}
}
