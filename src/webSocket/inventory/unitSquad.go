package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/gameObjects/unit"
)

func UnitSquad(ws *websocket.Conn, msg Message)  {
	if msg.Event == "AddUnit" || msg.Event == "ReplaceUnit" {
		if usersInventoryWs[ws].Squad != nil {
			//var gameUnit unit.Unit
			// todo проверка на занятость слота в который хотят добавить юнита
			if msg.WeaponID != 0 {
				//gameUnit.SetWeapon(detailUnit.GetWeapon(msg.WeaponID))
			}

			if msg.BodyID != 0 {
				//gameUnit.SetBody(detailUnit.GetBody(msg.BodyID))
			}


			if msg.Event == "AddUnit" {
				//usersInventoryWs[ws].Squad.AddUnit(&gameUnit, msg.UnitSlot)
			} else {
				//usersInventoryWs[ws].Squad.ReplaceUnit(&gameUnit, msg.UnitSlot)
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
			/*err := usersInventoryWs[ws].Squad.DelUnit(msg.UnitSlot)

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
			ws.WriteJSON(resp)*/
		}
	}
}

func UnitConstructor(ws *websocket.Conn, msg Message) {

	gameUnit, ok := usersInventoryWs[ws].Squad.Units[msg.UnitSlot]
	if !ok {
		gameUnit = &unit.Unit{}
	}

	if msg.WeaponID != 0 {
		//gameUnit.SetWeapon(detailUnit.GetWeapon(msg.WeaponID))
	} else {
		gameUnit.DelWeapon()
	}

	if msg.BodyID != 0 {
		//gameUnit.SetBody(detailUnit.GetBody(msg.BodyID))
	} else {
		gameUnit.DelBody()
	}

	resp := Response{Event: "UnitConstructorUpdate", Unit: *gameUnit}
	ws.WriteJSON(resp)
}
