package inventory

import (
	"../../mechanics/squadInventory"
	"errors"
	"github.com/gorilla/websocket"
)

func SetBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	var err error

	if msg.Event == "SetMotherShipBody" {
		// установить тело без отряда можно, и тогда создастся отряд
		err = squadInventory.SetMSBody(user, msg.BodyID, msg.InventorySlot, msg.Source)
	}

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	if msg.Event == "SetUnitBody" {
		err = squadInventory.SetUnitBody(user, msg.BodyID, msg.InventorySlot, msg.UnitSlot, msg.Source)
	}

	UpdateSquad(user, err, ws, msg)
}
