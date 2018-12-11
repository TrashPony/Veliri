package field

import (
	"../../mechanics/factories/games"
	"../../mechanics/gameObjects/detail"
	"../../mechanics/localGame/Phases/targetPhase"
	"github.com/gorilla/websocket"
)

func SetTargetMapEquip(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := games.Games.Get(client.GetGameID())

	ok := false
	equipSlot := &detail.BodyEquipSlot{}

	if msg.EquipType == 3 {
		equipSlot, ok = gameUnit.Body.EquippingIII[msg.NumberSlot]
	}

	if msg.EquipType == 2 {
		equipSlot, ok = gameUnit.Body.EquippingII[msg.NumberSlot]
	}

	if findUnit && findClient && findGame && !client.GetReady() && ok && equipSlot.Equip != nil {
		if equipSlot.Equip.Applicable == "map" {
			gameCoordinate, findCoordinate := activeGame.Map.GetCoordinate(msg.TargetQ, msg.TargetR)
			if findCoordinate {
				err := targetPhase.SetEquipTarget(gameUnit, gameCoordinate, equipSlot, client)
				if err == nil {
					ws.WriteJSON(Unit{Event: "UpdateUnit", Unit: gameUnit})
				} else {
					ws.WriteJSON(ErrorMessage{Event: "Error", Error: err.Error()})
				}
			} else {
				ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not find game coordinate"})
			}
		}
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not allow"})
	}
}
