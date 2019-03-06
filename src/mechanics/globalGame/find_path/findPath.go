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

type Points map[string]coordinate.Coordinate

func FindPath(client *player.Player, gameMap *_map.Map, start *coordinate.Coordinate, end *coordinate.Coordinate, gameUnit *unit.Unit, scaleMap int) (error, []*coordinate.Coordinate) {

	xSize, ySize := gameMap.SetXYSize(globalGame.HexagonWidth, globalGame.HexagonHeight, scaleMap) // расчтиамем высоту и ширину карты в ху

	start.X, start.Y = start.X/scaleMap, start.Y/scaleMap
	end.X, end.Y = end.X/scaleMap, end.Y/scaleMap

	if end.X > xSize || end.Y > ySize {
		return errors.New("no path"), nil
	}

	matrix := make([][]coordinate.Coordinate, xSize, xSize*ySize) //создаем матрицу для всех точек на карте
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]coordinate.Coordinate, ySize)
	}

	for x := 0; x < xSize; x++ { //заполняем матрицу координатами
		for y := 0; y < ySize; y++ {
			matrix[x][y] = coordinate.Coordinate{X: x, Y: y, State: FREE}
		}
	}

	openPoints, closePoints := Points{}, Points{} // создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints[start.Key()] = *start              // кладем в карту посещенных точек стартовую точку

	matrix[start.X][start.Y] = *start // кладем первую координату в путь
	matrix[end.X][end.Y] = *end       // кладем последнюю координату в уже провереные

	var path []*coordinate.Coordinate
	var noSortedPath []*coordinate.Coordinate

	for {
		if len(openPoints) <= 0 {
			return errors.New("no path"), nil
		}
		current := *MinF(openPoints, xSize, ySize) // Берем точку с мин стоимостью пути
		if current.EqualXY(end) {                  // если текущая точка и есть конец начинаем генерить путь
			for !current.EqualXY(start) { // если текущая точка не стартовая точка то цикл крутиться путь мутиться
				current = *current.Parent    // берем текущую точку и на ее место ставить ее родителя
				if !current.EqualXY(start) { // если текущая точка попрежнему не стартовая то
					matrix[current.X][current.Y].State = PATH // помечаем ее как часть пути
					noSortedPath = append(noSortedPath, &matrix[current.X][current.Y])
				}
			}
			break
		}
		parseNeighbours(client, current, &matrix, &openPoints, &closePoints, gameMap, end, gameUnit, xSize, ySize, scaleMap)
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

func parseNeighbours(client *player.Player, curr coordinate.Coordinate, m *[][]coordinate.Coordinate, open,
	close *Points, gameMap *_map.Map, end *coordinate.Coordinate, gameUnit *unit.Unit, xSize, ySize, scaleMap int) {
	delete(*open, curr.Key())   // удаляем ячейку из не посещенных
	(*close)[curr.Key()] = curr // добавляем в массив посещенные

	nCoordinate := generateNeighboursCoordinate(client, &curr, gameMap, gameUnit, scaleMap) // берем всех соседей этой клетки

	for _, xLine := range nCoordinate {
		for _, c := range xLine {
			if c.X < xSize && c.Y < ySize && c.X > 0 && c.Y > 0 {
				tmpPoint := (*m)[c.X][c.Y] // берем поинт из матрицы

				if _, inClose := (*close)[tmpPoint.Key()]; inClose || tmpPoint.State == BLOCKED {
					continue // если ячейка является блокированой или находиться в масиве посещенных то пропускаем ее
				}

				if _, inOpen := (*open)[tmpPoint.Key()]; inOpen {
					continue // если ячейка уже добавленна в массив еще не посещенных то пропускаем
				}

				// считаем для поинта значения пути
				tmpPoint.G = curr.GetXYG(tmpPoint) // стоимость клетки
				tmpPoint.H = GetH(tmpPoint, *end)  // приближение от точки до конечной цели.
				tmpPoint.F = tmpPoint.GetF()       // длина пути до цели
				tmpPoint.Parent = &curr            //ref is needed?

				(*open)[tmpPoint.Key()] = tmpPoint // добавляем точку в масив не посещеных
			}
		}
	}
}

func GetH(a, b coordinate.Coordinate) int { // эвристическое приближение стоимости пути от v до конечной цели.
	tmp := math.Abs(float64(a.X - b.X)) // вычисляем разницу между точкой и концом пути по Х
	tmp += math.Abs(float64(a.Y - b.Y)) // вычисляем разницу между точкой и концом пути по Y и сумируем с раницой по X

	return int(tmp)
}

func MinF(points Points, xSize, ySize int) (min *coordinate.Coordinate) { // берет точку с минимальной стоимостью пути из масива не посещеных
	min = &coordinate.Coordinate{F: xSize*ySize*10 + 1}

	for _, p := range points {
		if p.F < min.F {
			*min = p
		}
	}
	return
}
