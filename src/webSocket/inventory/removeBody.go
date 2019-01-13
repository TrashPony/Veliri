package inventory

import (
	"../../mechanics/squadInventory"
	"../storage"
	"github.com/gorilla/websocket"
)

func RemoveBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	var err error

	if msg.Event == "RemoveMotherShipBody" {
		squadInventory.RemoveMSBody(user)
	}

	if msg.Event == "RemoveUnitBody" {
		squadInventory.RemoveUnitBody(user, msg.UnitSlot, true)
	}

	if err != nil {
		ws.WriteJSON(Response{Event: msg.Event, Error: err.Error()})
	} else {
		if user.GetSquad() != nil {
			ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		} else {
			ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: 0})
		}
		storage.Updater(user.GetID())
	}
}
