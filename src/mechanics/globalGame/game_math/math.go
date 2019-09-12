package game_math

import (
	"math"
)

const HexagonHeight = 55 // Константы описывающие свойства гексов на игровом поле
const HexagonWidth = 50
const VerticalOffset = HexagonHeight * 3 / 4
const HorizontalOffset = HexagonWidth

func GetQRfromXY(x, y int) (int, int) {

	rF := (float64(y) - HexagonHeight/2) / VerticalOffset
	r := int(math.Round(rF))
	q := 0

	if r%2 != 0 {
		qF := (float64(x) - HexagonWidth) / HexagonWidth
		q = int(math.Round(qF))
	} else {
		qF := (float64(x) - HexagonWidth/2) / HexagonWidth
		q = int(math.Round(qF))
	}

	return q, r
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