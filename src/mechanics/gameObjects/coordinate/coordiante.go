package coordinate

import (
	"strconv"
	"../effect"
)

type Coordinate struct {
	ID            int              `json:"id"`
	Type          string           `json:"type"`
	TextureFlore  string           `json:"texture_flore"`
	TextureObject string           `json:"texture_object"`
	GameID        int              `json:"game_id"`
	X             int              `json:"x"`
	Y             int              `json:"y"`
	Z             int              `json:"z"`
	R             int              `json:"r"`
	Q             int              `json:"q"`
	State         int              `json:"state"`
	Effects       []*effect.Effect `json:"effects"`
	Move          bool             `json:"move"`
	View          bool             `json:"view"`
	Attack        bool             `json:"attack"`
	PassableEdges bool             `json:"passable_edges"`
	Level         int              `json:"level"`
	H, G, F       int
	Parent        *Coordinate
}

func (coor *Coordinate) GetZ() int {
	return coor.Z
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
	if target.Q != coor.Q && // настолько я понял если конец пути находиться на искосок то стоимость клетки 14
		target.R != coor.R { // можно реализовывать стоимость пути по различной поверхности
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
	return strconv.Itoa(coor.Q) + ":" + strconv.Itoa(coor.R)
}

func (coor *Coordinate) Equal(b *Coordinate) bool { // сравнивает точки на одинаковость
	return coor.Q == b.Q && coor.R == b.R
}
