package globalGame

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/getlantern/deepcopy"
	"math"
)

const HexagonHeight = 55 // Константы описывающие свойства гексов на игровом поле
const HexagonWidth = 50
const VerticalOffset = HexagonHeight * 3 / 4
const HorizontalOffset = HexagonWidth

func MoveUnit(moveUnit *unit.Unit, ToX, ToY float64, mp *_map.Map) ([]unit.PathUnit, error) {

	startX := float64(moveUnit.X)
	startY := float64(moveUnit.Y)
	rotate := moveUnit.Rotate

	//todo
	moveUnit.MinSpeed = 10

	// т.к. метод расчитывает в секунду а путь строится по 100 мс то скорость делим на 10
	maxSpeed := float64(moveUnit.Speed) / 10
	minSpeed := float64(moveUnit.MinSpeed) / 10

	// если текущая скорость выше стартовой то берем ее
	startSpeed := minSpeed
	if minSpeed < moveUnit.CurrentSpeed {
		startSpeed = moveUnit.CurrentSpeed
	}

	var fakeThoriumSlots map[int]*detail.ThoriumSlot

	// копируем что бы не произошло вычетание топлива на расчетах
	err := deepcopy.Copy(&fakeThoriumSlots, &moveUnit.Body.ThoriumSlots)
	if err != nil || len(fakeThoriumSlots) == 0 {
		println(err.Error())
		return nil, err
	}

	if moveUnit.Afterburner { // если форсаж то х2 скорости (доступно только МС)
		maxSpeed = maxSpeed * 2
	}

	err, path := MoveTo(startX, startY, maxSpeed, minSpeed, startSpeed, ToX, ToY, rotate, 5, mp, true,
		fakeThoriumSlots, moveUnit.Afterburner, moveUnit.HighGravity, moveUnit.Body)

	return path, err
}

func MoveTo(forecastX, forecastY, maxSpeed, minSpeed, speed, ToX, ToY float64, rotate, rotateAngle int, mp *_map.Map,
	ignoreObstacle bool, thoriumSlots map[int]*detail.ThoriumSlot, afterburner, gravity bool, body *detail.Body) (error, []unit.PathUnit) {

	path := make([]unit.PathUnit, 0)

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

		// скорость * (180/угол поворота) = длинна полокружности (грабая модель)
		// длинны окружности получаем ее радиус r= длинна полокружности /2 пи
		// r*2 получаем минимальное растояние от бтр до обьекта к которому он может повернутся не останавливаясь
		minDistRotate := ((speed * (180 / float64(rotateAngle))) / (2 * math.Pi)) * 2

		if dist > maxSpeed*5 {
			if int(maxSpeed)*10 != int(speed)*10 {

				if maxSpeed > speed {
					if len(path)%2 == 0 {
						speed += maxSpeed / 10
					}
				} else {
					speed -= maxSpeed / 10
				}

			} else {
				speed = maxSpeed
			}
		} else {
			if minSpeed < speed {
				if len(path)%2 == 0 {
					speed -= maxSpeed / 10
				}
			} else {
				speed = minSpeed
			}
		}

		radRotate := float64(rotate) * math.Pi / 180
		stopX := float64(speed) * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := float64(speed) * math.Sin(radRotate)

		for i := 0; i < rotateAngle; i++ { // т.к. за 1 учаток пути корпус может повернуться на много градусов тут этот for)
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

		possibleMove, q, r, front := CheckCollisionsOnStaticMap(int(forecastX+stopX), int(forecastY+stopY), rotate, mp, body, false)

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
			} else {
				// если мы поворачиваемся то наша скорость = 0
				speed = 0
			}
		}

		path = append(path, unit.PathUnit{X: int(forecastX), Y: int(forecastY), Rotate: rotate, Millisecond: 100,
			Q: forecastQ, R: forecastR, Speed: speed, Animate: true})
	}

	if len(path) > 1 {
		path[len(path)-1].Speed = 0 // на последней точке машина останавливается
	}

	return nil, path
}
