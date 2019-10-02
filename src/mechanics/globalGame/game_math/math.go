package game_math

import (
	"math"
)

const CellSize = 30
const DiscreteSize = 6 * CellSize

func GetBetweenDist(fromX, fromY, toX, toY int) float64 {
	var dx = toX - fromX
	var dy = toY - fromY
	return math.Sqrt(float64(dx*dx + dy*dy))
}

//needRad := math.Atan2(float64(y-pY), float64(x-pX))
//needAngle := int(needRad * 180 / 3.14)

func GetBetweenAngle(x, y, targetX, targetY float64) int {
	//math.Atan2 куда у - текущие у, куда х - текущие х, получаем угол
	needRad := math.Atan2(y-targetY, x-targetX)
	return int(needRad * 180 / 3.14)
}

func GetBetweenDistLinePoint(xPoint, yPoint, x1Line, y1Line, x2Line, y2Line int) int {

	A := y1Line - y2Line
	B := x1Line - x2Line
	C := y1Line*x2Line - y2Line*x1Line

	dist := int(float64(A*xPoint+B*yPoint+C) / math.Sqrt(float64(A*A+B*B)))
	if dist < 0 {
		dist *= -1
	}
	return dist
}
