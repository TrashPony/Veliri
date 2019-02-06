package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"math"
)

func GetQRfromXY(x, y int, mp *_map.Map) *coordinate.Coordinate {
	var minDist int
	var hexCoordinate *coordinate.Coordinate

	for _, qLine := range mp.OneLayerMap {
		for _, mapCoordinate := range qLine {
			xc, yc := GetXYCenterHex(mapCoordinate.Q, mapCoordinate.R)
			dist := int(GetBetweenDist(x, y, xc, yc))

			if dist < minDist || minDist == 0 {
				minDist = dist
				hexCoordinate = mapCoordinate
			}
		}
	}
	return hexCoordinate
}

func GetXYCenterHex(q, r int) (int, int) {
	var x, y int

	if r%2 != 0 {
		x = HexagonWidth + (HorizontalOffset * q)
	} else {
		x = HexagonWidth/2 + (HorizontalOffset * q)
	}
	y = HexagonHeight/2 + (r * VerticalOffset)

	return x, y
}

func GetBetweenDist(fromX, fromY, toX, toY int) float64 {
	var dx = toX - fromX
	var dy = toY - fromY
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func RotateUnit(unitRotate, needRotate *int) int {

	if *unitRotate < 0 {
		*unitRotate += 360
	}

	if *unitRotate > 360 {
		*unitRotate -= 360
	}

	if *needRotate < 0 {
		*needRotate += 360
	}

	if *needRotate > 360 {
		*needRotate -= 360
	}

	if unitRotate != needRotate {
		if directionRotate(*unitRotate, *needRotate) {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

func directionRotate(unitAngle, needAngle int) bool {
	// true ++
	// false --
	count := 0
	direction := false

	if unitAngle < needAngle {
		for unitAngle < needAngle {
			count++
			direction = true
			unitAngle++
		}
	} else {
		for unitAngle > needAngle {
			count++
			direction = false
			needAngle++
		}
	}

	if direction {
		return count <= 180
	} else {
		return !(count <= 180)
	}
}
