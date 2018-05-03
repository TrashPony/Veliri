package lobby

import (
	"../../lobby/DetailUnit"
	"../../lobby/Squad"
	"github.com/gorilla/websocket"
)

func UnitConstructor(ws *websocket.Conn, msg Message) {

	unit, ok := usersLobbyWs[ws].Squad.Units[msg.UnitSlot]
	if !ok {
		unit = &Squad.Unit{}
	}

	if msg.WeaponID != 0 {
		unit.SetWeapon(DetailUnit.GetWeapon(msg.WeaponID))
	} else {
		unit.DelWeapon()
	}

	if msg.ChassisID != 0 {
		unit.SetChassis(DetailUnit.GetChass(msg.ChassisID))
	} else {
		unit.DelChassis()
	}

	if msg.TowerID != 0 {
		unit.SetTower(DetailUnit.GetTower(msg.TowerID))
	} else {
		unit.DelTower()
	}

	if msg.BodyID != 0 {
		unit.SetBody(DetailUnit.GetBody(msg.BodyID))
	} else {
		unit.DelBody()
	}

	if msg.RadarID != 0 {
		unit.SetRadar(DetailUnit.GetRadar(msg.RadarID))
	} else {
		unit.DelRadar()
	}

	unit.CalculateParametersUnit()

	resp := Response{Event: "UnitConstructorUpdate", Unit: *unit}
	ws.WriteJSON(resp)
}
