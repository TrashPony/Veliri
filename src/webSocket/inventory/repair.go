package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func Repair(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error
	if msg.Event == "InventoryRepair" {
		err = squad_inventory.ItemsRepair(user)
		//if err != nil {
		//	ws.WriteJSON(Response{Event: "repair error", Error: err.Error()})
		//} else {
		//	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		//}
	}

	if msg.Event == "EquipsRepair" {
		err = squad_inventory.EquipRepair(user)
		//if err != nil {
		//	ws.WriteJSON(Response{Event: "repair error", Error: err.Error()})
		//} else {
		//	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		//}
	}

	if msg.Event == "AllRepair" {
		err = squad_inventory.AllRepair(user)
		//if err != nil {
		//	ws.WriteJSON(Response{Event: "repair error", Error: err.Error(), InventorySize: user.GetSquad().Inventory.GetSize()})
		//} else {
		//	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		//}
	}

	UpdateSquad("UpdateSquad", user, err, ws, msg)
}
