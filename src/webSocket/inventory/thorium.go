package inventory

import (
	"../../mechanics/squadInventory"
	"../storage"
	"github.com/gorilla/websocket"
)

func setThorium(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squadInventory.SetThorium(user, msg.InventorySlot, msg.ThoriumSlot, msg.Source)
	if err != nil {
		ws.WriteJSON(Response{Event: msg.Event, Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		storage.Updater(user.GetID())
	}
}

func removeThoriumThorium(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squadInventory.RemoveThorium(user, msg.ThoriumSlot, true)
	if err != nil {
		ws.WriteJSON(Response{Event: msg.Event, Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		storage.Updater(user.GetID())
	}
}
