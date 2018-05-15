package player

import (
	"../unit"
	"strconv"
)

func (client *Player) AddUnit(gameUnit *unit.Unit) {
	if client.units != nil {
		if client.units[strconv.Itoa(gameUnit.X)] != nil {
			client.units[strconv.Itoa(gameUnit.X)][strconv.Itoa(gameUnit.Y)] = gameUnit
		} else {
			client.units[strconv.Itoa(gameUnit.X)] = make(map[string]*unit.Unit)
			client.AddUnit(gameUnit)
		}
	} else {
		client.units = make(map[string]map[string]*unit.Unit)
		client.AddUnit(gameUnit)
	}
}

func (client *Player) AddHostileUnit(hostile *unit.Unit) {
	if client.hostileUnits != nil {
		if client.hostileUnits[strconv.Itoa(hostile.X)] != nil {
			client.hostileUnits[strconv.Itoa(hostile.X)][strconv.Itoa(hostile.Y)] = hostile
		} else {
			client.hostileUnits[strconv.Itoa(hostile.X)] = make(map[string]*unit.Unit)
			client.AddHostileUnit(hostile)
		}
	} else {
		client.hostileUnits = make(map[string]map[string]*unit.Unit)
		client.AddHostileUnit(hostile)
	}
}

func (client *Player) GetUnits() (units map[string]map[string]*unit.Unit) {
	return client.units
}

func (client *Player) SetUnits(units map[string]map[string]*unit.Unit) () {
	client.units = units
}

func (client *Player) GetUnit(x, y int) (gameUnit *unit.Unit, find bool) {
	gameUnit, find = client.units[strconv.Itoa(x)][strconv.Itoa(y)]
	return
}

func (client *Player) DelUnit(x, y int) {
	delete(client.units[strconv.Itoa(x)], strconv.Itoa(y))
}

func (client *Player) GetHostileUnits() (units map[string]map[string]*unit.Unit) {
	return client.hostileUnits
}

func (client *Player) SetHostileUnits(units map[string]map[string]*unit.Unit) () {
	client.hostileUnits = units
}

func (client *Player) GetHostileUnit(x, y int) (gameUnit *unit.Unit, find bool) {
	gameUnit, find = client.hostileUnits[strconv.Itoa(x)][strconv.Itoa(y)]
	return
}

func (client *Player) DelHostileUnit(x, y int) {
	delete(client.hostileUnits[strconv.Itoa(x)], strconv.Itoa(y))
}

func (client *Player) SetUnitsStorage(units []*unit.Unit) () {
	client.unitStorage = units
}

func (client *Player) GetUnitsStorage() (gameUnit []*unit.Unit) {
	return client.unitStorage
}

func (client *Player) GetUnitStorage(id int) (storageUnit *unit.Unit, find bool) {
	for _, storageUnit := range client.GetUnitsStorage() {
		if id == storageUnit.Id {
			return storageUnit, true
		}
	}

	return
}