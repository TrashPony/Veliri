package game

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

type UserStat struct {
	IdGame int
	Name   string
	Price  int
	Ready  string
	RespX int
	RespY int
}
