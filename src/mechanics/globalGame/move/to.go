package move

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
)

func To(forecastX, forecastY, maxSpeed, minSpeed, speed, ToX, ToY float64, rotate, rotateAngle int) (error, []*unit.PathUnit) {

	path := make([]*unit.PathUnit, 0)

	//// скорость * (180/угол поворота) = длинна полокружности (грабая модель)
	//// длинны окружности получаем ее радиус r= длинна полокружности /2 пи
	//// r*2 получаем минимальное растояние от бтр до обьекта к которому он может повернутся не останавливаясь
	//// todo анализировать будующий путь - если нельзя проекхать по большой траектории из за препятсвий или
	//// todo или радиус большой траектории слишком большой
	//// todo то тормазить машину до скорости в которой юнит сможет проехать
	//minDistRotate := 10 + ((maxSpeed*(180/float64(rotateAngle)))/(2*math.Pi))*2

	for {

		// находим длинную вектора до цели
		dist := game_math.GetBetweenDist(int(forecastX), int(forecastY), int(ToX), int(ToY))
		if dist < maxSpeed+5 {
			break
		}

		percentsSpeed := (speed * 100) / maxSpeed
		if dist > (maxSpeed*10)/percentsSpeed {
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

		//math.Atan2 куда у - текущие у, куда х - текущие х, получаем угол
		// находим какой угол необходимо принять телу
		needRad := math.Atan2(ToY-forecastY, ToX-forecastX)
		needRotate := int(needRad * 180 / 3.14)
		// и принимаем его
		RotateUnit(&rotate, &needRotate, rotateAngle)

		forecastX = forecastX + stopX
		forecastY = forecastY + stopY

		path = append(path, &unit.PathUnit{X: int(forecastX), Y: int(forecastY), Rotate: rotate, Millisecond: 100, Speed: speed, Animate: true})
	}

	if len(path) > 1 {
		path[len(path)-1].Speed = 0 // на последней точке машина останавливается
	}

	return nil, path
}

//func DetailCheckStaticCollision(startX, startY float64, rotate int, mp *_map.Map, body *detail.Body) (float64, float64) {
//
//	radRotate := float64(rotate) * math.Pi / 180
//
//	oldX, oldY := 0.0, 0.0
//
//	for {
//		stopX := 1 * math.Cos(radRotate) // идем по вектору движения корпуса
//		stopY := 1 * math.Sin(radRotate)
//
//		possibleMove, _, _, _:= collisions.CheckCollisionsOnStaticMap(int(startX+stopX), int(startY+stopY), rotate, mp, body)
//		if possibleMove {
//			oldX = stopX
//			oldY = stopY
//		} else {
//			return oldX, oldY
//		}
//	}
//}
