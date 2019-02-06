package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func itemToStorage(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	if msg.Event == "itemToStorage" {
		err := squadInventory.ItemToStorage(user, msg.InventorySlot)
		UpdateSquad(user, err, ws, msg)
	}

	if msg.Event == "itemsToStorage" {
		for _, i := range msg.InventorySlots {
			err := squadInventory.ItemToStorage(user, i)
			UpdateSquad(user, err, ws, msg)
		}
	}
}
