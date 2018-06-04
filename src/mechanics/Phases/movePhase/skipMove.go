package movePhase

import (
	"../../unit"
	"../../game"
	"../../db"
)

func SkipMove(gameUnit *unit.Unit, game *game.Game)  {
	gameUnit.Action = true

	queue := Queue(game)
	gameUnit.Queue = queue

	db.UpdateUnit(gameUnit)
}
