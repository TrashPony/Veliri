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
	"math"
	"strconv"
	"time"
)

func LeftHandAlgorithm(moveUnit *unit.Unit, startX, startY, ToX, ToY float64, uuid string) ([]*coordinate.Coordinate, error) {

	mp, _ := maps.Maps.GetByID(moveUnit.MapID)

	// 0 пытаемя проложить путь от начала пути до конечной точки по прямой
	collision := collisions.SearchCollisionInLine(startX, startY, ToX, ToY, mp, moveUnit.Body, 5)
	if !collision {
		return []*coordinate.Coordinate{{X: int(moveUnit.ToX), Y: int(moveUnit.ToY)}}, nil
	} else {
		// 0.3 на прямой были найдены препятвия
		possibleMove, _ := collisions.CheckCollisionsOnStaticMap(int(ToX), int(ToY), 0, mp, moveUnit.Body, false, true)
		// если конечная точка находится в препятсвие то смотрим куда ближе идти ко входу или к выходу
		if !possibleMove {
			ToX, ToY, collision = SearchEndPoint(startX, startY, ToX, ToY, moveUnit, mp)
			moveUnit.ToX, moveUnit.ToY = ToX, ToY
		}

		if !collision {
			return []*coordinate.Coordinate{{X: int(ToX), Y: int(ToY)}}, nil
		} else {
			return startFind(moveUnit, int(startX), int(startY), ToX, ToY, uuid, game_math.CellSize, mp)
		}
	}
}

func SearchEndPoint(startX, startY, ToX, ToY float64, moveUnit *unit.Unit, mp *_map.Map) (float64, float64, bool) {

	startTime := time.Now()
	defer func() {
		if debug.Store.Move {
			elapsed := time.Since(startTime)
			fmt.Println("time search end point: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64))
		}
	}()

	if debug.Store.MoveEndPoint {
		debug.Store.AddMessage("CreateRect", "red", int(ToX), int(ToY), 0, 0, 5, mp.Id, 20)
	}

	toTmpX, toTmpY := ToX, ToY
	toTmpX2, toTmpY2 := ToX, ToY

	angle := game_math.GetBetweenAngle(ToX, ToY, startX, startY)
	radian := float64(angle) * math.Pi / 180

	// идем в обе стороны по направлению вектра пока не найдем пригодную точку выхода
	for {
		stopX, stopY := float64(game_math.CellSize)*math.Cos(radian), float64(game_math.CellSize)*math.Sin(radian)
		stopX2, stopY2 := float64(-game_math.CellSize)*math.Cos(radian), float64(-game_math.CellSize)*math.Sin(radian)

		toTmpX, toTmpY = toTmpX+stopX, toTmpY+stopY
		toTmpX2, toTmpY2 = toTmpX2+stopX2, toTmpY2+stopY2

		possibleMove1, _ := collisions.CheckCollisionsOnStaticMap(int(toTmpX), int(toTmpY), 0, mp, moveUnit.Body, false, true)
		if possibleMove1 {

			if debug.Store.MoveEndPoint {
				debug.Store.AddMessage("CreateRect", "green", int(toTmpX), int(toTmpY), 0, 0, game_math.CellSize, mp.Id, 20)
			}

			return toTmpX, toTmpY, collisions.SearchCollisionInLine(startX, startY, toTmpX, toTmpY, mp, moveUnit.Body, game_math.CellSize)
		}
		possibleMove2, _ := collisions.CheckCollisionsOnStaticMap(int(toTmpX2), int(toTmpY2), 0, mp, moveUnit.Body, false, true)
		if possibleMove2 {

			if debug.Store.MoveEndPoint {
				debug.Store.AddMessage("CreateRect", "green", int(toTmpX2), int(toTmpY2), 0, 0, game_math.CellSize, mp.Id, 20)
			}

			return toTmpX2, toTmpY2, collisions.SearchCollisionInLine(startX, startY, toTmpX2, toTmpY2, mp, moveUnit.Body, game_math.CellSize)
		}
	}
}

func startFind(moveUnit *unit.Unit, x, y int, ToX, ToY float64, uuid string, size int, mp *_map.Map) ([]*coordinate.Coordinate, error) {

	path := make([]*coordinate.Coordinate, 0)
	last := false
	var points []*coordinate.Coordinate

	for moveUnit.MoveUUID == uuid {

		if !last {

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
			x, y, last = SearchPoint(&points, x, y, mp, moveUnit, float64(size))
			if x == 0 && y == 0 {
				return nil, errors.New("line not to cell")
			}

			if last {
				// если последняя точка то не добавляем ее тут а прокидываем в конец
				continue
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

func SearchPoint(points *[]*coordinate.Coordinate, unitX, unitY int, mp *_map.Map, gameUnit *unit.Unit, size float64) (int, int, bool) {

	startTime := time.Now()
	defer func() {
		if debug.Store.Move {
			elapsed := time.Since(startTime)
			fmt.Println("time search line: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64))
		}
	}()

	// todo самый дорогой метод на дальних дистациях из за того что он считается много раз, возможно можно просчитать его в 1 фор
	// 	но у меня чет не вышло
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

	for countClose < len(*points)-1 {
		time.Sleep(time.Millisecond)
	}

	return x, y, len(*points)-1 == lastIndex
}
