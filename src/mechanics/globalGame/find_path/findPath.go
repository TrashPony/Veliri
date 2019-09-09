package find_path

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"math"
)

//** SOURCE CODE https://github.com/JavaDar/aStar **//

const (
	FREE = iota
	BLOCKED
	START
	END
	PATH
)

type Points map[string]*coordinate.Coordinate

func MoveUnit(moveUnit *unit.Unit, ToX, ToY float64, mp *_map.Map) ([]unit.PathUnit, error) {

	startX := moveUnit.X
	startY := moveUnit.Y

	path := make([]unit.PathUnit, 0)

	allUnits := globalGame.Clients.GetAllShortUnits(mp.Id, true)
	_, path2 := FindPath(mp, &coordinate.Coordinate{X: startX, Y: startY},
		&coordinate.Coordinate{X: int(ToX), Y: int(ToY)}, moveUnit, moveUnit.Speed, allUnits)

	for _, unitPath := range path2 {
		path = append(path, unit.PathUnit{X: unitPath.X, Y: unitPath.Y, Rotate: unitPath.Rotate, Millisecond: 1000,
			Speed: float64(moveUnit.Speed), Animate: true})
	}

	return path, nil
}

func FindPath(gameMap *_map.Map, start, end *coordinate.Coordinate, gameUnit *unit.Unit, scaleMap int, allUnits map[int]*unit.ShortUnitInfo) (error, []*coordinate.Coordinate) {

	xSize, ySize := gameMap.SetXYSize(globalGame.HexagonWidth, globalGame.HexagonHeight, scaleMap) // расчтиамем высоту и ширину карты в ху

	start.X, start.Y = start.X/scaleMap, start.Y/scaleMap
	start.Rotate = gameUnit.Rotate

	end.X, end.Y = end.X/scaleMap, end.Y/scaleMap

	if end.X >= xSize || end.Y >= ySize || end.X < 0 || end.Y < 0 || start.X >= xSize || start.Y >= ySize || start.X < 0 || start.Y < 0 {
		return errors.New("no path"), nil
	}

	openPoints, closePoints := Points{}, Points{} // создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints[start.Key()] = start               // кладем в карту посещенных точек стартовую точку

	var path []*coordinate.Coordinate
	var noSortedPath []*coordinate.Coordinate

	for {
		if len(openPoints) <= 0 {
			return errors.New("no path"), nil
		}
		current := MinF(openPoints, xSize, ySize) // Берем точку с мин стоимостью пути
		if current.EqualXY(end) {                 // если текущая точка и есть конец начинаем генерить путь
			for !current.EqualXY(start) { // идем обратно до тех пока пока не дойдем до стартовой точки
				current = current.Parent     // по родительским точкам
				if !current.EqualXY(start) { // если текущая точка попрежнему не стартовая то
					noSortedPath = append(noSortedPath, current)
				}
			}
			break
		}
		parseNeighbours(current, &openPoints, &closePoints, gameMap, end, gameUnit, xSize, ySize, scaleMap, allUnits)
	}

	for i := len(noSortedPath); i > 0; i-- {
		noSortedPath[i-1].X *= scaleMap
		noSortedPath[i-1].Y *= scaleMap
		path = append(path, noSortedPath[i-1])
	}

	end.X *= scaleMap
	end.Y *= scaleMap
	path = append(path, end)
	return nil, path
}

func parseNeighbours(curr *coordinate.Coordinate, open, close *Points, gameMap *_map.Map,
	end *coordinate.Coordinate, gameUnit *unit.Unit, xSize, ySize, scaleMap int, allUnits map[int]*unit.ShortUnitInfo) {

	delete(*open, curr.Key())   // удаляем ячейку из не посещенных
	(*close)[curr.Key()] = curr // добавляем в массив посещенные

	nCoordinate := generateNeighboursCoordinate(curr, gameMap, gameUnit, scaleMap, allUnits, end) // берем всех соседей этой клетки

	for _, xLine := range nCoordinate {
		for _, c := range xLine {

			if c.X < xSize && c.Y < ySize && c.X > 0 && c.Y > 0 {
				if (*close)[c.Key()] != nil || (*open)[c.Key()] != nil {
					continue // если ячейка является блокированой или находиться в масиве посещенных то пропускаем ее
				}
				// считаем для поинта значения пути
				c.G = curr.GetXYG(c) // стоимость клетки
				c.H = GetH(c, end)   // приближение от точки до конечной цели.
				c.F = c.GetF()       // длина пути до цели
				c.Parent = curr      //ref is needed?

				(*open)[c.Key()] = c // добавляем точку в масив не посещеных
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
