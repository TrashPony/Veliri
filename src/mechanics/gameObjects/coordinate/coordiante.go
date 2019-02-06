package coordinate

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamicMapObject"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/effect"
	"strconv"
)

type Coordinate struct {
	ID                  int                             `json:"id"`
	Type                string                          `json:"type"`
	TextureFlore        string                          `json:"texture_flore"`
	TextureOverFlore    string                          `json:"texture_over_flore"`
	TextureObject       string                          `json:"texture_object"`
	AnimateSpriteSheets string                          `json:"animate_sprite_sheets"`
	AnimateLoop         bool                            `json:"animate_loop"`
	ImpactRadius        int                             `json:"impact_radius"`
	Impact              *Coordinate                     `json:"impact"`
	GameID              int                             `json:"game_id"`
	X                   int                             `json:"x"`
	Y                   int                             `json:"y"`
	Z                   int                             `json:"z"`
	R                   int                             `json:"r"`
	Q                   int                             `json:"q"`
	State               int                             `json:"state"`
	Effects             []*effect.Effect                `json:"effects"`
	Move                bool                            `json:"move"`
	View                bool                            `json:"view"`
	Attack              bool                            `json:"attack"`
	Level               int                             `json:"level"`
	Scale               int                             `json:"scale"`
	Shadow              bool                            `json:"shadow"`
	UnitOverlap         bool                            `json:"unit_overlap"`
	ObjRotate           int                             `json:"obj_rotate"`
	AnimationSpeed      int                             `json:"animation_speed"`
	XOffset             int                             `json:"x_offset"`
	YOffset             int                             `json:"y_offset"`
	DynamicObject       *dynamicMapObject.DynamicObject `json:"dynamic_object"`
	H, G, F             int
	Parent              *Coordinate

	/* если тру то с течением времени или по эвенту игрока эвакуируют с этой клетки без его желания */
	Transport bool `json:"transport"`
	/* если строка не пуста значит эта клетка прослушивается, например вход в базу (base) или переход в другой сектор (sector),
	   и когда игрок на ней происходит событие */
	Handler string `json:"handler"`

	/* говорит работает хендлер или нет, например занята ячейка перехода и тп не работает*/
	HandlerOpen bool `json:"handler_open"`

	/* соотвественно место куда попадает игрок после ивента */
	ToQ      int `json:"to_q"`
	ToR      int `json:"to_r"`
	ToBaseID int `json:"to_base_id"`
	ToMapID  int `json:"to_map_id"`
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

func (coor *Coordinate) CalculateXYZ() {
	coor.X = coor.Q - (coor.R-(coor.R&1))/2
	coor.Z = coor.R
	coor.Y = -coor.X - coor.Z
}
