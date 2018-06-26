package targetPhase

import (
	"../../../gameObjects/unit"
	"../../../db/localGame/update"
	"../../../localGame"
	"math"
)

func SetTarget(gameUnit *unit.Unit, game *localGame.Game, targetX, targetY int) {

	target, _ := game.Map.GetCoordinate(targetX, targetY)
	rotate := rotateUnit(gameUnit, targetX, targetY)

	gameUnit.Target = target
	gameUnit.Rotate = rotate

	update.Unit(gameUnit)
}

func rotateUnit(gameUnit *unit.Unit, targetX, targetY int)  int{
	rotate := math.Atan2(float64(targetY - gameUnit.Y), float64(targetX - gameUnit.X))

	rotate = rotate * 180/math.Pi

	if rotate < 0 {
		rotate = 360 + rotate
	}

	return int(rotate)
}
