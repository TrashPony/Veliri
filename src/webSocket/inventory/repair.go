package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func Repair(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	if msg.Event == "InventoryRepair" {
		err := squadInventory.ItemsRepair(user)
		if err != nil {
			ws.WriteJSON(Response{Event: "repair error", Error: err.Error()})
		} else {
			ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		}
	}

	if msg.Event == "EquipsRepair" {
		err := squadInventory.EquipRepair(user)
		if err != nil {
			ws.WriteJSON(Response{Event: "repair error", Error: err.Error()})
		} else {
			ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		}
	}

	if msg.Event == "AllRepair" {
		err := squadInventory.AllRepair(user)
		if err != nil {
			ws.WriteJSON(Response{Event: "repair error", Error: err.Error(), InventorySize: user.GetSquad().Inventory.GetSize()})
		} else {
			ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		}
	}
}
