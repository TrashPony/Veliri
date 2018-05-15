package lobby

import (
	"../../detailUnit"
	"../../lobby/Squad"
	"github.com/gorilla/websocket"
)

func UnitConstructor(ws *websocket.Conn, msg Message) {

	unit, ok := usersLobbyWs[ws].Squad.Units[msg.UnitSlot]
	if !ok {
		unit = &Squad.Unit{}
	}

	if msg.WeaponID != 0 {
		unit.SetWeapon(detailUnit.GetWeapon(msg.WeaponID))
	} else {
		unit.DelWeapon()
	}

	if msg.ChassisID != 0 {
		unit.SetChassis(detailUnit.GetChass(msg.ChassisID))
	} else {
		unit.DelChassis()
	}

	if msg.TowerID != 0 {
		unit.SetTower(detailUnit.GetTower(msg.TowerID))
	} else {
		unit.DelTower()
	}

	if msg.BodyID != 0 {
		unit.SetBody(detailUnit.GetBody(msg.BodyID))
	} else {
		unit.DelBody()
	}

	if msg.RadarID != 0 {
		unit.SetRadar(detailUnit.GetRadar(msg.RadarID))
	} else {
		unit.DelRadar()
	}

	unit.CalculateParametersUnit()

	resp := Response{Event: "UnitConstructorUpdate", Unit: *unit}
	ws.WriteJSON(resp)
}
