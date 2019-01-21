package globalGame

import (
	"../gameObjects/detail"
	"../gameObjects/map"
	"../player"
	"errors"
	"github.com/getlantern/deepcopy"
	"math"
)

const HexagonHeight = 55 // Константы описывающие свойства гексов на игровом поле
const HexagonWidth = 50
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

func MoveSquad(user *player.Player, ToX, ToY float64, mp *_map.Map) ([]PathUnit, error) {
	startX := float64(user.GetSquad().GlobalX)
	startY := float64(user.GetSquad().GlobalY)
	rotate := user.GetSquad().MatherShip.Rotate

	maxSpeed := float64(user.GetSquad().MatherShip.Speed*3)
	minSpeed := float64(user.GetSquad().MatherShip.Speed)
	speed := float64(user.GetSquad().MatherShip.Speed)

	// если текущая скорость выше стартовой то берем ее
	if float64(user.GetSquad().MatherShip.Speed) < user.GetSquad().CurrentSpeed {
		speed = user.GetSquad().CurrentSpeed
	}

	var fakeThoriumSlots map[int]*detail.ThoriumSlot

	// копируем что бы не произошло вычетание топлива на расчетах
	err := deepcopy.Copy(&fakeThoriumSlots, &user.GetSquad().MatherShip.Body.ThoriumSlots)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	if user.GetSquad().Afterburner { // если форсаж то х2 скорости
		maxSpeed = maxSpeed * 2
	}

	err, path := MoveTo(startX, startY, maxSpeed, minSpeed, speed, ToX, ToY, rotate, mp, false,
		fakeThoriumSlots, user.GetSquad().Afterburner, user.GetSquad().HighGravity)

	return path, err
}

func MoveTo(forecastX, forecastY, maxSpeed, minSpeed, speed, ToX, ToY float64, rotate int,
	mp *_map.Map, ignoreObstacle bool, thoriumSlots map[int]*detail.ThoriumSlot, afterburner, gravity bool) (error, []PathUnit) {

	path := make([]PathUnit, 0)

	fullMax := maxSpeed

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

		if thoriumSlots != nil {
			efficiency := WorkOutThorium(thoriumSlots, afterburner, gravity) // отрабатываем прредпологалаемое топливо
			maxSpeed = (fullMax * efficiency) / 100                          // высчитываем максимальную скорость по состоянию топлива
			if efficiency == 0 {
				// кончилось топливо, выходим с ошибкой
				return errors.New("not thorium"), path
			}
		}

		minDistRotate := float64(speed) / ((2 * math.Pi) / float64(360/speed))

		if dist > maxSpeed*25 { // TODO не правильно, тут надо расчитать растояние когда надо сбрасывать скорость
			if int(maxSpeed)*10 != int(speed)*10 {
				if maxSpeed > speed {
					if len(path)%2 == 0 {
						speed += minSpeed / 10
					}
				} else {
					speed -= minSpeed / 10
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

		radRotate := float64(rotate) * math.Pi / 180
		stopX := float64(speed) * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := float64(speed) * math.Sin(radRotate)

		for i := 0; i < int(minSpeed); i++ { // т.к. за 1 учаток пути корпус может повернуться на много градусов тут этот for)
			needRad := math.Atan2(ToY-forecastY, ToX-forecastX)
			//  math.Atan2 куда у - текущие у, куда х - текущие х, получаем угол
			needRotate := int(needRad * 180 / 3.14) // находим какой угол необходимо принять телу

			newRotate := RotateUnit(&rotate, &needRotate)

			diffRotate = rotate - needRotate
			if diffRotate < 0 {
				diffRotate = 360 - diffRotate
			}

			if diffRotate != 0 { // если разница есть то поворачиваем корпус
				rotate += newRotate
				if minSpeed < speed {
					speed -= minSpeed / (10 * speed) // сбрасывает скорость на поворотах
				}
			} else {
				break
			}
		}

		possibleMove, q, r, front := CheckCollisionsOnStaticMap(int(forecastX+stopX), int(forecastY+stopY), rotate, mp)

		if (diffRotate == 0 || dist > minDistRotate) && (possibleMove || ignoreObstacle) {
			forecastX = forecastX + stopX
			forecastY = forecastY + stopY

			forecastQ = q
			forecastR = r
		} else {
			if diffRotate == 0 {
				break
			}

			if !possibleMove {
				if front { // если препятвие спереди то идем назад
					forecastX = forecastX - stopX
					forecastY = forecastY - stopY
				} else { // иначе вперед
					forecastX = forecastX + stopX
					forecastY = forecastY + stopY
				}
			}
		}

		path = append(path, PathUnit{X: int(forecastX), Y: int(forecastY), Rotate: rotate, Millisecond: 100,
			Q: forecastQ, R: forecastR, Speed: speed})
	}

	if len(path) > 1 {
		path[len(path)-1].Speed = 0 // на последней точке машина останавливается
	}

	return nil, path
}
