package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
)

func RemoveMotherShipWeapon(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.RemoveMSWeapon(user, msg.EquipSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

func RemoveUnitWeapon(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.RemoveUnitWeapon(user, msg.EquipSlot, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
