package movePhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
)

func SkipMove(gameUnit *unit.Unit, game *localGame.Game) {
	gameUnit.ActionPoints = 0
	gameUnit.Move = false
	QueueMove(game)
}
