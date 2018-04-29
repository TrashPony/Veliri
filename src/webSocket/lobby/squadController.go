package lobby

import (
	"../../lobby/DetailUnit"
	"../../lobby/Squad"
	"github.com/gorilla/websocket"
)

func SquadSettings(ws *websocket.Conn, msg Message)  {

	if msg.Event == "AddNewSquad" {
		err, squad := Squad.AddNewSquad(msg.SquadName, usersLobbyWs[ws].Id)

		var resp Response

		if err != nil {
			resp = Response{Event: "AddNewSquad", Error: err.Error()}
			ws.WriteJSON(resp)
		} else {
			usersLobbyWs[ws].Squads = append(usersLobbyWs[ws].Squads, squad)
			resp = Response{Event: "AddNewSquad", Error: "none", Squad: squad}
			ws.WriteJSON(resp)
		}
	}

	if msg.Event == "SelectSquad" {
		for _, squad := range  usersLobbyWs[ws].Squads {
			if squad.ID == msg.SquadID {
				usersLobbyWs[ws].Squad = squad
				resp := Response{Event: "SelectSquad", Error: "none", Squad: squad}
				ws.WriteJSON(resp)
			}
		}
	}

	if msg.Event == "SelectMatherShip" {
		if usersLobbyWs[ws].Squad != nil {
			if usersLobbyWs[ws].Squad.MatherShip != nil {
				usersLobbyWs[ws].Squad.ReplaceMatherShip(msg.MatherShipID)
			} else {
				usersLobbyWs[ws].Squad.AddMatherShip(msg.MatherShipID)
			}
			resp := Response{Event: "UpdateSquad", Squad: usersLobbyWs[ws].Squad}
			ws.WriteJSON(resp)
		}
	}
}

func GetDetailSquad(ws *websocket.Conn, msg Message)  {
	if msg.Event == "GetDetailOfUnits" {

		weapons := DetailUnit.GetWeapons()
		chassis := DetailUnit.GetChassis()
		towers := DetailUnit.GetTowers()
		bodies := DetailUnit.GetBodies()
		radars := DetailUnit.GetRadars()

		var resp = Response{Event: msg.Event, Weapons: weapons, Chassis: chassis, Towers: towers, Bodies: bodies, Radars: radars}
		ws.WriteJSON(resp)
	}

	if msg.Event == "GetEquipping" {
		var equipping = Squad.GetTypeEquipping()
		var resp = Response{Event: msg.Event, Equipping: equipping}
		ws.WriteJSON(resp)
	}

	if msg.Event == "GetListSquad" {
		squads, err := Squad.GetUserSquads(usersLobbyWs[ws].Id)

		var resp Response

		if err != nil {
			resp = Response{Event: "GetListSquad", Error: err.Error()}
			ws.WriteJSON(resp)
		} else {
			usersLobbyWs[ws].Squads = squads
			resp = Response{Event: "GetListSquad", Error: "none", Squads: squads}
			ws.WriteJSON(resp)
		}
	}

	if msg.Event == "GetMatherShips" {
		var matherShips = Squad.GetTypeMatherShips()
		var resp = Response{Event: msg.Event, MatherShips: matherShips}
		ws.WriteJSON(resp)
	}
}

func UnitSquad(ws *websocket.Conn, msg Message)  {
	if msg.Event == "AddUnitInSquad" || msg.Event == "ReplaceUnitInSquad" {
		if usersLobbyWs[ws].Squad != nil {
			var unit Squad.Unit

			weapon := DetailUnit.GetWeapon(msg.WeaponID)
			chassis := DetailUnit.GetChass(msg.ChassisID)
			tower := DetailUnit.GetTower(msg.TowerID)
			body := DetailUnit.GetBody(msg.BodyID)
			radar := DetailUnit.GetRadar(msg.RadarID)

			unit.SetChassis(&chassis)
			unit.SetWeapon(&weapon)
			unit.SetTower(&tower)
			unit.SetBody(&body)
			unit.SetRadar(&radar)

			if msg.Event == "AddUnitInSquad" {
				usersLobbyWs[ws].Squad.AddUnit(&unit, msg.UnitSlot)
			} else {
				usersLobbyWs[ws].Squad.ReplaceUnit(&unit, msg.UnitSlot)
			}
		} else {
			resp := Response{Event: msg.Event, Error: "No select squad"}
			ws.WriteJSON(resp)
		}
	}

	if msg.Event == "RemoveUnitInSquad" {
		if usersLobbyWs[ws].Squad != nil {
			err := usersLobbyWs[ws].Squad.DelUnit(msg.UnitSlot)
			if err == nil {
				resp := Response{Event: "RemoveUnitInSquad", Error: "none", UnitSlot: msg.UnitSlot}
				ws.WriteJSON(resp)
			} else {
				resp := Response{Event: "RemoveUnitInSquad", Error: err.Error(), UnitSlot: msg.UnitSlot}
				ws.WriteJSON(resp)
			}
		} else {
			resp := Response{Event: msg.Event, Error: "No select squad"}
			ws.WriteJSON(resp)
		}
	}
}

func EquipSquad(ws *websocket.Conn, msg Message)  {
	if msg.Event == "AddEquipment" || msg.Event == "ReplaceEquipment" {
		if usersLobbyWs[ws].Squad != nil {
			equip := Squad.GetTypeEquip(msg.EquipID)

			if msg.Event == "AddEquipment" {
				usersLobbyWs[ws].Squad.AddEquip(&equip, msg.EquipSlot)
			} else {
				usersLobbyWs[ws].Squad.ReplaceEquip(&equip, msg.EquipSlot)
			}

			resp := Response{Event: "UpdateSquad", Squad: usersLobbyWs[ws].Squad}
			ws.WriteJSON(resp)

		} else {
			resp := Response{Event: msg.Event, Error: "No select squad"}
			ws.WriteJSON(resp)
		}
	}

	if msg.Event == "RemoveEquipment" {
		if usersLobbyWs[ws].Squad != nil {
			err := usersLobbyWs[ws].Squad.DelEquip(msg.EquipSlot)
			if err == nil {
				resp := Response{Event: msg.Event, Error: "none", UnitSlot: msg.EquipSlot}
				ws.WriteJSON(resp)
			} else {
				resp := Response{Event: msg.Event, Error: err.Error(), UnitSlot: msg.EquipSlot}
				ws.WriteJSON(resp)
			}

			resp := Response{Event: "UpdateSquad", Squad: usersLobbyWs[ws].Squad}
			ws.WriteJSON(resp)
		} else {
			resp := Response{Event: msg.Event, Error: "No select squad"}
			ws.WriteJSON(resp)
		}
	}
}