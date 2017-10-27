package mechanics
//** SOURCE CODE https://github.com/JavaDar/aStar **//
import (
	"log"
	"math"
	"strconv"
	"../objects"
)

const (
	FREE = iota
	BLOCKED
	START
	END
	PATH
)

var (
	START_POINT, END_POINT Point
	WIDTH, HEIGHT int
	matrix [][]Point
)

type Point struct {
	x, y, state int
	H, G, F int
	parent *Point
}

type Points map[string]Point

func FindPath(gameMap objects.Map, start objects.Coordinate, end objects.Coordinate, obstacles []objects.Coordinate)  {

	START_POINT = Point{x: start.X, y: start.Y, state:START} // начальная точка
	END_POINT = Point{x: end.X, y: end.Y, state:END} 		  // конечная точка
	WIDTH = gameMap.Xsize									  // ширина карты
	HEIGHT = gameMap.Ysize									  // высота карты

	matrix = make([][]Point,WIDTH, WIDTH*HEIGHT)             //создаем матрицу для всех точек на карте
	for i:=0; i<len(matrix);i ++ {
		matrix[i] = make([]Point, HEIGHT)
	}

	for x := 0; x < WIDTH; x++ {							  //заполняем матрицу координатами
		for y := 0; y < HEIGHT; y++ {
			matrix[x][y] = Point{x:x, y:y, state:FREE}
		}
	}

	openPoints, closePoints := Points{}, Points{}			  // создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints[START_POINT.Key()] = START_POINT				  // кладем в карту посещенных точек стартовую точку

	matrix[START_POINT.x][START_POINT.y] = START_POINT        // магия 	//set start & finish
	matrix[END_POINT.x][END_POINT.y] = END_POINT			  // магия

	for _, o := range obstacles {							  // ставим препятсвиям статус в точке Блокировано
		matrix[o.X][o.Y].state = BLOCKED
	}

	for {
		current := *MinF(openPoints)                          // Берем точку с мин стоимостью пути
		if current.Equal(END_POINT) {                         // если текущая точка и есть конец начинаем генерить путь
			log.Println("Path is found")

			log.Print("Points in path: ")
			for !current.Equal(START_POINT) {				  // если текущая точка не стартовая точка то цикл крутиться путь мутиться
				current = *current.parent					  // берем текущую точку и на ее место ставить ее родителя
				if !current.Equal(START_POINT){				  // если текущая точка попрежнему не стартовая то
					matrix[current.x][current.y].state = PATH // помечаем ее как часть пути
					log.Print(matrix[current.x][current.y], " ")
				}
			}
			break
		}
		parseNeighbours(current, &matrix, &openPoints, &closePoints)
	}

	log.Println("Exited!")
	//end
}

func parseNeighbours(curr Point, m *[][]Point, open, close *Points) {
	delete(*open, curr.Key()) 									// удаляем ячейку из не посещенных
	(*close)[curr.Key()] = curr 								// добавляем в массив посещенные

	nCoord := generateNeighboursCoord(curr) 					// берем всех соседей этой клетки

	for _, c := range nCoord{
		tmpPoint := (*m)[c.x][c.y]								// берем поинт из матрицы

		if _, inClose := (*close)[tmpPoint.Key()]; inClose || tmpPoint.state == BLOCKED {
			continue 											// если ячейка является блокированой или находиться в масиве посещенных то пропускаем ее
		}

		if _, inOpen := (*open)[tmpPoint.Key()]; inOpen{
			continue 											// если ячейка уже добавленна в массив еще не посещенных то пропускаем
		}

																// считаем для поинта значения пути
		tmpPoint.G = curr.GetG(tmpPoint)  						// стоимость клетки
		tmpPoint.H = GetH(tmpPoint, END_POINT) 					// приближение от точки до конечной цели.
		tmpPoint.F = tmpPoint.GetF() 							// длина пути до цели
		tmpPoint.parent = &curr 								//ref is needed?

		(*open)[tmpPoint.Key()] = tmpPoint 						// добавляем точку в масив не посещеных
	}
}

func GetH(a, b Point) int {                   // эвристическое приближение стоимости пути от v до конечной цели.
	tmp := math.Abs(float64(a.x - b.x))	      // вычисляем разницу между точкой и концом пути по Х
	tmp += math.Abs(float64(a.y - b.y))		  // вычисляем разницу между точкой и концом пути по Y и сумируем с раницой по X

	return int(tmp)
}

func (current Point) GetG(target Point) int { // наименьшая стоимость пути в End из стартовой вершины
	if target.x != current.x &&               // настолько я понял если конец пути находиться на искосок то стоимость клетки 14
		target.y != current.y {
		return current.G + 14
	}

	return current.G + 10                     // если находиться на 1 линии по Х или У то стоимость 10
}

/* Фактически, функция f(v) — длина пути до цели, которая складывается из пройденного расстояния g(v) и оставшегося расстояния h(v). Исходя из этого, чем меньше значение f(v),
тем раньше мы откроем вершину v, так как через неё мы предположительно достигнем расстояние до цели быстрее всего. Открытые алгоритмом вершины можно хранить в очереди с приоритетом
по значению f(v). А* действует подобно алгоритму Дейкстры и просматривает среди всех маршрутов ведущих к цели сначала те, которые благодаря имеющейся информации
(эвристическая функция) в данный момент являются наилучшими. */

func (p Point) GetF() int {  				// длина пути до цели, которая складывается из пройденного расстояния g(v) и оставшегося расстояния h(v).
	return p.G + p.H		 				// складываем пройденое расстония и оставшееся
}

func (p Point) Key() string { 				//создает уникальный ключ для карты "X:Y"
	return strconv.Itoa(p.x) + ":" + strconv.Itoa(p.y)
}

func (a Point) Equal(b Point) bool { 		// сравнивает точки на одинаковость
	return a.x == b.x && a.y == b.y
}

func MinF(points Points) (min *Point){ 		// берет точку с минимальной стоимостью пути из масива не посещеных
	min = &Point{F:WIDTH*HEIGHT*10+1}

	for _, p := range points{
		if p.F < min.F {
			*min = p
		}
	}
	return
}

func addCoordIfValid(coords *[]Point, x,y  int){

	if x >= 0  && y >= 0 &&
		x < WIDTH && y < HEIGHT{
		*coords = append(*coords, Point{x:x, y:y})
	}
}

func generateNeighboursCoord(curr Point) (res []Point)  { // берет все соседние клетки от текущей

	//верх лево
	addCoordIfValid(&res, curr.x -1, curr.y +1)
	//верх центр
	addCoordIfValid(&res, curr.x, curr.y +1)
	//верх право
	addCoordIfValid(&res, curr.x +1, curr.y +1)

	//строго лево
	addCoordIfValid(&res, curr.x-1, curr.y)
	//строго право
	addCoordIfValid(&res, curr.x+1, curr.y)

	//низ лево
	addCoordIfValid(&res, curr.x -1, curr.y -1)
	//низ центр
	addCoordIfValid(&res, curr.x, curr.y -1)
	//низ право
	addCoordIfValid(&res, curr.x +1, curr.y -1)

	return
}