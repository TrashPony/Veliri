package inventory

import (
	"../../mechanics/squadInventory"
	"../storage"
	"github.com/gorilla/websocket"
)

func itemToStorage(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if msg.Event == "itemToStorage" {
		err := squadInventory.ItemToStorage(user, msg.InventorySlot)
		if err != nil {
			ws.WriteJSON(Response{Event: "Error", Error: err.Error()})
		} else {
			storage.Updater(user.GetID())
			ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		}
	}

	if msg.Event == "itemsToStorage" {
		for _, i := range msg.InventorySlots {
			err := squadInventory.ItemToStorage(user, i)
			if err != nil {
				ws.WriteJSON(Response{Event: "Error", Error: err.Error()})
			} else {
				storage.Updater(user.GetID())
				//concurrent map iteration and map write
				ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
			}
		}
	}
}
