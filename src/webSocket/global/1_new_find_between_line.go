package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
)

// возвращает xКолизии, yКолизии, хВыхода из колизии, yВыхода из колизии прошли без столкновений
func BetweenLine(startX, startY, ToX, ToY float64, mp *_map.Map, body *detail.Body, startMove bool) (
	entryPoints, collisionPoints, outPoints []*coordinate.Coordinate, collision, endIsObstacle bool) {

	// TODO учитывать угол поворота юнита
	// идем по линии со скорость 10 рх
	speed := 10.0

	// угол от старта до конца
	angle := game_math.GetBetweenAngle(ToX, ToY, startX, startY)
	radian := float64(angle) * math.Pi / 180

	// текущее положение курсора
	currentX, currentY := startX, startY

	// точки входа в колизию
	entryPoints = make([]*coordinate.Coordinate, 0)

	// точки колизии
	collisionPoints = make([]*coordinate.Coordinate, 0)

	// точки выхода из колизии
	outPoints = make([]*coordinate.Coordinate, 0)

	currentCollision := false

	extraExit := false
	for {
		// находим длинную вектора до цели
		distToEnd := game_math.GetBetweenDist(int(currentX), int(currentY), int(ToX), int(ToY))
		//distToStart := game_math.GetBetweenDist(int(currentX), int(currentY), int(startX), int(startY))

		if distToEnd < speed+15 || extraExit {
			if currentCollision {
				endIsObstacle = true
				extraExit = true
			} else {
				if extraExit {
					// добавляем последнюю точку что бы найти по ней ближайшую
					outPoints = append(outPoints, &coordinate.Coordinate{X: int(currentX), Y: int(currentY)})
				}
				return entryPoints, outPoints, collisionPoints, collision, endIsObstacle
			}
		}

		stopX, stopY := float64(speed)*math.Cos(radian), float64(speed)*math.Sin(radian)

		possibleMove, _, _, _ := collisions.CheckCollisionsOnStaticMap(int(currentX), int(currentY), angle, mp, body, true)
		if !possibleMove {
			// если юнит по каким то причинам стартует из колизии то дать ему выйти и потом уже искать колизию
			//if !(distToStart < speed+2 && startMove) {
			// фиксируем 1 точку колизии
			if !currentCollision {

				entryPoints = append(entryPoints, &coordinate.Coordinate{X: int(currentX), Y: int(currentY)})
				collisionPoints = append(collisionPoints, &coordinate.Coordinate{X: int(currentX + stopX), Y: int(currentY + stopY)})

				collision = true
				currentCollision = true
			}
			//}
		} else {
			// фиксируем точку выхода из колизии
			if currentCollision {
				currentCollision = false
				outPoints = append(outPoints, &coordinate.Coordinate{X: int(currentX), Y: int(currentY)})
			}
		}

		currentX += stopX
		currentY += stopY
	}
}
