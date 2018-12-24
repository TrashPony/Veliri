package globalGame

import (
	"../factories/bases"
	"../gameObjects/base"
	"../gameObjects/map"
	"../player"
	"errors"
	"math"
	"sync"
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
	Speed       float64
}

func MoveSquad(user *player.Player, ToX, ToY float64, mp *_map.Map) []PathUnit {
	startX := float64(user.GetSquad().GlobalX)
	startY := float64(user.GetSquad().GlobalY)
	rotate := user.GetSquad().MatherShip.Rotate

	maxSpeed := float64(user.GetSquad().MatherShip.Speed * 3)
	minSpeed := float64(user.GetSquad().MatherShip.Speed)
	speed := float64(user.GetSquad().MatherShip.Speed)

	// если текущая скорость выше стартовой то берем ее
	if float64(user.GetSquad().MatherShip.Speed) < user.GetSquad().CurrentSpeed {
		speed = user.GetSquad().CurrentSpeed
	}

	return MoveTo(startX, startY, maxSpeed, minSpeed, speed, ToX, ToY, rotate, mp, false)
}

func LaunchEvacuation(user *player.Player, mp *_map.Map) ([]PathUnit, int, *base.Transport, error) {
	mapBases := bases.Bases.GetBasesByMap(mp.Id)
	minDist := 0.0
	evacuationBase := &base.Base{}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	for _, mapBase := range mapBases {

		x, y := GetXYCenterHex(mapBase.Q, mapBase.R)
		dist := GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
		transport := mapBase.GetFreeTransport()

		if (dist < minDist || minDist == 0) && transport != nil {
			minDist = dist
			evacuationBase = mapBase
		}
	}

	if evacuationBase != nil {
		transport := evacuationBase.GetFreeTransport()
		if transport != nil {
			var startX, startY int

			transport.Job = true

			if transport.X == 0 && transport.Y == 0 {
				startX, startY = GetXYCenterHex(evacuationBase.Q, evacuationBase.R)
			} else {
				startX = transport.X
				startY = transport.Y
			}

			return MoveTo(float64(startX), float64(startY), 250, 15, 15,
					float64(user.GetSquad().GlobalX), float64(user.GetSquad().GlobalY), 0, mp, true),
				evacuationBase.ID, transport, nil
		} else {
			return nil, 0, nil, errors.New("no available transport")
		}
	} else {
		return nil, 0, nil, errors.New("no available base")
	}
}

func ReturnEvacuation(user *player.Player, mp *_map.Map, baseID int) []PathUnit {
	mapBase, _ := bases.Bases.Get(baseID)
	endX, endY := GetXYCenterHex(mapBase.Q, mapBase.R)

	return MoveTo(float64(user.GetSquad().GlobalX), float64(user.GetSquad().GlobalY), 250, 15, 15,
		float64(endX), float64(endY), 0, mp, true)
}

func MoveTo(forecastX, forecastY, maxSpeed, minSpeed, speed, ToX, ToY float64, rotate int, mp *_map.Map, ignoreObstacle bool) []PathUnit {

	path := make([]PathUnit, 0)

	diffRotate := 0 // разница между углом цели и носа корпуса
	dist := 900.0

	for {
		forecastQ := 0
		forecastR := 0
		// находим длинную вектора до цели
		dist = GetBetweenDist(int(forecastX), int(forecastY), int(ToX), int(ToY))
		if dist < 10 {
			break
		}

		minDist := float64(speed) / ((2 * math.Pi) / float64(360/speed)) // TODO не правильно

		if dist > maxSpeed*25 {
			if maxSpeed > speed {
				if len(path)%2 == 0 {
					speed += minSpeed / 10
				}
			} else {
				speed = maxSpeed
			}
		} else {
			if minSpeed < speed {
				if len(path)%2 == 0 {
					speed -= minSpeed / 10
				}
			} else {
				speed = minSpeed
			}
		}

		for i := 0; i < int(speed); i++ { // т.к. за 1 учаток пути корпус может повернуться на много градусов тут этот for)
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

		radRotate := float64(rotate) * math.Pi / 180
		stopX := float64(speed) * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := float64(speed) * math.Sin(radRotate)

		possibleMove, q, r := CheckXYinMove(int(forecastX+stopX), int(forecastY+stopY), rotate, mp)

		if (diffRotate == 0 || dist > minDist) && (possibleMove || ignoreObstacle) {
			forecastX = forecastX + stopX
			forecastY = forecastY + stopY

			forecastQ = q
			forecastR = r
		} else {
			if diffRotate == 0 {
				break
			}
		}

		path = append(path, PathUnit{X: int(forecastX), Y: int(forecastY), Rotate: rotate, Millisecond: 100,
			Q: forecastQ, R: forecastR, Speed: speed})
	}

	return path
}

func CheckXYinMove(x, y, rotate int, mp *_map.Map) (bool, int, int) {
	bodyRadius := 55 // размеры подобраны методом тыка)
	coordinateRadius := HexagonHeight / 2

	minDist := 999

	var q, r int

	for _, qLine := range mp.OneLayerMap {
		for _, mapCoordinate := range qLine {
			xc, yc := GetXYCenterHex(mapCoordinate.Q, mapCoordinate.R)

			//находим растояние координаты от места остановки
			dist := int(GetBetweenDist(x, y, xc, yc))

			// если координата находиться в теоритическом радиусе радиусе то проверяем на колизии
			if dist <= HexagonHeight {

				if minDist > dist {
					minDist = dist
					q = mapCoordinate.Q
					r = mapCoordinate.R
				}

				for i := rotate - 35; i < rotate+35; i++ { // смотрим только предметы по курсу )
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
