package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
)

func RemoveMotherShipBody(ws *websocket.Conn) {
	user := usersInventoryWs[ws]

	inventory.RemoveMSBody(user)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

func RemoveUnitBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.RemoveUnitBody(user, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
