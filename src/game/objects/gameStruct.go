package objects

type UnitType struct {
	Id          int
	Type        string
	Damage      int
	Hp          int
	MoveSpeed   int
	Init        int
	RangeAttack int
	WatchZone   int
	AreaAttack  int
	TypeAttack  string
	Price       int
}

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

type UserStat struct {
	IdGame int
	Name   string
	IdResp int
	Price  int
	Ready  string
}

type Coordinate struct {
	X int
	Y int
	Type string
	Texture string
}
