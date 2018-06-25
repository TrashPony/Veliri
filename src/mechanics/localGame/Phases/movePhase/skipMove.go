package movePhase

import (
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../db"
)

func SkipMove(gameUnit *unit.Unit, game *localGame.Game)  {
	gameUnit.Action = true

	queue := Queue(game)
	gameUnit.QueueAttack = queue

	db.UpdateUnit(gameUnit)
}
