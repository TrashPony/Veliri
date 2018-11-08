package movePhase

import (
	"../../../gameObjects/unit"
	"../../../localGame"
)

func SkipMove(gameUnit *unit.Unit, game *localGame.Game) {
	gameUnit.ActionPoints = 0
	gameUnit.Move = false
	QueueMove(game)
}
