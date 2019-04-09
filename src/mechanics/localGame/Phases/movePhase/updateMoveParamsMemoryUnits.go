package movePhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
)

func updateMoveParamsMemoryUnits(game *localGame.Game) {
	//todo вершина моего говнокода и он работает
	for _, gameUser := range game.GetPlayers() {
		for _, memoryUnit := range gameUser.GetMemoryHostileUnits() {
			for _, gameUser2 := range game.GetPlayers() {
				for _, q := range gameUser2.GetUnits() {
					for _, realUnit := range q {
						if memoryUnit.ID == realUnit.ID {
							gameUser.SetMoveParamsMemoryUnit(memoryUnit.ID, realUnit.Move, realUnit.ActionPoints)
						}
					}

					if gameUser2.GetSquad() != nil && memoryUnit.ID == gameUser2.GetSquad().MatherShip.ID {
						gameUser.SetMoveParamsMemoryUnit(memoryUnit.ID, gameUser2.GetSquad().MatherShip.Move, gameUser2.GetSquad().MatherShip.ActionPoints)
					}
				}
			}
		}
	}
}
