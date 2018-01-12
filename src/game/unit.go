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
	NameType    string      `json:"name_type"`
	NameUser    string      `json:"name_user"`
	Hp          int         `json:"hp"`
	Action      bool        `json:"action"`
	Target      *Coordinate `json:"target"`
	X           int         `json:"x"`
	Y           int         `json:"y"`
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

func (unit *Unit) getNameUser() string {
	return unit.NameUser
}