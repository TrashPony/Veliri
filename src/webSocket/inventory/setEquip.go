package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
)

func SetMotherShipEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.SetMSEquip(user, msg.EquipID, msg.InventorySlot, msg.EquipSlot, msg.EquipSlotType)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

func SetUnitEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.SetUnitEquip(user, msg.EquipID, msg.InventorySlot, msg.EquipSlot, msg.EquipSlotType, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

