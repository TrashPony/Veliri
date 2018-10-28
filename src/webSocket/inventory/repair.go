package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
)

func Repair(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if msg.Event == "InventoryRepair" {
		err := inventory.ItemsRepair(user)
		if err != nil {
			ws.WriteJSON(Response{Event: "repair error", Error: err.Error()})
		} else {
			ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
		}
	}

	if msg.Event == "EquipsRepair" {
		err := inventory.EquipRepair(user)
		if err != nil {
			ws.WriteJSON(Response{Event: "repair error", Error: err.Error()})
		} else {
			ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
		}
	}

	if msg.Event == "AllRepair" {
		err := inventory.AllRepair(user)
		if err != nil {
			ws.WriteJSON(Response{Event: "repair error", Error: err.Error(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
		} else {
			ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
		}
	}
}
