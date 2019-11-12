package coordinate

import (
	"strconv"
)

type Coordinate struct {
	ID     int `json:"id"`
	X      int `json:"x"`
	Y      int `json:"y"`
	Rotate int `json:"rotate"` // используется при расчете поиска пути
	State  int `json:"state"`

	MapID      int `json:"map_id"`
	RespRotate int `json:"resp_rotate"`

	H, G, F int
	Parent  *Coordinate

	/* мета слушателей */
	/* если тру то с течением времени или по эвенту игрока эвакуируют с этой клетки без его желания */
	Transport bool `json:"transport"`
	/* если строка не пуста значит эта клетка прослушивается, например вход в базу (base) или переход в другой сектор (sector),
	   и когда игрок на ней происходит событие */
	Handler string `json:"handler"`

	/* говорит работает хендлер или нет, например занята ячейка перехода и тп не работает*/
	HandlerOpen bool `json:"handler_open"`

	/* соотвественно место куда попадает игрок после ивента */
	Positions []*Coordinate `json:"positions"`
	ToBaseID  int           `json:"to_base_id"`
	ToMapID   int           `json:"to_map_id"`

	Find bool `json:"-"`
}

func (coor *Coordinate) GetY() int {
	return coor.Y
}

func (coor *Coordinate) GetX() int {
	return coor.X
}

func (coor *Coordinate) GetR() int {
	return coor.X
}

func (coor *Coordinate) GetQ() int {
	return coor.X
}
func (coor *Coordinate) GetG(target Coordinate) int { // наименьшая стоимость пути в End из стартовой вершины
	if target.X != coor.X && // настолько я понял если конец пути находиться на искосок то стоимость клетки 14
		target.Y != coor.Y { // можно реализовывать стоимость пути по различной поверхности
		return coor.G + 14
	}

	return coor.G + 10 // если находиться на 1 линии по Х или У то стоимость 10
}

func (coor *Coordinate) GetXYG(target *Coordinate) int { // наименьшая стоимость пути в End из стартовой вершины
	if target.X != coor.X && // настолько я понял если конец пути находиться на искосок то стоимость клетки 14
		target.Y != coor.Y { // можно реализовывать стоимость пути по различной поверхности
		return coor.G + 14
	}

	return coor.G + 10 // если находиться на 1 линии по Х или У то стоимость 10
}

/* Фактически, функция f(v) — длина пути до цели, которая складывается из пройденного расстояния g(v) и оставшегося расстояния h(v). Исходя из этого, чем меньше значение f(v),
тем раньше мы откроем вершину v, так как через неё мы предположительно достигнем расстояние до цели быстрее всего. Открытые алгоритмом вершины можно хранить в очереди с приоритетом
по значению f(v). А* действует подобно алгоритму Дейкстры и просматривает среди всех маршрутов ведущих к цели сначала те, которые благодаря имеющейся информации
(эвристическая функция) в данный момент являются наилучшими. */

func (coor *Coordinate) GetF() int { // длина пути до цели, которая складывается из пройденного расстояния g(v) и оставшегося расстояния h(v).
	return coor.G + coor.H // складываем пройденое расстония и оставшееся
}

func (coor *Coordinate) Key() string { //создает уникальный ключ для карты "X:Y"
	return strconv.Itoa(coor.X) + ":" + strconv.Itoa(coor.Y)
}

func (coor *Coordinate) Equal(b *Coordinate) bool { // сравнивает точки на одинаковость
	return coor.X == b.X && coor.Y == b.Y
}
