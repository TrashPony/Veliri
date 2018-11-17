package inventory

import (
	"../../mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func RemoveMotherShipBody(ws *websocket.Conn) {
	user := usersInventoryWs[ws]

	squadInventory.RemoveMSBody(user)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
}

func RemoveUnitBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	squadInventory.RemoveUnitBody(user, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
}
