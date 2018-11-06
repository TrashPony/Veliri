package field

import (
	"../../mechanics/gameObjects/detail"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/localGame/Phases/targetPhase"
	"github.com/gorilla/websocket"
)

func SetTargetUnitEquip(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := Games.Get(client.GetGameID())

	ok := false
	equipSlot := &detail.BodyEquipSlot{}

	if msg.EquipType == 3 {
		equipSlot, ok = gameUnit.Body.EquippingIII[msg.NumberSlot]
	}

	if msg.EquipType == 2 {
		equipSlot, ok = gameUnit.Body.EquippingII[msg.NumberSlot]
	}

	if findClient && findUnit && findGame && ok && equipSlot.Equip != nil {
		if !client.GetReady() && !equipSlot.Used {

			var targetUnits []*unit.Unit

			if equipSlot.Equip.Applicable == "my_units" {
				targetUnits = targetPhase.GetEquipMyUnitsTarget(gameUnit, equipSlot.Equip, activeGame, client)
			}

			if equipSlot.Equip.Applicable == "hostile_units" {
				targetUnits = targetPhase.GetEquipHostileUnitsTarget(gameUnit, equipSlot.Equip, activeGame, client)
			}

			if equipSlot.Equip.Applicable == "myself" {
				// todo на себя
			}

			if equipSlot.Equip.Applicable == "all" {
				// todo  и свои и чужие но не карта GetEquipAllUnitTarget
			}

			for _, targetUnit := range targetUnits {
				if targetUnit.Q == msg.ToQ && targetUnit.R == msg.ToR {
					targetCoordinate, ok := activeGame.Map.GetCoordinate(targetUnit.Q, targetUnit.R)
					if ok {
						err := targetPhase.SetEquipTarget(gameUnit, targetCoordinate, equipSlot, client)
						if err != nil {
							ws.WriteJSON(ErrorMessage{Event: "Error", Error: err.Error()})
						} else {
							ws.WriteJSON(Unit{Event: "UpdateUnit", Unit: gameUnit})
						}
					}
				}
			}
		} else {
			ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not allow"})
		}
	}
}
