package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/squadInventory"
)

func setThorium(ws *websocket.Conn, msg Message)  {
	user := usersInventoryWs[ws]

	err := squadInventory.SetThorium(user, msg.InventorySlot, msg.ThoriumSlot)
	if err != nil {
		// TODO
	}
	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}
