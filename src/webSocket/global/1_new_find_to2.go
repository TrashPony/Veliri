package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
	"math"
)

func To2(forecastX, forecastY, speed, ToX, ToY float64, rotate, rotateAngle, ms int) (error, []*unit.PathUnit) {

	// TODO искать предварительно часть пути до тех пока не будет разница в углу 0
	// TODO в разных вариантах, разворот на скорости или развород на месте, выбирать то где мешье частей пути (выше скорость)
	//  или если 1 из способов не пройти
	speed = speed / float64(1000/ms)
	path := make([]*unit.PathUnit, 0)

	for {

		// находим длинную вектора до цели
		dist := game_math.GetBetweenDist(int(forecastX), int(forecastY), int(ToX), int(ToY))
		if dist < speed+5 {
			break
		}

		radRotate := float64(rotate) * math.Pi / 180
		stopX := float64(speed) * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := float64(speed) * math.Sin(radRotate)

		//находим какой угол необходимо принять телу
		needRotate := game_math.GetBetweenAngle(ToX, ToY, forecastX, forecastY)
		move.RotateUnit(&rotate, &needRotate, rotateAngle)

		forecastX = forecastX + stopX
		forecastY = forecastY + stopY

		q, r := game_math.GetQRfromXY(int(forecastX), int(forecastY))
		path = append(path, &unit.PathUnit{X: int(forecastX), Y: int(forecastY), Rotate: rotate, Millisecond: ms,
			Q: q, R: r, Speed: speed, Animate: true})
	}

	if len(path) > 1 {
		path[len(path)-1].Speed = 0 // на последней точке машина останавливается
	}

	return nil, path
}
