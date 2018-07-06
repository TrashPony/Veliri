package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
)

func SetMotherShipBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.SetBody(user, msg.BodyID, msg.InventorySlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
