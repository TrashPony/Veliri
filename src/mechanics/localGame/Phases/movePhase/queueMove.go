package movePhase

import (
	"../../../localGame"
)

func Queue(game *localGame.Game) int {
	queue := 0

	for _, xLine := range game.GetUnits() {
		for _, gameUnit := range xLine {
			if gameUnit.ActionPoints == 0 {
				if gameUnit.QueueAttack > queue {
					queue = gameUnit.QueueAttack
				}
			}
		}
	}

	return queue + 1
}
