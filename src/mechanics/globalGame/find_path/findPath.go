package find_path

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
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

func FindPath(client *player.Player, gameMap *_map.Map, start, end *coordinate.Coordinate, gameUnit *unit.Unit, scaleMap int) (error, []*coordinate.Coordinate) {

	xSize, ySize := gameMap.SetXYSize(globalGame.HexagonWidth, globalGame.HexagonHeight, scaleMap) // расчтиамем высоту и ширину карты в ху

	start.X, start.Y = start.X/scaleMap, start.Y/scaleMap
	end.X, end.Y = end.X/scaleMap, end.Y/scaleMap

	if end.X >= xSize || end.Y >= ySize || end.X < 0 || end.Y < 0 || start.X >= xSize || start.Y >= ySize || start.X < 0 || start.Y < 0 {
		return errors.New("no path"), nil
	}

	//matrix := make([][]coordinate.Coordinate, xSize, xSize*ySize) //создаем матрицу для всех точек на карте
	//for i := 0; i < len(matrix); i++ {
	//	matrix[i] = make([]coordinate.Coordinate, ySize)
	//}
	//
	//for x := 0; x < xSize; x++ { //заполняем матрицу координатами
	//	for y := 0; y < ySize; y++ {
	//		matrix[x][y] = coordinate.Coordinate{X: x, Y: y, State: FREE}
	//	}
	//}

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
		parseNeighbours(client, current, &openPoints, &closePoints, gameMap, end, gameUnit, xSize, ySize, scaleMap)
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

func parseNeighbours(client *player.Player, curr *coordinate.Coordinate, open, close *Points, gameMap *_map.Map,
	end *coordinate.Coordinate, gameUnit *unit.Unit, xSize, ySize, scaleMap int) {

	delete(*open, curr.Key())   // удаляем ячейку из не посещенных
	(*close)[curr.Key()] = curr // добавляем в массив посещенные

	nCoordinate := generateNeighboursCoordinate(client, curr, gameMap, gameUnit, scaleMap) // берем всех соседей этой клетки

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
	min = &coordinate.Coordinate{F: xSize*ySize*10 + 1}

	for _, p := range points {
		if p.F < min.F {
			min = p
		}
	}
	return
}
