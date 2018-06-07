package targetPhase

import (
	"../../unit"
	"../../game"
	"../../db"
	"math"
)

func SetTarget(gameUnit *unit.Unit, game *game.Game, targetX, targetY int) {

	target, _ := game.Map.GetCoordinate(targetX, targetY)
	rotate := rotateUnit(gameUnit, targetX, targetY)

	gameUnit.Target = target
	gameUnit.Rotate = rotate

	db.UpdateUnit(gameUnit)
}

func rotateUnit(gameUnit *unit.Unit, targetX, targetY int)  int{
	rotate := math.Atan2(float64(targetY - gameUnit.Y), float64(targetX - gameUnit.X))

	rotate = rotate * 180/math.Pi

	if rotate < 0 {
		rotate = 360 + rotate
	}

	return int(rotate)
}
