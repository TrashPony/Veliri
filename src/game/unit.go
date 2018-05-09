package game

import "../DetailUnit"

type Unit struct {
	Id    int    `json:"id"`
	Owner string `json:"owner"`

	Chassis *DetailUnit.Chassis `json:"chassis"`
	Weapon  *DetailUnit.Weapon  `json:"weapon"`
	Tower   *DetailUnit.Tower   `json:"tower"`
	Body    *DetailUnit.Body    `json:"body"`
	Radar   *DetailUnit.Radar   `json:"radar"`

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
	Action bool        `json:"action"`
	Target *Coordinate `json:"target"`
	Queue  int         `json:"queue"`
}

func (unit *Unit) getX() int {
	return unit.X
}

func (unit *Unit) getY() int {
	return unit.Y
}

func (unit *Unit) getWatchZone() int {
	return unit.RangeView
}

func (unit *Unit) getOwnerUser() string {
	return unit.Owner
}

func (unit *Unit) SetChassis(chassis *DetailUnit.Chassis) {
	unit.Chassis = chassis
}

func (unit *Unit) SetWeapon(weapon *DetailUnit.Weapon) {
	unit.Weapon = weapon
}

func (unit *Unit) SetTower(tower *DetailUnit.Tower) {
	unit.Tower = tower
}

func (unit *Unit) SetBody(body *DetailUnit.Body) {
	unit.Body = body
}

func (unit *Unit) SetRadar(radar *DetailUnit.Radar) {
	unit.Radar = radar
}
