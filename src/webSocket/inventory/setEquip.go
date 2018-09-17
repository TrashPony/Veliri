package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
)

func SetMotherShipEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := inventory.SetMSEquip(user, msg.EquipID, msg.InventorySlot, msg.EquipSlot, msg.EquipSlotType)

	if err != nil {
		ws.WriteJSON(Error{Event: "ms error", Error: err.Error()})
	}

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

func SetUnitEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := inventory.SetUnitEquip(user, msg.EquipID, msg.InventorySlot, msg.EquipSlot, msg.EquipSlotType, msg.UnitSlot)

	if err != nil {
		ws.WriteJSON(Error{Event: "unit error", Error: err.Error()})
	}

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
