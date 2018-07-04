package inventory

import ("github.com/gorilla/websocket"
"../../mechanics/inventory")

func SetMotherShipWeapon(ws *websocket.Conn, msg Message)  {
	user := usersInventoryWs[ws]

	inventory.SetWeapon(user, msg.BodyID, msg.InventorySlot, msg.EquipSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
