package globalGame

import (
	"../gameObjects/map"
	"../player"
	"math"
)

const HexagonHeight = 111 // Константы описывающие свойства гексов на игровом поле
const HexagonWidth = 100
const VerticalOffset = HexagonHeight * 3 / 4
const HorizontalOffset = HexagonWidth

type PathUnit struct {
	X           int `json:"x"`
	Y           int `json:"y"`
	Q           int `json:"q"`
	R           int `json:"r"`
	Rotate      int `json:"rotate"`
	Millisecond int `json:"millisecond"`
}

func MoveTo(user *player.Player, ToX, ToY float64, mp *_map.Map) []PathUnit {

	path := make([]PathUnit, 0)

	forecastX := float64(user.GetSquad().GlobalX)
	forecastY := float64(user.GetSquad().GlobalY)
	speed := user.GetSquad().MatherShip.Speed * 3
	rotate := user.GetSquad().MatherShip.Rotate

	diffRotate := 0 // разница между углом цели и носа корпуса
	dist := 900.0

	for {
		forecastQ := 0
		forecastR := 0
		// находим длинную вектора до цели
		dist = math.Sqrt(((forecastX - ToX) * (forecastX - ToX)) + ((forecastY - ToY) * (forecastY - ToY)))
		if dist < 10 {
			break
		}

		for i := 0; i < speed; i++ { // т.к. за 1 учаток пути корпус может повернуться на много градусов тут этот for)
			needRad := math.Atan2(ToY-forecastY, ToX-forecastX)
			needRotate := int(needRad * 180 / 3.14) // находим какой угол необходимо принять телу

			newRotate := RotateUnit(&rotate, &needRotate)

			if rotate >= needRotate {
				diffRotate = rotate - needRotate
			} else {
				diffRotate = needRotate - rotate
			}

			if diffRotate != 0 { // если разница есть то поворачиваем корпус
				rotate += newRotate
			} else {
				break
			}
		}

		minDist := float64(speed) / ((2 * math.Pi) / float64(360/speed)) // TODO не правильно

		radRotate := float64(rotate) * math.Pi / 180
		stopX := float64(speed) * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := float64(speed) * math.Sin(radRotate)

		possibleMove, q, r := CheckXYinMove(int(forecastX+stopX), int(forecastY+stopY), rotate, mp)

		if (diffRotate == 0 || dist > minDist) && possibleMove {
			forecastX = forecastX + stopX
			forecastY = forecastY + stopY

			forecastQ = q
			forecastR = r
		} else {
			if diffRotate == 0 {
				break
			}
		}

		path = append(path, PathUnit{X: int(forecastX), Y: int(forecastY), Rotate: rotate, Millisecond: 100, Q: forecastQ, R: forecastR})
	}

	return path
}

func CheckXYinMove(x, y, rotate int, mp *_map.Map) (bool, int, int) {
	bodyRadius := 55 // размеры подобраны методом тыка)
	coordinateRadius := HexagonHeight/2

	minDist := 999

	var q, r int

	for _, qLine := range mp.OneLayerMap {
		for _, mapCoordinate := range qLine {
			xc, yc := GetXYCenterHex(mapCoordinate.Q, mapCoordinate.R)

			//находим растояние координаты от места остановки
			dist := (x-xc)*(x-xc) + (y-yc)*(y-yc)

			// если координата находиться в теоритическом радиусе радиусе то проверяем на колизии
			if dist <= HexagonHeight*HexagonHeight {

				if minDist > dist {
					minDist = dist
					q = mapCoordinate.Q
					r = mapCoordinate.R
				}

				for i := rotate - 35; i < rotate + 35; i++ { // смотрим только предметы по курсу )
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyRadius)*math.Cos(rad)) + x // точки окружности корпуса
					bY := int(float64(bodyRadius)*math.Sin(rad)) + y

					if ((bX-xc)*(bX-xc) + (bY-yc)*(bY-yc)) <= coordinateRadius*coordinateRadius {
						if !mapCoordinate.Move {
							return false, q, r
						}
					}
				}
			}
		}
	}
	return true, q, r
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
