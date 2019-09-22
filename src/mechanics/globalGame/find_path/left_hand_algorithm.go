package find_path

import (
	"errors"
	"fmt"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"strconv"
	"time"
)

func LeftHandAlgorithm(moveUnit *unit.Unit, startX, startY, ToX, ToY float64, uuid string) ([]*coordinate.Coordinate, error) {

	mp, _ := maps.Maps.GetByID(moveUnit.MapID)

	// 0 пытаемя проложить путь от начала пути до конечной точки по прямой
	collision := collisions.SearchCollisionInLine(startX, startY, ToX, ToY, mp, moveUnit.Body, game_math.CellSize)
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
			return startFind(moveUnit, int(startX), int(startY), ToX, ToY, uuid, game_math.CellSize, mp)
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
			x, y = SearchPoint(&points, x, y, mp, moveUnit, float64(size))
			if x == 0 && y == 0 {
				return nil, errors.New("line not to cell")
			}

			if debug.Store.HandAlgorithm && len(path) > 0 {
				debug.Store.AddMessage("CreateLine", "red", path[len(path)-1].X,
					path[len(path)-1].Y, x, y, size, mp.Id, 20)
			}

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

func SearchPoint(points *[]*coordinate.Coordinate, unitX, unitY int, mp *_map.Map, gameUnit *unit.Unit, size float64) (int, int) {

	startTime := time.Now()
	defer func() {
		if debug.Store.Move {
			elapsed := time.Since(startTime)
			fmt.Println("time search line: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64))
		}
	}()

	// todo самый дорогой метод на дальних дистациях из за того что он считается много раз, возможно можно просчитать его в 1 фор
	// ищем самую дальнюю точку до которой можем дойти

	x, y := 0, 0
	lastIndex := 0
	countClose := 0

	for i := len(*points) - 1; i >= 0; i-- {
		go func(index int) {

			defer func() {
				countClose++
			}()

			collision := collisions.SearchCollisionInLine(
				float64((*points)[index].X)+size/2,
				float64((*points)[index].Y)+size/2,
				float64(unitX),
				float64(unitY), mp, gameUnit.Body, size)
			if !collision {
				if index > lastIndex {
					x, y, lastIndex = (*points)[index].X+int(size)/2, (*points)[index].Y+int(size)/2, index
				}
			}
		}(i)
	}

	for countClose < len(*points) {
		time.Sleep(time.Millisecond)
	}

	return x, y
}
