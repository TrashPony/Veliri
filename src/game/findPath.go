package game

//** SOURCE CODE https://github.com/JavaDar/aStar **//
import (
	"math"
	"strconv"
)

const (
	FREE = iota
	BLOCKED
	START
	END
	PATH
)

var (
	START_POINT, END_POINT Coordinate
	WIDTH, HEIGHT          int
	matrix                 [][]Coordinate
)


// TODO переделать POINT в координаты, обьеденить методы с методами из файла "moveUnit"
type Points map[string]Coordinate

func FindPath(gameMap *Map, start Coordinate, end Coordinate, obstacles map[int]map[int]*Coordinate) []Coordinate {

	START_POINT = Coordinate{X: start.X, Y: start.Y, State: START} // начальная точка
	END_POINT = Coordinate{X: end.X, Y: end.Y, State: END}         // конечная точка
	WIDTH = gameMap.Xsize                                     // ширина карты
	HEIGHT = gameMap.Ysize                                    // высота карты

	matrix = make([][]Coordinate, WIDTH, WIDTH*HEIGHT) //создаем матрицу для всех точек на карте
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]Coordinate, HEIGHT)
	}

	for x := 0; x < WIDTH; x++ { //заполняем матрицу координатами
		for y := 0; y < HEIGHT; y++ {
			matrix[x][y] = Coordinate{X: x, Y: y, State: FREE}
		}
	}

	openPoints, closePoints := Points{}, Points{} // создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints[START_POINT.Key()] = START_POINT   // кладем в карту посещенных точек стартовую точку

	matrix[START_POINT.X][START_POINT.Y] = START_POINT // магия 	//set start & finish
	matrix[END_POINT.X][END_POINT.Y] = END_POINT       // магия

	var path []Coordinate
	var noSortedPath []Coordinate
	for {
		current := *MinF(openPoints)  // Берем точку с мин стоимостью пути
		if current.Equal(END_POINT) { // если текущая точка и есть конец начинаем генерить путь
			for !current.Equal(START_POINT) { // если текущая точка не стартовая точка то цикл крутиться путь мутиться
				current = *current.Parent        // берем текущую точку и на ее место ставить ее родителя
				if !current.Equal(START_POINT) { // если текущая точка попрежнему не стартовая то
					matrix[current.X][current.Y].State = PATH // помечаем ее как часть пути
					noSortedPath = append(noSortedPath, Coordinate{X: matrix[current.X][current.Y].X, Y: matrix[current.X][current.Y].Y})
				}
			}
			break
		}
		parseNeighbours(current, &matrix, &openPoints, &closePoints, obstacles)
	}

	for i := len(noSortedPath); i > 0; i-- {
		path = append(path, noSortedPath[i-1])
	}

	path = append(path, end)
	return path
}

func parseNeighbours(curr Coordinate, m *[][]Coordinate, open, close *Points, obstacles map[int]map[int]*Coordinate) {
	delete(*open, curr.Key())   // удаляем ячейку из не посещенных
	(*close)[curr.Key()] = curr // добавляем в массив посещенные

	nCoord := generateNeighboursCoord(&curr, obstacles) // берем всех соседей этой клетки

	for _, xLine := range nCoord {
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

func GetH(a, b Coordinate) int { // эвристическое приближение стоимости пути от v до конечной цели.
	tmp := math.Abs(float64(a.X - b.X)) // вычисляем разницу между точкой и концом пути по Х
	tmp += math.Abs(float64(a.Y - b.Y)) // вычисляем разницу между точкой и концом пути по Y и сумируем с раницой по X

	return int(tmp)
}

func (current Coordinate) GetG(target Coordinate) int { // наименьшая стоимость пути в End из стартовой вершины
	if target.X != current.X && // настолько я понял если конец пути находиться на искосок то стоимость клетки 14
		target.Y != current.Y { // можно реализовывать стоимость пути по различной поверхности
		return current.G + 14
	}

	return current.G + 10 // если находиться на 1 линии по Х или У то стоимость 10
}

/* Фактически, функция f(v) — длина пути до цели, которая складывается из пройденного расстояния g(v) и оставшегося расстояния h(v). Исходя из этого, чем меньше значение f(v),
тем раньше мы откроем вершину v, так как через неё мы предположительно достигнем расстояние до цели быстрее всего. Открытые алгоритмом вершины можно хранить в очереди с приоритетом
по значению f(v). А* действует подобно алгоритму Дейкстры и просматривает среди всех маршрутов ведущих к цели сначала те, которые благодаря имеющейся информации
(эвристическая функция) в данный момент являются наилучшими. */

func (p Coordinate) GetF() int { // длина пути до цели, которая складывается из пройденного расстояния g(v) и оставшегося расстояния h(v).
	return p.G + p.H // складываем пройденое расстония и оставшееся
}

func (p Coordinate) Key() string { //создает уникальный ключ для карты "X:Y"
	return strconv.Itoa(p.X) + ":" + strconv.Itoa(p.Y)
}

func (a Coordinate) Equal(b Coordinate) bool { // сравнивает точки на одинаковость
	return a.X == b.X && a.Y == b.Y
}

func MinF(points Points) (min *Coordinate) { // берет точку с минимальной стоимостью пути из масива не посещеных
	min = &Coordinate{F: WIDTH*HEIGHT*10 + 1}

	for _, p := range points {
		if p.F < min.F {
			*min = p
		}
	}
	return
}

func addCoordIfValid(res map[int]map[int]*Coordinate, obstacles map[int]map[int]*Coordinate, x int, y int) {
	coor := Coordinate{X:x , Y:y}

	_, ok := obstacles[x][y]
	if !ok && (x >= 0 && y >= 0){
		if res[x] != nil {
			res[x][y] = &coor
		} else {
			res[x] = make(map[int]*Coordinate)
			res[x][y] = &coor
		}
	}
}

func generateNeighboursCoord(curr *Coordinate, obstacles map[int]map[int]*Coordinate) (res map[int]map[int]*Coordinate) { // берет все соседние клетки от текущей
	res = make(map[int]map[int]*Coordinate)

	//строго лево
	_, left := obstacles[curr.X-1][curr.Y]
	addCoordIfValid(res, obstacles, curr.X-1, curr.Y)
	//строго право
	_, right := obstacles[curr.X+1][curr.Y]
	addCoordIfValid(res, obstacles, curr.X+1, curr.Y)
	//верх центр
	_, top := obstacles[curr.X][curr.Y-1]
	addCoordIfValid(res, obstacles, curr.X, curr.Y-1)
	//низ центр
	_, bottom := obstacles[curr.X][curr.Y+1]
	addCoordIfValid(res, obstacles, curr.X, curr.Y+1)


	//верх лево/    ЛЕВО И верх
	if !(left || top) {
		addCoordIfValid(res, obstacles, curr.X-1, curr.Y-1)
	}
	//верх право/   ПРАВО И верх
	if !(right || top) {
		addCoordIfValid(res, obstacles, curr.X+1, curr.Y-1)
	}
	//низ лево/  если ЛЕВО И низ
	if !(left || bottom) {
		addCoordIfValid(res, obstacles, curr.X-1, curr.Y+1)
	}
	//низ право/  низ И ВЕРХ
	if !(right || bottom) {
		addCoordIfValid(res, obstacles, curr.X+1, curr.Y+1)
	}

	return
}
