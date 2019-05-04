package inventory

import (
	"github.com/gorilla/websocket"
)

func changeColor(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil || user.GetSquad().MatherShip == nil || user.GetSquad().MatherShip.Body == nil {
		return
	}

	if msg.UnitSlot == 0 {
		user.GetSquad().MatherShip.BodyColor1 = msg.BodyColor1
		user.GetSquad().MatherShip.BodyColor2 = msg.BodyColor2
		user.GetSquad().MatherShip.WeaponColor1 = msg.WeaponColor1
		user.GetSquad().MatherShip.WeaponColor2 = msg.WeaponColor2
	} else {
		unitSlot, ok := user.GetSquad().MatherShip.Units[msg.UnitSlot]
		if ok && unitSlot != nil && unitSlot.Unit != nil {
			unitSlot.Unit.BodyColor1 = msg.BodyColor1
			unitSlot.Unit.BodyColor2 = msg.BodyColor2
			unitSlot.Unit.WeaponColor1 = msg.WeaponColor1
			unitSlot.Unit.WeaponColor2 = msg.WeaponColor2
		}
	}

	UpdateSquad("UpdateSquad", user, nil, ws, msg)
}
