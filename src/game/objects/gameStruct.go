package objects

type UnitType struct {
	Id int
	Type string
	Damage int
	Hp int
	MoveSpeed int
	Init int
	RangeAttack int
	WatchZone int
	AreaAttack int
	TypeAttack string
	Price int
}

type Unit struct {
	Id int
	IdGame int
	Damage int
	MoveSpeed int
	Init int
	RangeAttack int
	WatchZone int
	Watch map[string]*Coordinate //KEY format X:Y
	WatchUnit map[string]*Unit //KEY format X:Y
	AreaAttack int
	TypeAttack string
	Price int
	NameType string
	NameUser string
	Hp int
	Action bool
	Target string
	X int
	Y int
}

type UserStat struct {
	IdGame int
	Name   string
	IdResp int
	Price  int
	Ready string
}

type Coordinate struct {
	X int
	Y int
}
