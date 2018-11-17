package inventory

import (
	"../../mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func SetMotherShipAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squadInventory.SetMSAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot)
	if err != nil {
		ws.WriteJSON(Response{Event: "error", Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
	}
}

func SetUnitAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squadInventory.SetUnitAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot)
	if err != nil {
		ws.WriteJSON(Response{Event: "error", Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
	}
}
