package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func itemToInventory(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}

	if msg.Event == "itemToInventory" {
		err := squad_inventory.ItemToInventory(user, msg.StorageSlot)
		UpdateSquad("UpdateSquad", user, err, ws, msg)
	}

	if msg.Event == "itemsToInventory" {
		for _, i := range msg.StorageSlots {
			squad_inventory.ItemToInventory(user, i)
		}
		UpdateSquad("UpdateSquad", user, nil, ws, msg)
	}
}
