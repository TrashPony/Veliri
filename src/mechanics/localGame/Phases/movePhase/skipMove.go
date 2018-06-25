package movePhase

import (
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../db/update"
)

func SkipMove(gameUnit *unit.Unit, game *localGame.Game)  {
	gameUnit.Action = true

	queue := Queue(game)
	gameUnit.QueueAttack = queue

	update.Unit(gameUnit)
}
