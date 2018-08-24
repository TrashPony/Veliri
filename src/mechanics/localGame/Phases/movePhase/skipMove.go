package movePhase

import (
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../player"
	"../../../db/updateSquad"
)

func SkipMove(gameUnit *unit.Unit, game *localGame.Game, client *player.Player)  {
	gameUnit.Action = true

	queue := Queue(game)
	gameUnit.QueueAttack = queue

	updateSquad.Squad(client.GetSquad())
}
