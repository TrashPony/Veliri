package movePhase

import (
	"../../../localGame"
)

func Queue(game *localGame.Game) int {
	queue := 0

	for _, xLine := range game.GetUnits() {
		for _, gameUnit := range xLine {
			if gameUnit.Action {
				if gameUnit.QueueAttack > queue {
					queue = gameUnit.QueueAttack
				}
			}
		}
	}

	return queue + 1
}
