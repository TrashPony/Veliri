package movePhase

import (
	"../../../db/updateSquad"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../player"
)

func SkipMove(gameUnit *unit.Unit, game *localGame.Game, client *player.Player) {
	gameUnit.ActionPoints = 0

	queue := Queue(game)
	gameUnit.QueueAttack = queue

	updateSquad.Squad(client.GetSquad())
}
