package find_path

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
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

// TODO переделать POINT в координаты, обьеденить методы с методами из файла "moveUnit"
// todo и использовать координаты из существующей карты
type Points map[string]coordinate.Coordinate

func FindPath(client *player.Player, gameMap *_map.Map, start *coordinate.Coordinate, end *coordinate.Coordinate, gameUnit *unit.Unit) (error, []*coordinate.Coordinate) {

	// TODO +100 это какойто ебаный кастыль
	matrix := make([][]coordinate.Coordinate, gameMap.QSize+100, gameMap.QSize*gameMap.RSize+100) //создаем матрицу для всех точек на карте
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]coordinate.Coordinate, gameMap.RSize+100)
	}

	for q := 0; q < gameMap.QSize; q++ { //заполняем матрицу координатами
		for r := 0; r < gameMap.RSize; r++ {
			matrix[q][r] = coordinate.Coordinate{Q: q, R: r, State: FREE}
		}
	}

	openPoints, closePoints := Points{}, Points{} // создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints[start.Key()] = *start              // кладем в карту посещенных точек стартовую точку

	matrix[start.Q][start.R] = *start // кладем первую координату в путь
	matrix[end.Q][end.R] = *end       // кладем последнюю координату в уже провереные

	var path []*coordinate.Coordinate
	var noSortedPath []*coordinate.Coordinate
	for {
		if len(openPoints) <= 0 {
			return errors.New("no path"), nil
		}
		current := *MinF(openPoints, gameMap) // Берем точку с мин стоимостью пути
		if current.Equal(end) {               // если текущая точка и есть конец начинаем генерить путь
			for !current.Equal(start) { // если текущая точка не стартовая точка то цикл крутиться путь мутиться
				current = *current.Parent  // берем текущую точку и на ее место ставить ее родителя
				if !current.Equal(start) { // если текущая точка попрежнему не стартовая то
					matrix[current.Q][current.R].State = PATH // помечаем ее как часть пути

					gameCoordinate, find := gameMap.GetCoordinate(current.Q, current.R)
					if find {
						noSortedPath = append(noSortedPath, gameCoordinate)
					}
				}
			}
			break
		}
		parseNeighbours(client, current, &matrix, &openPoints, &closePoints, gameMap, end, gameUnit)
	}

	for i := len(noSortedPath); i > 0; i-- {
		path = append(path, noSortedPath[i-1])
	}

	path = append(path, end)
	return nil, path
}

func parseNeighbours(client *player.Player, curr coordinate.Coordinate, m *[][]coordinate.Coordinate, open, close *Points, gameMap *_map.Map, end *coordinate.Coordinate, gameUnit *unit.Unit) {
	delete(*open, curr.Key())   // удаляем ячейку из не посещенных
	(*close)[curr.Key()] = curr // добавляем в массив посещенные

	nCoordinate := generateNeighboursCoordinate(client, &curr, gameMap, gameUnit) // берем всех соседей этой клетки

	for _, qLine := range nCoordinate {
		for _, c := range qLine {
			if c.Q < gameMap.QSize && c.R < gameMap.RSize {
				tmpPoint := (*m)[c.Q][c.R] // берем поинт из матрицы

				if _, inClose := (*close)[tmpPoint.Key()]; inClose || tmpPoint.State == BLOCKED {
					continue // если ячейка является блокированой или находиться в масиве посещенных то пропускаем ее
				}

				if _, inOpen := (*open)[tmpPoint.Key()]; inOpen {
					continue // если ячейка уже добавленна в массив еще не посещенных то пропускаем
				}

				// считаем для поинта значения пути
				tmpPoint.G = curr.GetG(tmpPoint)  // стоимость клетки
				tmpPoint.H = GetH(tmpPoint, *end) // приближение от точки до конечной цели.
				tmpPoint.F = tmpPoint.GetF()      // длина пути до цели
				tmpPoint.Parent = &curr           //ref is needed?

				(*open)[tmpPoint.Key()] = tmpPoint // добавляем точку в масив не посещеных
			}
		}
	}
}

func GetH(a, b coordinate.Coordinate) int { // эвристическое приближение стоимости пути от v до конечной цели.
	tmp := math.Abs(float64(a.Q - b.Q)) // вычисляем разницу между точкой и концом пути по Х
	tmp += math.Abs(float64(a.R - b.R)) // вычисляем разницу между точкой и концом пути по Y и сумируем с раницой по X

	return int(tmp)
}

func MinF(points Points, gameMap *_map.Map) (min *coordinate.Coordinate) { // берет точку с минимальной стоимостью пути из масива не посещеных
	min = &coordinate.Coordinate{F: gameMap.QSize*gameMap.RSize*10 + 1}

	for _, p := range points {
		if p.F < min.F {
			*min = p
		}
	}
	return
}
