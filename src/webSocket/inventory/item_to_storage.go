package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func itemToStorage(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}

	if msg.Event == "itemToStorage" {
		err := squad_inventory.ItemToStorage(user, msg.InventorySlot)
		UpdateSquad("UpdateSquad", user, err, ws, msg)
	}

	if msg.Event == "itemsToStorage" {
		for _, i := range msg.InventorySlots {
			squad_inventory.ItemToStorage(user, i)
		}
		UpdateSquad("UpdateSquad", user, nil, ws, msg)
	}
}
