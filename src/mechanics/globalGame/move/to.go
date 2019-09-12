package move

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
)

func To(forecastX, forecastY, maxSpeed, minSpeed, speed, ToX, ToY float64, rotate, rotateAngle int, mp *_map.Map,
	ignoreObstacle bool, thoriumSlots map[int]*detail.ThoriumSlot, afterburner, gravity bool, body *detail.Body) (error, []unit.PathUnit) {

	path := make([]unit.PathUnit, 0)

	fullMax := maxSpeed

	diffRotate := 0 // разница между углом цели и носа корпуса
	dist := 900.0

	for {
		forecastQ := 0
		forecastR := 0
		// находим длинную вектора до цели
		dist = game_math.GetBetweenDist(int(forecastX), int(forecastY), int(ToX), int(ToY))
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

		// скорость * (180/угол поворота) = длинна полокружности (грабая модель)
		// длинны окружности получаем ее радиус r= длинна полокружности /2 пи
		// r*2 получаем минимальное растояние от бтр до обьекта к которому он может повернутся не останавливаясь
		minDistRotate := 10 + ((speed*(180/float64(rotateAngle)))/(2*math.Pi))*2
		// todo анализировать будующий путь - если нельзя проекхать по большой траектории из за препятсвий или
		// todo или радиус большой траектории слишком большой
		// todo то тормазить машину до скорости в которой юнит сможет проехать

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

		possibleMove, q, r := collisions.CheckCollisionsOnStaticMap(int(forecastX+stopX), int(forecastY+stopY), rotate, mp, body)

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
				// смотрим точку где мы моем доехать до упора и поворачиваемся
				forecastX, forecastY = DetailCheckStaticCollision(forecastX, forecastY, rotate, mp, body)
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

func DetailCheckStaticCollision(startX, startY float64, rotate int, mp *_map.Map, body *detail.Body) (float64, float64) {

	radRotate := float64(rotate) * math.Pi / 180

	oldX, oldY := 0.0, 0.0

	for {
		stopX := 1 * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := 1 * math.Sin(radRotate)

		possibleMove, _, _ := collisions.CheckCollisionsOnStaticMap(int(startX+stopX), int(startY+stopY), rotate, mp, body)
		if possibleMove {
			oldX = stopX
			oldY = stopY
		} else {
			return oldX, oldY
		}
	}
}
