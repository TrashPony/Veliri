package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
)

func RemoveMotherShipWeapon(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.RemoveWeapon(user, msg.EquipSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
