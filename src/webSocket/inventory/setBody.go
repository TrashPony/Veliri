package inventory

import (
	"../../mechanics/squadInventory"
	"../storage"
	"github.com/gorilla/websocket"
)

func SetBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	var err error

	if msg.Event == "SetMotherShipBody" {
		err = squadInventory.SetMSBody(user, msg.BodyID, msg.InventorySlot, msg.Source)
	}

	if msg.Event == "SetUnitBody" {
		err = squadInventory.SetUnitBody(user, msg.BodyID, msg.InventorySlot, msg.UnitSlot, msg.Source)
	}

	if err != nil {
		ws.WriteJSON(Response{Event: msg.Event, Error: err.Error()})
	} else {
		ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
		storage.Updater(user.GetID())
	}
}
