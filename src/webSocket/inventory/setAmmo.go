package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
)

func SetMotherShipAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := inventory.SetMSAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot)
	if err != nil {
		ws.WriteJSON(Response{Event: "error", Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
	}
}

func SetUnitAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := inventory.SetUnitAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot)
	if err != nil {
		ws.WriteJSON(Response{Event: "error", Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
	}
}
