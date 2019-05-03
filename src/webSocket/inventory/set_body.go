package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func SetBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	var err error

	if msg.Event == "SetMotherShipBody" {
		// установить тело без отряда можно, и тогда создастся отряд
		err = squad_inventory.SetMSBody(user, msg.BodyID, msg.InventorySlot, msg.Source)
	}

	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}

	if msg.Event == "SetUnitBody" {
		err = squad_inventory.SetUnitBody(user, msg.BodyID, msg.InventorySlot, msg.UnitSlot, msg.Source)
	}

	UpdateSquad("UpdateSquad", user, err, ws, msg)
}
