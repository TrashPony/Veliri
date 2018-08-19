package inventory

import ("github.com/gorilla/websocket"
"../../mechanics/inventory")

func SetMotherShipWeapon(ws *websocket.Conn, msg Message)  {

	user := usersInventoryWs[ws]

	inventory.SetMSWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

func SetUnitWeapon(ws *websocket.Conn, msg Message)  {
	user := usersInventoryWs[ws]

	inventory.SetUnitWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}