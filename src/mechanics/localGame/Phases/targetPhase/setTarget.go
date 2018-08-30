package targetPhase

import (
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../player"
	"../../../db/updateSquad"
	"math"
)

func SetTarget(gameUnit *unit.Unit, game *localGame.Game, targetX, targetY int, client *player.Player) {

	target, _ := game.Map.GetCoordinate(targetX, targetY)
	rotate := rotateUnit(gameUnit, targetX, targetY)

	gameUnit.Target = target
	gameUnit.Rotate = rotate

	updateSquad.Squad(client.GetSquad())
}

func rotateUnit(gameUnit *unit.Unit, targetX, targetY int)  int{
	rotate := math.Atan2(float64(targetY - gameUnit.Q), float64(targetX - gameUnit.R))

	rotate = rotate * 180/math.Pi

	if rotate < 0 {
		rotate = 360 + rotate
	}

	return int(rotate)
}
