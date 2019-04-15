package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/attackPhase"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/targetPhase"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func SetTargetUnitEquip(msg Message, client *player.Player) {

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

	if findUnit && findGame && ok && equipSlot.Equip != nil {
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
							SendMessage(ErrorMessage{Event: "Error", Error: err.Error()}, client.GetID(), activeGame.Id)
						} else {
							SendMessage(Unit{Event: "UpdateUnit", Unit: gameUnit}, client.GetID(), activeGame.Id)
							SendMessage(Message{Event: "QueueAttack", QueueAttack: attackPhase.CreateQueueAttack(client.GetUnitsINTKEY())}, client.GetID(), activeGame.Id)
						}
					}
				}
			}
		} else {
			SendMessage(ErrorMessage{Event: "Error", Error: "not allow"}, client.GetID(), activeGame.Id)
		}
	}
}