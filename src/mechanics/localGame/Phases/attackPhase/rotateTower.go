package attackPhase

import (
	"../../../gameObjects/coordinate"
	"math"
)

func rotateTower(unitCoordinate, target *coordinate.Coordinate) int {

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
