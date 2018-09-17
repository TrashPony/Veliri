package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
)

func SetMotherShipBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.SetMSBody(user, msg.BodyID, msg.InventorySlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

func SetUnitBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.SetUnitBody(user, msg.BodyID, msg.InventorySlot, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
