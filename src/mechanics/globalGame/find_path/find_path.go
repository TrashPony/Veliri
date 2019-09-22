package find_path

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
)

//** SOURCE CODE https://github.com/JavaDar/aStar **//

type Points map[string]*coordinate.Coordinate

func MoveUnit(moveUnit *unit.Unit, ToX, ToY float64, mp *_map.Map, size int, uuid string) ([]*coordinate.Coordinate, error) {

	startX := moveUnit.X
	startY := moveUnit.Y

	allUnits := globalGame.Clients.GetAllShortUnits(mp.Id, true)

	start := &coordinate.Coordinate{X: startX, Y: startY}
	end := &coordinate.Coordinate{X: int(ToX), Y: int(ToY)}

	err, regions := FindRegionPath(mp, start, end, moveUnit, uuid)
	if err != nil {

		// если не удалось построить путь по регионам то делаем без регионов
		// todo однако это не правильно

		println("без регионов")
		err, path := FindPath(mp, start, end, moveUnit, size, allUnits, uuid, regions)
		return path, err
	} else {

		if debug.Store.RegionResult {
			for _, region := range regions {
				drawRegion(region, mp.Id, "blue")
			}
		}

		err, path := FindPath(mp, start, end, moveUnit, size, allUnits, uuid, regions)
		if err != nil {

			// если не удалось построить путь по регионам то делаем без регионов
			// todo однако это не правильно

			// координаты портятся поэтому заного создаем
			start := &coordinate.Coordinate{X: startX, Y: startY}
			end := &coordinate.Coordinate{X: int(ToX), Y: int(ToY)}

			err, path = FindPath(mp, start, end, moveUnit, size, allUnits, uuid, nil)
			println(len(path))
		}

		if debug.Store.AStartResult {
			for _, cell := range path {
				debug.Store.AddMessage("CreateRect", "green", cell.X, cell.Y, 0, 0, game_math.CellSize, mp.Id, 0)
			}
		}

		return path, err
	}
}

func PrepareInData(mp *_map.Map, start, end *coordinate.Coordinate, gameUnit *unit.Unit, scaleMap int) (*coordinate.Coordinate, *coordinate.Coordinate, int, int, error) {

	println(mp.QSize * game_math.HexagonWidth)
	println(int(float64(mp.RSize) * float64(game_math.HexagonHeight) * 0.75))

	xSize, ySize := mp.SetXYSize(game_math.HexagonWidth, game_math.HexagonHeight, scaleMap) // расчтиамем высоту и ширину карты в ху

	start.X, start.Y = start.X/scaleMap, start.Y/scaleMap
	start.Rotate = gameUnit.Rotate

	end.X, end.Y = end.X/scaleMap, end.Y/scaleMap

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
	allUnits map[int]*unit.ShortUnitInfo, uuid string, regions []*_map.Region) (error, []*coordinate.Coordinate) {

	start, end, xSize, ySize, err := PrepareInData(gameMap, start, end, gameUnit, scaleMap)

	openPoints, closePoints := Points{}, Points{} // создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints[start.Key()] = start               // кладем в карту посещенных точек стартовую точку

	if err != nil {
		return err, nil
	}

	var path []*coordinate.Coordinate
	var noSortedPath []*coordinate.Coordinate

	for uuid == gameUnit.MoveUUID {

		if len(openPoints) <= 0 {
			return errors.New("a star path no find"), nil
		}
		current := MinF(openPoints, xSize, ySize) // Берем точку с мин стоимостью пути

		if current.EqualXY(end) { // если текущая точка и есть конец начинаем генерить путь

			for !current.EqualXY(start) { // идем обратно до тех пока пока не дойдем до стартовой точки

				current = current.Parent     // по родительским точкам
				if !current.EqualXY(start) { // если текущая точка попрежнему не стартовая то добавляем в путь координату
					noSortedPath = append(noSortedPath, current)
				}
			}
			break
		}
		parseNeighbours(current, &openPoints, &closePoints, gameMap, end, gameUnit, xSize, ySize, scaleMap, allUnits, regions)
	}

	for i := len(noSortedPath); i > 0; i-- {
		noSortedPath[i-1].X *= scaleMap
		noSortedPath[i-1].Y *= scaleMap
		path = append(path, noSortedPath[i-1])
	}

	end.X *= scaleMap
	end.Y *= scaleMap

	if len(path) > 0 {
		end.Rotate = game_math.GetBetweenAngle(float64(end.X), float64(end.Y), float64(path[len(path)-1].X), float64(path[len(path)-1].Y))
	} else {
		start.X, start.Y = start.X*scaleMap, start.Y*scaleMap
		end.Rotate = game_math.GetBetweenAngle(float64(end.X), float64(end.Y), float64(start.X), float64(start.Y))
	}

	path = append(path, end)
	return nil, path
}

func parseNeighbours(curr *coordinate.Coordinate, open, close *Points, gameMap *_map.Map, end *coordinate.Coordinate,
	gameUnit *unit.Unit, xSize, ySize, scaleMap int, allUnits map[int]*unit.ShortUnitInfo, regions []*_map.Region) {

	delete(*open, curr.Key())   // удаляем ячейку из не посещенных
	(*close)[curr.Key()] = curr // добавляем в массив посещенные

	nCoordinate := generateNeighboursCoordinate(curr, gameMap, gameUnit, scaleMap, allUnits, xSize, ySize, regions) // берем всех соседей этой клетки

	for _, xLine := range nCoordinate {
		for _, c := range xLine {

			if (*close)[c.Key()] != nil || (*open)[c.Key()] != nil {
				continue // если ячейка является блокированой или находиться в масиве посещенных то пропускаем ее
			}

			// считаем для поинта значения пути
			c.G = curr.GetXYG(c) // стоимость клетки
			c.H = GetH(c, end)   // приближение от точки до конечной цели.
			c.F = c.GetF()       // длина пути до цели
			c.Parent = curr      //ref is needed?

			(*open)[c.Key()] = c // добавляем точку в масив не посещеных

			if debug.Store.AStartNeighbours {
				debug.Store.AddMessage("CreateRect", "orange", c.X*scaleMap, c.Y*scaleMap, 0, 0, scaleMap, gameMap.Id, 20)
			}
		}
	}
}

func GetH(a, b *coordinate.Coordinate) int { // эвристическое приближение стоимости пути от v до конечной цели.
	tmp := math.Abs(float64(a.X - b.X)) // вычисляем разницу между точкой и концом пути по Х
	tmp += math.Abs(float64(a.Y - b.Y)) // вычисляем разницу между точкой и концом пути по Y и сумируем с раницой по X

	return int(tmp)
}

func MinF(points Points, xSize, ySize int) (min *coordinate.Coordinate) { // берет точку с минимальной стоимостью пути из масива не посещеных
	min = &coordinate.Coordinate{F: xSize*ySize + 1}

	for _, p := range points {
		if p.F < min.F {
			min = p
		}
	}

	return
}
