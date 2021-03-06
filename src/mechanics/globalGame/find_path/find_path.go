package find_path

import (
	"errors"
	"fmt"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
	"strconv"
	"sync"
	"time"
)

//** SOURCE CODE https://github.com/JavaDar/aStar **//

type Points struct {
	points map[string]*coordinate.Coordinate
	mx     sync.Mutex
}

func MoveUnit(moveUnit *unit.Unit, ToX, ToY float64, mp *_map.Map, size int, uuid string,
	units map[int]*unit.ShortUnitInfo, unitsID []int) ([]*coordinate.Coordinate, error) {

	startX := moveUnit.X
	startY := moveUnit.Y

	start := &coordinate.Coordinate{X: startX, Y: startY}
	end := &coordinate.Coordinate{X: int(ToX), Y: int(ToY)}

	// чисто А стар
	// err, path := FindPath(mp, start, end, moveUnit, size, allUnits, uuid, nil)
	// return path, err

	err, regions := FindRegionPath(mp, start, end, moveUnit, uuid)
	if err != nil {

		// если не удалось построить путь по регионам то делаем без регионов
		// todo однако это не правильно

		err, path := FindPath(mp, start, end, moveUnit, size, units, uuid, regions, unitsID)
		return path, err
	} else {

		if debug.Store.RegionResult {
			for _, region := range regions {
				drawRegion(region, mp.Id, "blue")
			}
		}

		err, path := FindPath(mp, start, end, moveUnit, size, units, uuid, regions, unitsID)
		if err != nil {

			// если не удалось построить путь по регионам то делаем без регионов
			// todo однако это не правильно

			// координаты портятся поэтому заного создаем
			start := &coordinate.Coordinate{X: startX, Y: startY}
			end := &coordinate.Coordinate{X: int(ToX), Y: int(ToY)}

			err, path = FindPath(mp, start, end, moveUnit, size, units, uuid, nil, unitsID)
		}

		if debug.Store.AStartResult {
			for _, cell := range path {
				debug.Store.AddMessage("CreateRect", "green", cell.X, cell.Y, 0, 0, game_math.CellSize, mp.Id, 0)
			}
		}

		return path, err
	}
}

func PrepareInData(mp *_map.Map, start, end *coordinate.Coordinate, gameUnit *unit.Unit, scaleMap int,
	regions []*_map.Region, units map[int]*unit.ShortUnitInfo, unitsID []int) (*coordinate.Coordinate, *coordinate.Coordinate, int, int, error) {

	xSize, ySize := mp.SetXYSize(scaleMap) // расчтиамем высоту и ширину карты в ху

	start.X, start.Y = start.X/scaleMap, start.Y/scaleMap
	start.Rotate = gameUnit.Rotate // todo возможно не нужно
	start.State = 1

	end.X, end.Y = end.X/scaleMap, end.Y/scaleMap

	// если конечная точка является невалидной то ищем ближайшую валидную точку и говорим что это цель
	_, free, _ := checkValidForMoveCoordinate(mp, end.X, end.Y, xSize, ySize, gameUnit, scaleMap, regions, units, unitsID)
	if !free {
		getValidEnd := func() *coordinate.Coordinate {
			for radius := 0; radius < 10; radius++ {
				for angle := 10; angle < 360; angle += 10 {

					x := int(math.Round(float64(float64(end.X) + float64(radius)*math.Cos(float64(angle)))))
					y := int(math.Round(float64(float64(end.Y) + float64(radius)*math.Sin(float64(angle)))))

					_, free, _ := checkValidForMoveCoordinate(mp, x, y, xSize, ySize, gameUnit, scaleMap, regions, units, unitsID)
					if free {
						return &coordinate.Coordinate{X: x, Y: y}
					}
				}
			}
			return nil
		}
		end = getValidEnd()
	}

	if end.X >= xSize || end.Y >= ySize || end.X < 0 || end.Y < 0 || start.X >= xSize || start.Y >= ySize || start.X < 0 || start.Y < 0 {
		return nil, nil, 0, 0, errors.New("start or end out the range")
	}

	if end == nil {
		return nil, nil, 0, 0, errors.New("end point not valid")
	} else {
		return start, end, xSize, ySize, nil
	}
}

func FindPath(gameMap *_map.Map, start, end *coordinate.Coordinate, gameUnit *unit.Unit, scaleMap int,
	units map[int]*unit.ShortUnitInfo, uuid string, regions []*_map.Region, unitsID []int) (error, []*coordinate.Coordinate) {

	startTime := time.Now()
	defer func() {
		if debug.Store.Move {
			elapsed := time.Since(startTime)
			fmt.Println("time aStar path: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64))
		}
	}()

	start, end, xSize, ySize, err := PrepareInData(gameMap, start, end, gameUnit, scaleMap, regions, units, unitsID)
	if debug.Store.AStartNeighbours {
		debug.Store.AddMessage("CreateRect", "blue", end.X*scaleMap, end.Y*scaleMap, 0, 0, scaleMap, gameMap.Id, 0)
	}

	// создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints, closePoints := Points{points: make(map[string]*coordinate.Coordinate)}, Points{points: make(map[string]*coordinate.Coordinate)}
	openPoints.points[start.Key()] = start // кладем в карту посещенных точек стартовую точку

	if err != nil {
		return err, nil
	}

	var path []*coordinate.Coordinate
	var noSortedPath []*coordinate.Coordinate

	exit := false
	countWorker := 0

	for uuid == gameUnit.MoveUUID {

		if exit {
			break
		}

		// 32 идиальное колво воркеров, подобрано эксперементально, скорее всего зависит от камня (fx8350)
		if len(openPoints.points) <= 0 || countWorker > 32 {
			time.Sleep(time.Millisecond)

			if len(openPoints.points) <= 0 && countWorker == 0 {
				return errors.New("a star path no find"), nil
			}

			continue
		}

		current, find := MinF(&openPoints, &closePoints, xSize, ySize) // Берем точку с мин стоимостью пути
		if !find || exit {
			continue
		}
		countWorker++

		go func() {
			defer func() {
				countWorker--
			}()

			if current.Equal(end) { // если текущая точка и есть конец начинаем генерить путь

				for !current.Equal(start) { // идем обратно до тех пока пока не дойдем до стартовой точки

					current = current.Parent // по родительским точкам

					if !current.Equal(start) { // если текущая точка попрежнему не стартовая то добавляем в путь координату
						noSortedPath = append(noSortedPath, current)
					}
				}
				exit = true
				return
			}

			if !exit {
				parseNeighbours(current, &openPoints, &closePoints, gameMap, end, gameUnit, xSize, ySize, scaleMap, units, regions, unitsID)
			}
		}()

		if exit {
			break
		}
	}

	for i := len(noSortedPath); i > 0; i-- {
		noSortedPath[i-1].X *= scaleMap
		noSortedPath[i-1].Y *= scaleMap
		path = append(path, noSortedPath[i-1])
	}

	end.X *= scaleMap
	end.Y *= scaleMap

	if len(path) > 0 { // todo возможно не нужно
		end.Rotate = game_math.GetBetweenAngle(float64(end.X), float64(end.Y), float64(path[len(path)-1].X), float64(path[len(path)-1].Y))
	} else {
		start.X, start.Y = start.X*scaleMap, start.Y*scaleMap
		end.Rotate = game_math.GetBetweenAngle(float64(end.X), float64(end.Y), float64(start.X), float64(start.Y))
	}

	path = append(path, end)
	return nil, path
}

func parseNeighbours(curr *coordinate.Coordinate, open, close *Points, gameMap *_map.Map, end *coordinate.Coordinate,
	gameUnit *unit.Unit, xSize, ySize, scaleMap int, units map[int]*unit.ShortUnitInfo, regions []*_map.Region, unitsID []int) {

	nCoordinate := generateNeighboursCoordinate(curr, gameMap, gameUnit, scaleMap, units, xSize, ySize, regions, unitsID) // берем всех соседей этой клетки

	open.mx.Lock()
	defer open.mx.Unlock()

	close.mx.Lock()
	defer close.mx.Unlock()

	for _, xLine := range nCoordinate {
		for _, c := range xLine {

			if close.points[c.Key()] != nil || open.points[c.Key()] != nil {
				continue // если ячейка является блокированой или находиться в масиве посещенных то пропускаем ее
			}

			// считаем для поинта значения пути
			c.G = curr.GetXYG(c) // стоимость клетки
			c.H = GetH(c, end)   // приближение от точки до конечной цели.
			c.F = c.GetF()       // длина пути до цели
			c.Parent = curr      //ref is needed?

			open.points[c.Key()] = c // добавляем точку в масив не посещеных

			if debug.Store.AStartNeighbours {
				debug.Store.AddMessage("CreateRect", "orange", c.X*scaleMap, c.Y*scaleMap, 0, 0, scaleMap, gameMap.Id, 0)
			}
		}
	}
}

func GetH(a, b *coordinate.Coordinate) int { // эвристическое приближение стоимости пути от v до конечной цели.
	tmp := math.Abs(float64(a.X - b.X)) // вычисляем разницу между точкой и концом пути по Х
	tmp += math.Abs(float64(a.Y - b.Y)) // вычисляем разницу между точкой и концом пути по Y и сумируем с раницой по X

	return int(tmp)
}

func MinF(open, close *Points, xSize, ySize int) (min *coordinate.Coordinate, find bool) { // берет точку с минимальной стоимостью пути из масива не посещеных
	min = &coordinate.Coordinate{F: xSize*ySize + 1}

	open.mx.Lock()
	defer open.mx.Unlock()

	close.mx.Lock()
	defer close.mx.Unlock()

	find = false
	for _, p := range open.points {
		if p.F < min.F {
			min = p
			find = true
		}
	}

	if find {
		delete(open.points, min.Key()) // удаляем ячейку из не посещенных
		close.points[min.Key()] = min  // добавляем в массив посещенные
	}

	return
}
