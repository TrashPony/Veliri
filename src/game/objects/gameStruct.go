package objects

import (
	"strconv"
	"errors"
)

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
	Price  int
	Ready  string
	RespX int
	RespY int
}

type Coordinate struct {
	X int
	Y int
	Type string
	Texture string
}

type Structure struct {
	Type string
	NameUser    string
	WatchZone   int
	X int
	Y int
}

type Watcher interface {
	Watch() float64
}

func (unit *Unit) Watch(login string, units map[int]map[int]*Unit, allStructures map[int]map[int]*Structure) (allCoordinate map[string]*Coordinate, unitsCoordinate map[int]map[int]*Unit, structureCoordinate map[int]map[int]*Structure, Err error) {
	allCoordinate = make(map[string]*Coordinate)
	unitsCoordinate = make(map[int]map[int]*Unit)
	structureCoordinate = make(map[int]map[int]*Structure)

	if login == unit.NameUser {
		PermCoordinates := GetCoordinates(unit.X, unit.Y, unit.WatchZone)
		for i := 0; i < len(PermCoordinates); i++ {
			unitInMap, ok := units[PermCoordinates[i].X][PermCoordinates[i].Y]
			if ok {
				if unitsCoordinate[PermCoordinates[i].X] != nil {
					unitsCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = unitInMap
				} else {
					unitsCoordinate[PermCoordinates[i].X] = make(map[int]*Unit)
					unitsCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = unitInMap
				}
			} else {
				var structureInMap *Structure
				structureInMap, ok = allStructures[PermCoordinates[i].X][PermCoordinates[i].Y]
				if ok {
					if structureCoordinate[PermCoordinates[i].X] != nil {
						structureCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = structureInMap
					} else {
						structureCoordinate[PermCoordinates[i].X] = make(map[int]*Structure)
						structureCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = structureInMap
					}
				} else {
					allCoordinate[strconv.Itoa(PermCoordinates[i].X)+":"+strconv.Itoa(PermCoordinates[i].Y)] = PermCoordinates[i]
				}
			}
		}
	} else {
		return allCoordinate, unitsCoordinate, structureCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitsCoordinate, structureCoordinate, nil
}

func (structure *Structure) Watch(login string, units map[int]map[int]*Unit, allStructures map[int]map[int]*Structure) (allCoordinate map[string]*Coordinate, unitCoordinate map[int]map[int]*Unit, structureCoordinate map[int]map[int]*Structure,  Err error) {
	allCoordinate = make(map[string]*Coordinate)
	unitCoordinate = make(map[int]map[int]*Unit)
	structureCoordinate = make(map[int]map[int]*Structure)

	if login == structure.NameUser {
		PermCoordinates := GetCoordinates(structure.X, structure.Y, structure.WatchZone)
		for i := 0; i < len(PermCoordinates); i++ {
			unitInMap, ok := units[PermCoordinates[i].X][PermCoordinates[i].Y]
			if ok {
				if unitCoordinate[PermCoordinates[i].X] != nil {
					unitCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = unitInMap
				} else {
					unitCoordinate[PermCoordinates[i].X] = make(map[int]*Unit)
					unitCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = unitInMap
				}
			} else {
				var structureInMap *Structure
				structureInMap, ok = allStructures[PermCoordinates[i].X][PermCoordinates[i].Y]
				if ok {
					if structureCoordinate[PermCoordinates[i].X] != nil {
						structureCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = structureInMap
					} else {
						structureCoordinate[PermCoordinates[i].X] = make(map[int]*Structure)
						structureCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = structureInMap
					}
				} else {
					allCoordinate[strconv.Itoa(PermCoordinates[i].X)+":"+strconv.Itoa(PermCoordinates[i].Y)] = PermCoordinates[i]
				}
			}
		}
	} else {
		return allCoordinate, unitCoordinate, structureCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitCoordinate, structureCoordinate, nil
}

