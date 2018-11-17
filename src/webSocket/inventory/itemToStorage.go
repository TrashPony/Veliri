package inventory

import (
	"../../mechanics/squadInventory"
	"../storage"
	"github.com/gorilla/websocket"
)

func itemToStorage(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squadInventory.ItemToStorage(user, msg.InventorySlot)

	if err != nil {
		// TODO
	} else {
		storage.Updater(user.GetID())
	}
}
