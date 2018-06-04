package movePhase

import (
	"../../game"
)

func Queue(game *game.Game) int {
	queue := 0

	for _, xLine := range game.GetUnits() {
		for _, gameUnit := range xLine {
			if gameUnit.Action {
				if gameUnit.Queue > queue {
					queue = gameUnit.Queue
				}
			}
		}
	}

	return queue + 1
}
