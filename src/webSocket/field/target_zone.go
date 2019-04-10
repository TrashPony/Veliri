package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/targetPhase"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func GetTargetZone(msg Message, client *player.Player) {

	activeGame, findGame := games.Games.Get(client.GetGameID())

	gameUnit, findUnit := client.GetUnitStorage(msg.UnitID)
	if !findUnit {
		gameUnit, findUnit = client.GetHostileUnit(msg.Q, msg.R)
		if !findUnit {
			gameUnit, findUnit = client.GetUnit(msg.Q, msg.R)
		}
	}

	if findUnit && findGame {

		tmpUnit := *gameUnit

		tmpUnit.SetQ(msg.ToQ)
		tmpUnit.SetR(msg.ToR)

		SendMessage(
			TargetCoordinate{
				Event:   "GetFirstTargets",
				Unit:    gameUnit,
				Targets: targetPhase.GetWeaponTargetCoordinate(&tmpUnit, activeGame, client, "GetFirstTargets"),
			},
			client.GetID(),
			activeGame.Id,
		)
	}
}
