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

func RotatePoint(x, y, x0, y0 float64, rotate int) (newX, newY float64) {
	// поворачиваем квадрат по формуле (x0:y0 - центр)
	//X = (x — x0) * cos(alpha) — (y — y0) * sin(alpha) + x0;
	//Y = (x — x0) * sin(alpha) + (y — y0) * cos(alpha) + y0;

	alpha := float64(rotate) * math.Pi / 180
	newX = (x-x0)*math.Cos(alpha) - (y-y0)*math.Sin(alpha) + x0
	newY = (x-x0)*math.Sin(alpha) + (y-y0)*math.Cos(alpha) + y0
	return
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
