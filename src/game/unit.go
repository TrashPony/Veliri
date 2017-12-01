package game


type Unit struct {
	Id          int
	IdGame      int
	Damage      int
	MoveSpeed   int
	Initiative  int
	RangeAttack int
	WatchZone   int
	AreaAttack  int
	TypeAttack  string
	Price       int
	NameType    string
	NameUser    string
	Hp          int
	Action      bool
	Target      *Coordinate
	X           int
	Y           int
	Queue       int
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