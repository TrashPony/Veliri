package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func LeftHandAlgorithm(moveUnit *unit.Unit, startX, startY, ToX, ToY float64, uuid string) ([]*coordinate.Coordinate, error) {

	mp, _ := maps.Maps.GetByID(moveUnit.MapID)
	rectSize := 30

	// 0 пытаемя проложить путь от начала пути до конечной точки по прямой
	collision := collisions.SearchCollisionInLine(startX, startY, ToX, ToY, mp, moveUnit.Body)
	if !collision {
		return []*coordinate.Coordinate{{X: int(moveUnit.ToX), Y: int(moveUnit.ToY)}}, nil
	} else {
		// 0.3 на прямой были найдены препятвия
		// size 5 потому что из за большой скорость можнео не заметить препятвия на конце линии, важно для endIsObstacle
		entryPoint, outPoint, _, collision, endIsObstacle := collisions.BetweenLine(startX, startY, ToX, ToY, mp, moveUnit.Body, true, 5)

		// 0.1 если конечная точка находится в препятсвие то смотрим куда ближе идти ко входу или к выходу
		if endIsObstacle {
			// последние точки это колизия вокруг точки назначения
			lastEntryX, lastEntryY := &entryPoint[len(entryPoint)-1].X, &entryPoint[len(entryPoint)-1].Y
			lastOutX, lastOutY := &outPoint[len(outPoint)-1].X, &outPoint[len(outPoint)-1].Y

			// ищем ближайшую точку которая не в колизии
			EndIsObstacle(&ToX, &ToY, lastEntryX, lastEntryY, lastOutX, lastOutY, &collision, moveUnit, len(entryPoint))
		}

		if !collision {
			return []*coordinate.Coordinate{{X: int(ToX), Y: int(ToY)}}, nil
		} else {
			return startFind(moveUnit, int(startX), int(startY), ToX, ToY, uuid, rectSize, mp)
		}
	}
}

func startFind(moveUnit *unit.Unit, x, y int, ToX, ToY float64, uuid string, size int, mp *_map.Map) ([]*coordinate.Coordinate, error) {

	path := make([]*coordinate.Coordinate, 0)

	var points []*coordinate.Coordinate

	for moveUnit.MoveUUID == uuid {

		_, _, _, collision, _ := collisions.BetweenLine(float64(x), float64(y), ToX, ToY, mp, moveUnit.Body, false, size)

		if collision {

			// ищем путь алгоритмом А*
			if points == nil {
				// т.к. он не будет менятся нам нет смысла искать его всегда заного
				var err error
				points, err = MoveUnit(moveUnit, ToX, ToY, mp, size, uuid)
				if err != nil {
					return nil, err
				}
			}
			// находим максимальную отдаленную точку куда может попать юнит
			x, y = SearchPoint(&points, x, y, mp, moveUnit)
			path = append(path, &coordinate.Coordinate{X: x, Y: y})
		} else {
			//  2.1.1 если между координатой истиного пути и целью нет препятсвий формируем путь. Выходим из функции.
			path = append(path, &coordinate.Coordinate{X: int(moveUnit.ToX), Y: int(moveUnit.ToY)})
			break
		}
	}

	return path, nil
}

func EndIsObstacle(ToX, ToY *float64, lastEntryX, lastEntryY, lastOutX, lastOutY *int, collision *bool, moveUnit *unit.Unit, countCollision int) {

	collisionStartDist := game_math.GetBetweenDist(int(*ToX), int(*ToY), *lastEntryX, *lastEntryY)
	collisionEndDist := game_math.GetBetweenDist(int(*ToX), int(*ToY), *lastOutX, *lastOutY)

	// если то старта колизии ближе чем до конца то считаем что маршрут без колизий
	if collisionStartDist < collisionEndDist {
		*ToX, *ToY = float64(*lastEntryX), float64(*lastEntryY)
		moveUnit.ToX, moveUnit.ToY = float64(*lastEntryX), float64(*lastEntryY)

		// говорим что нет колизий если она всего одна
		if countCollision == 1 {
			*collision = false
		}

	} else {
		// иначе переназначаем конечный пункт что бы не искать путь вечно
		*ToX, *ToY = float64(*lastOutX), float64(*lastOutY)
		moveUnit.ToX, moveUnit.ToY = float64(*lastOutX), float64(*lastOutY)
	}
}

func SearchPoint(points *[]*coordinate.Coordinate, unitX, unitY int, mp *_map.Map, gameUnit *unit.Unit) (int, int) {
	// ищем самую дальнюю точку до которой можем дойти
	for i := len(*points) - 1; i >= 0; i-- {
		collision := collisions.SearchCollisionInLine(float64((*points)[i].X), float64((*points)[i].Y), float64(unitX), float64(unitY), mp, gameUnit.Body)
		if !collision {
			return (*points)[i].X, (*points)[i].Y
		}
	}
	return 0, 0
}
