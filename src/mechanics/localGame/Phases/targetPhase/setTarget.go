package targetPhase

import (
	"../../../db/updateSquad"
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../player"
	"math"
)

func SetTarget(gameUnit *unit.Unit, game *localGame.Game, targetQ, targetR int, client *player.Player) {

	target, _ := game.Map.GetCoordinate(targetQ, targetR)
	unitCoordinate, _ := game.Map.GetCoordinate(gameUnit.Q, gameUnit.R)

	rotate := rotateUnit(unitCoordinate, target)

	gameUnit.Target = target
	gameUnit.Rotate = rotate

	updateSquad.Squad(client.GetSquad())
}

func rotateUnit(unitCoordinate, target *coordinate.Coordinate) int {

	//http://zvold.blogspot.com/2010/01/bresenhams-line-drawing-algorithm-on_26.html

	var rotate float64

	if unitCoordinate.R%2 != 0 {
		if (target.R-unitCoordinate.R)%2 != 0 {
			rotate = math.Atan2(float64(target.R)-float64(unitCoordinate.R), (float64(target.Q)-0.5)-float64(unitCoordinate.Q))
		} else {
			rotate = math.Atan2(float64(target.R-unitCoordinate.R), float64(target.Q-unitCoordinate.Q))
		}
	} else {
		if (target.R-unitCoordinate.R)%2 != 0 {
			rotate = math.Atan2(float64(target.R)-float64(unitCoordinate.R), float64(target.Q)-(float64(unitCoordinate.Q)-0.5))
		} else {
			rotate = math.Atan2(float64(target.R-unitCoordinate.R), float64(target.Q-unitCoordinate.Q))
		}
	}

	rotate = rotate * 180 / math.Pi
	if rotate < 0 {
		rotate = 360 + rotate
	}

	return int(rotate)
}
