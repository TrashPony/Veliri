package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
)

func SetMotherShipAmmo(ws *websocket.Conn, msg Message)  {
	user := usersInventoryWs[ws]

	inventory.SetMSAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

func SetUnitAmmo(ws *websocket.Conn, msg Message)  {
	user := usersInventoryWs[ws]

	inventory.SetUnitAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

