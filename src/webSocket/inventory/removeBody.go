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
		squadInventory.RemoveUnitBody(user, msg.UnitSlot)
	}

	if err != nil {
		ws.WriteJSON(Response{Event: msg.Event, Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		storage.Updater(user.GetID())
	}
}
