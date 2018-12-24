package inventory

import (
	"../../mechanics/squadInventory"
	"../storage"
	"github.com/gorilla/websocket"
)

func itemToInventory(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squadInventory.ItemToInventory(user, msg.StorageSlot)

	if err != nil {
		// TODO
	} else {
		storage.Updater(user.GetID())
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
	}
}
