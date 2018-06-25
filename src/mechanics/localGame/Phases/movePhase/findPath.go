package movePhase

//** SOURCE CODE https://github.com/JavaDar/aStar **//
import (
	"math"
	"../../../gameObjects/map"
	"../../map/coordinate"
	"../../../player"
)

const (
	FREE    = iota
	BLOCKED
	START
	END
	PATH
)

var (
	START_POINT, END_POINT coordinate.Coordinate
	WIDTH, HEIGHT          int
	matrix                 [][]coordinate.Coordinate
)

// TODO переделать POINT в координаты, обьеденить методы с методами из файла "moveUnit"
type Points map[string]coordinate.Coordinate

func FindPath(client *player.Player, gameMap *_map.Map, start *coordinate.Coordinate, end *coordinate.Coordinate) []*coordinate.Coordinate {

	START_POINT = coordinate.Coordinate{X: start.X, Y: start.Y, State: START} // начальная точка
	END_POINT = coordinate.Coordinate{X: end.X, Y: end.Y, State: END}         // конечная точка
	WIDTH = gameMap.XSize                                                     // ширина карты
	HEIGHT = gameMap.YSize                                                    // высота карты

	matrix = make([][]coordinate.Coordinate, WIDTH, WIDTH*HEIGHT) //создаем матрицу для всех точек на карте
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]coordinate.Coordinate, HEIGHT)
	}

	for x := 0; x < WIDTH; x++ { //заполняем матрицу координатами
		for y := 0; y < HEIGHT; y++ {
			matrix[x][y] = coordinate.Coordinate{X: x, Y: y, State: FREE}
		}
	}

	openPoints, closePoints := Points{}, Points{} // создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints[START_POINT.Key()] = START_POINT   // кладем в карту посещенных точек стартовую точку

	matrix[START_POINT.X][START_POINT.Y] = START_POINT // магия 	//set start & finish
	matrix[END_POINT.X][END_POINT.Y] = END_POINT       // магия

	var path []*coordinate.Coordinate
	var noSortedPath []*coordinate.Coordinate
	for {
		current := *MinF(openPoints) // Берем точку с мин стоимостью пути
		if current.Equal(&END_POINT) { // если текущая точка и есть конец начинаем генерить путь
			for !current.Equal(&START_POINT) { // если текущая точка не стартовая точка то цикл крутиться путь мутиться
				current = *current.Parent // берем текущую точку и на ее место ставить ее родителя
				if !current.Equal(&START_POINT) { // если текущая точка попрежнему не стартовая то
					matrix[current.X][current.Y].State = PATH // помечаем ее как часть пути

					gameCoordinate, find := gameMap.GetCoordinate(current.X, current.Y)
					if find {
						noSortedPath = append(noSortedPath, gameCoordinate)
					}
				}
			}
			break
		}
		parseNeighbours(client, current, &matrix, &openPoints, &closePoints, gameMap)
	}

	for i := len(noSortedPath); i > 0; i-- {
		path = append(path, noSortedPath[i-1])
	}

	path = append(path, end)
	return path
}

func parseNeighbours(client *player.Player, curr coordinate.Coordinate, m *[][]coordinate.Coordinate, open, close *Points, gameMap *_map.Map) {
	delete(*open, curr.Key())   // удаляем ячейку из не посещенных
	(*close)[curr.Key()] = curr // добавляем в массив посещенные

	nCoordinate := generateNeighboursCoordinate(client, &curr, gameMap) // берем всех соседей этой клетки

	for _, xLine := range nCoordinate {
		for _, c := range xLine {
			if c.X < WIDTH && c.Y < HEIGHT {
				tmpPoint := (*m)[c.X][c.Y] // берем поинт из матрицы

				if _, inClose := (*close)[tmpPoint.Key()]; inClose || tmpPoint.State == BLOCKED {
					continue // если ячейка является блокированой или находиться в масиве посещенных то пропускаем ее
				}

				if _, inOpen := (*open)[tmpPoint.Key()]; inOpen {
					continue // если ячейка уже добавленна в массив еще не посещенных то пропускаем
				}

				// считаем для поинта значения пути
				tmpPoint.G = curr.GetG(tmpPoint)       // стоимость клетки
				tmpPoint.H = GetH(tmpPoint, END_POINT) // приближение от точки до конечной цели.
				tmpPoint.F = tmpPoint.GetF()           // длина пути до цели
				tmpPoint.Parent = &curr                //ref is needed?

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

func MinF(points Points) (min *coordinate.Coordinate) { // берет точку с минимальной стоимостью пути из масива не посещеных
	min = &coordinate.Coordinate{F: WIDTH*HEIGHT*10 + 1}

	for _, p := range points {
		if p.F < min.F {
			*min = p
		}
	}
	return
}
