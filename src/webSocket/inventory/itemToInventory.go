package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func itemToInventory(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	if msg.Event == "itemToInventory" {
		err := squadInventory.ItemToInventory(user, msg.StorageSlot)
		UpdateSquad(user, err, ws, msg)
	}

	if msg.Event == "itemsToInventory" {
		for _, i := range msg.StorageSlots {
			err := squadInventory.ItemToInventory(user, i)
			UpdateSquad(user, err, ws, msg)
		}
	}
}
