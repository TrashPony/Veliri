package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
)

func RemoveMotherShipEquip(ws *websocket.Conn, msg Message)  {
	user := usersInventoryWs[ws]

	inventory.RemoveMSEquip(user, msg.EquipSlot, msg.EquipSlotType)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

func RemoveUnitEquip(ws *websocket.Conn, msg Message)  {
	user := usersInventoryWs[ws]

	inventory.RemoveUnitEquip(user, msg.EquipSlot, msg.EquipSlotType, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

