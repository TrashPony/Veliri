package inventory

import (
	"../../mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func SetMotherShipBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	squadInventory.SetMSBody(user, msg.BodyID, msg.InventorySlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}

func SetUnitBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	squadInventory.SetUnitBody(user, msg.BodyID, msg.InventorySlot, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}
