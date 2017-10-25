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
	AreaAttack int
	TypeAttack string
	Price int
	NameType string
	NameUser string
	Hp int
	Action bool
	Target int
	X int
	Y int
}

type Coordinate struct {
	X int
	Y int
}
