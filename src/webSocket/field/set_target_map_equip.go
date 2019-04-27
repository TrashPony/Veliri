package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/attackPhase"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/targetPhase"
)

func SetTargetMapEquip(msg Message, client *player.Player) {

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

	if findUnit && findGame && !client.GetReady() && ok && equipSlot.Equip != nil {
		if equipSlot.Equip.Applicable == "map" {
			gameCoordinate, findCoordinate := activeGame.Map.GetCoordinate(msg.TargetQ, msg.TargetR)
			if findCoordinate {
				err := targetPhase.SetEquipTarget(gameUnit, gameCoordinate, equipSlot, client)
				if err == nil {
					SendMessage(Unit{Event: "UpdateUnit", Unit: gameUnit}, client.GetID(), activeGame.Id)
					SendMessage(Message{Event: "QueueAttack", QueueAttack: attackPhase.CreateQueueAttack(client.GetUnitsINTKEY())}, client.GetID(), activeGame.Id)
				} else {
					SendMessage(ErrorMessage{Event: "Error", Error: err.Error()}, client.GetID(), activeGame.Id)
				}
			} else {
				SendMessage(ErrorMessage{Event: "Error", Error: "not find game coordinate"}, client.GetID(), activeGame.Id)
			}
		}
	} else {
		SendMessage(ErrorMessage{Event: "Error", Error: "not allow"}, client.GetID(), activeGame.Id)
	}
}
