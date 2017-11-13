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
	WatchUnit map[int]map[int]*Unit // map[X]map[Y]
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
	Queue int
}

func (unit *Unit) AddWatchUnit(watchUnit *Unit) {
	if unit.WatchUnit != nil {
		if unit.WatchUnit[watchUnit.X] != nil {
			unit.WatchUnit[watchUnit.X][watchUnit.Y] = watchUnit
		} else {
			unit.WatchUnit[watchUnit.X] = make(map[int]*Unit)
			unit.AddWatchUnit(watchUnit)
		}
	} else {
		unit.WatchUnit = make(map[int]map[int]*Unit)
		unit.AddWatchUnit(watchUnit)
	}
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
