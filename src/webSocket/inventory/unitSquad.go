package inventory

import (
	"github.com/gorilla/websocket"
	"../../inventory"
	"../../detailUnit"
)

func UnitSquad(ws *websocket.Conn, msg Message)  {
	if msg.Event == "AddUnit" || msg.Event == "ReplaceUnit" {
		if usersInventoryWs[ws].Squad != nil {
			var unit inventory.Unit
			// todo проверка на занятость слота в который хотят добавить юнита
			if msg.WeaponID != 0 {
				unit.SetWeapon(detailUnit.GetWeapon(msg.WeaponID))
			}

			if msg.ChassisID != 0 {
				unit.SetChassis(detailUnit.GetChass(msg.ChassisID))
			}

			if msg.TowerID != 0 {
				unit.SetTower(detailUnit.GetTower(msg.TowerID))
			}

			if msg.BodyID != 0 {
				unit.SetBody(detailUnit.GetBody(msg.BodyID))
			}

			if msg.RadarID != 0 {
				unit.SetRadar(detailUnit.GetRadar(msg.RadarID))
			}

			unit.CalculateParametersUnit()

			if msg.Event == "AddUnit" {
				usersInventoryWs[ws].Squad.AddUnit(&unit, msg.UnitSlot)
			} else {
				usersInventoryWs[ws].Squad.ReplaceUnit(&unit, msg.UnitSlot)
			}

			resp := Response{Event: "UpdateSquad", Squad: usersInventoryWs[ws].Squad}
			ws.WriteJSON(resp)
		} else {
			resp := Response{Event: msg.Event, Error: "No select squad"}
			ws.WriteJSON(resp)
		}
	}

	if msg.Event == "RemoveUnit" {
		if usersInventoryWs[ws].Squad != nil {
			err := usersInventoryWs[ws].Squad.DelUnit(msg.UnitSlot)

			if err == nil {
				resp := Response{Event: msg.Event, Error: "none", UnitSlot: msg.UnitSlot}
				ws.WriteJSON(resp)
			} else {
				resp := Response{Event: msg.Event, Error: err.Error(), UnitSlot: msg.UnitSlot}
				ws.WriteJSON(resp)
			}

			resp := Response{Event: "UpdateSquad", Squad: usersInventoryWs[ws].Squad}
			ws.WriteJSON(resp)
		} else {
			resp := Response{Event: msg.Event, Error: "No select squad"}
			ws.WriteJSON(resp)
		}
	}
}

func UnitConstructor(ws *websocket.Conn, msg Message) {

	unit, ok := usersInventoryWs[ws].Squad.Units[msg.UnitSlot]
	if !ok {
		unit = &inventory.Unit{}
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
