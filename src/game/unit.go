package game

type Unit struct {
	Id          int         `json:"id"`
	IdGame      int         `json:"id_game"`
	Damage      int         `json:"damage"`
	MoveSpeed   int         `json:"move_speed"`
	Initiative  int         `json:"initiative"`
	RangeAttack int         `json:"range_attack"`
	WatchZone   int         `json:"watch_zone"`
	AreaAttack  int         `json:"area_attack"`
	TypeAttack  string      `json:"type_attack"`
	Price       int         `json:"price"`
	ChassisType string      `json:"chassis_type"`
	WeaponType  string      `json:"weapon_type"`
	Owner       string      `json:"owner"`
	Hp          int         `json:"hp"`
	Action      bool        `json:"action"`
	Target      *Coordinate `json:"target"`
	X           int         `json:"x"`
	Y           int         `json:"y"`
	Rotate      int         `json:"rotate"`
	Queue       int         `json:"queue"`
}

func (unit *Unit) getX() int {
	return unit.X
}

func (unit *Unit) getY() int {
	return unit.Y
}

func (unit *Unit) getWatchZone() int {
	return unit.WatchZone
}

func (unit *Unit) getOwnerUser() string {
	return unit.Owner
}
