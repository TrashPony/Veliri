package unit

import (
	"../../detailUnit"
	"../coordinate"
	"../effect"
)

type Unit struct {
	Id    int    `json:"id"`
	Owner string `json:"owner"`

	Chassis *detailUnit.Chassis `json:"chassis"`
	Weapon  *detailUnit.Weapon  `json:"weapon"`
	Tower   *detailUnit.Tower   `json:"tower"`
	Body    *detailUnit.Body    `json:"body"`
	Radar   *detailUnit.Radar   `json:"radar"`

	Weight int `json:"weight"`

	// Движение
	MoveSpeed  int `json:"move_speed"`
	Initiative int `json:"initiative"`

	// Атака
	Damage         int    `json:"damage"`
	RangeAttack    int    `json:"range_attack"`
	MinAttackRange int    `json:"min_attack_range"`
	AreaAttack     int    `json:"area_attack"`
	TypeAttack     string `json:"type_attack"`

	// Выживаемость
	HP              int `json:"hp"`
	Armor           int `json:"armor"`
	EvasionCritical int `json:"evasion_critical"`
	VulKinetics     int `json:"vul_kinetics"`
	VulThermal      int `json:"vul_thermal"`
	VulEM           int `json:"vul_em"`
	VulExplosive    int `json:"vul_explosive"`

	// Навигация
	RangeView int  `json:"range_view"`
	Accuracy  int  `json:"accuracy"`
	WallHack  bool `json:"wall_hack"`

	// Позиция
	X      int  `json:"x"`
	Y      int  `json:"y"`
	Rotate int  `json:"rotate"`
	OnMap  bool `json:"on_map"`

	// Игровая статистика
	Action bool                   `json:"action"`
	Target *coordinate.Coordinate `json:"target"`
	Queue  int                    `json:"queue"`

	// Бафы, дебафы
	Effects []effect.Effect `json:"effect"`
}

func (unit *Unit) SetX(x int) {
	unit.X = x
}

func (unit *Unit) GetX() int {
	return unit.X
}

func (unit *Unit) SetY(y int) {
	unit.Y = y
}

func (unit *Unit) GetY() int {
	return unit.Y
}

func (unit *Unit) GetWatchZone() int {
	return unit.RangeView
}

func (unit *Unit) GetOwnerUser() string {
	return unit.Owner
}

func (unit *Unit) GetOnMap() bool {
	return unit.OnMap
}

func (unit *Unit) SetOnMap(bool bool) {
	unit.OnMap = bool
}

func (unit *Unit) SetChassis(chassis *detailUnit.Chassis) {
	unit.Chassis = chassis
}

func (unit *Unit) SetWeapon(weapon *detailUnit.Weapon) {
	unit.Weapon = weapon
}

func (unit *Unit) SetTower(tower *detailUnit.Tower) {
	unit.Tower = tower
}

func (unit *Unit) SetBody(body *detailUnit.Body) {
	unit.Body = body
}

func (unit *Unit) SetRadar(radar *detailUnit.Radar) {
	unit.Radar = radar
}
