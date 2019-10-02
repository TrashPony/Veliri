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

func LeftHandAlgorithm(moveUnit *unit.Unit, startX, startY, ToX, ToY float64, uuid string, units map[int]*unit.ShortUnitInfo) ([]*coordinate.Coordinate, error) {

	mp, _ := maps.Maps.GetByID(moveUnit.MapID)

	// 0 пытаемя проложить путь от начала пути до конечной точки по прямой
	collision := collisions.SearchCollisionInLine(startX, startY, ToX, ToY, mp, moveUnit, 5, units, false)
	if !collision {
		return []*coordinate.Coordinate{{X: int(moveUnit.ToX), Y: int(moveUnit.ToY)}}, nil
	} else {

		// конец пути находится в препятсвие
		possibleMove, _ := collisions.BodyCheckCollisionsOnStaticMap(int(ToX), int(ToY), 0, mp, moveUnit.Body, false, true)
		free, _ := collisions.CheckCollisionsPlayers(moveUnit, int(ToX), int(ToY), 0, units, false, true, false)
		// если конечная точка находится в препятсвие то смотрим куда ближе идти ко входу или к выходу
		if !possibleMove || !free {
			ToX, ToY, collision = SearchEndPoint(startX, startY, ToX, ToY, moveUnit, mp, units)
			moveUnit.ToX, moveUnit.ToY = ToX, ToY
		}

		if !collision {
			return []*coordinate.Coordinate{{X: int(ToX), Y: int(ToY)}}, nil
		} else {
			return startFind(moveUnit, int(startX), int(startY), ToX, ToY, uuid, game_math.CellSize, mp, units)
		}
	}
}

// метод пускает от точки ToX, ToY 8 лузей по разным направлениям, что бы найти первое ближайшее место куда может встать юнит
func SearchEndPoint(startX, startY, ToX, ToY float64, moveUnit *unit.Unit, mp *_map.Map, units map[int]*unit.ShortUnitInfo) (float64, float64, bool) {

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

	search := func(x, y, steep float64, angle int) (float64, float64, bool, bool) {

		radian := float64(angle) * math.Pi / 180
		stopX, stopY := float64(game_math.CellSize*steep)*math.Cos(radian)+x, float64(game_math.CellSize*steep)*math.Sin(radian)+y

		possibleMove, _ := collisions.BodyCheckCollisionsOnStaticMap(int(stopX), int(stopY), 0, mp, moveUnit.Body, false, true)
		free, _ := collisions.CheckCollisionsPlayers(moveUnit, int(stopX), int(stopY), 0, units, false, true, true)
		if possibleMove && free {

			if debug.Store.MoveEndPoint {
				debug.Store.AddMessage("CreateRect", "green", int(stopX), int(stopY), 0, 0, game_math.CellSize, mp.Id, 20)
			}

			return stopX, stopY,
				collisions.SearchCollisionInLine(startX, startY, stopX, stopY, mp, moveUnit, game_math.CellSize, units, false), true
		}

		if debug.Store.MoveEndPoint {
			debug.Store.AddMessage("CreateRect", "white", int(stopX), int(stopY), 0, 0, game_math.CellSize, mp.Id, 20)
		}

		return 0, 0, false, false
	}

	step := 0
	for {
		for angle := 0; angle <= 360; angle += 45 {
			x, y, collision, passed := search(ToX, ToY, float64(step), angle)
			if passed {
				return x, y, collision
			}
		}
		step++
	}
}

func startFind(moveUnit *unit.Unit, x, y int, ToX, ToY float64, uuid string, size int, mp *_map.Map, units map[int]*unit.ShortUnitInfo) ([]*coordinate.Coordinate, error) {

	path := make([]*coordinate.Coordinate, 0)
	last := false
	tryCount := 3

	var points []*coordinate.Coordinate

	for moveUnit.MoveUUID == uuid {

		if !last {

			// ищем путь алгоритмом А*
			if points == nil {
				// т.к. он не будет менятся нам нет смысла искать его всегда заного
				var err error
				points, err = MoveUnit(moveUnit, ToX, ToY, mp, size, uuid, units)
				if err != nil {

					collision := false
					ToX, ToY, collision = SearchEndPoint(float64(x), float64(y), ToX, ToY, moveUnit, mp, units)
					moveUnit.ToX, moveUnit.ToY = ToX, ToY
					points = nil
					tryCount--

					if !collision || tryCount == 0 {
						return nil, err
					}
				}
			}

			// находим максимальную отдаленную точку куда может попать юнит
			x, y, last = SearchPoint(&points, x, y, mp, moveUnit, float64(size), units)
			if x == 0 && y == 0 {
				points = nil
				tryCount--
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

func SearchPoint(points *[]*coordinate.Coordinate, unitX, unitY int, mp *_map.Map, gameUnit *unit.Unit, size float64, units map[int]*unit.ShortUnitInfo) (int, int, bool) {

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

	if len(*points) > 0 {
		for i := len(*points) - 1; i >= 0; i-- {

			go func(index int, cell *coordinate.Coordinate) {

				defer func() {
					countClose++
				}()

				collision := collisions.SearchCollisionInLine(
					float64(cell.X)+size/2,
					float64(cell.Y)+size/2,
					float64(unitX),
					float64(unitY), mp, gameUnit, size, units, false)
				if !collision {
					if index > lastIndex {
						x, y, lastIndex = cell.X+int(size)/2, cell.Y+int(size)/2, index
					}
				}
			}(i, (*points)[i])
		}
	}

	for countClose < len(*points)-1 {
		time.Sleep(time.Millisecond)
	}

	return x, y, len(*points)-1 == lastIndex
}
