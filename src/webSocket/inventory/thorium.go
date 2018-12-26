package inventory

import (
	"../../mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func setThorium(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squadInventory.SetThorium(user, msg.InventorySlot, msg.ThoriumSlot)
	if err != nil {
		// TODO
	}
	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}

func removeThoriumThorium(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squadInventory.RemoveThorium(user, msg.ThoriumSlot)
	if err != nil {
		// TODO
	}
	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}
