package lobby

import (
	"../../lobby/DetailUnit"
	"../../lobby/Squad"
	"github.com/gorilla/websocket"
)

func UnitConstructor(ws *websocket.Conn, msg Message) {
	var unit Squad.Unit

	if msg.WeaponID != 0 {
		unit.SetWeapon(DetailUnit.GetWeapon(msg.WeaponID))
	}

	if msg.ChassisID != 0 {
		unit.SetChassis(DetailUnit.GetChass(msg.ChassisID))
	}

	if msg.TowerID != 0 {
		unit.SetTower(DetailUnit.GetTower(msg.TowerID))
	}

	if msg.BodyID != 0 {
		unit.SetBody(DetailUnit.GetBody(msg.BodyID))
	}

	if msg.RadarID != 0 {
		unit.SetRadar(DetailUnit.GetRadar(msg.RadarID))
	}

	unit.CalculateParametersUnit()

	resp := Response{Event: "UnitConstructorUpdate", Unit: unit}
	ws.WriteJSON(resp)
}
