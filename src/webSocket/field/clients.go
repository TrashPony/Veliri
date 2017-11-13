package field

import "../../game/objects"

type Clients struct { // структура описывающая клиента ws соеденение
	Login string
	Id int
	Watch map[int]map[int]*objects.Coordinate  // map[X]map[Y]
	Units map[int]map[int]*objects.Unit        // map[X]map[Y]
	HostileUnits map[int]map[int]*objects.Unit // map[X]map[Y]
	Map objects.Map
	Respawn objects.Respawn
	CreateZone map[string]*objects.Coordinate
	GameStat objects.Game
	Players []objects.UserStat
}

func (client *Clients) getAllWatchObject(units map[string]*objects.Unit) {
	client.Units = nil
	client.HostileUnits = nil
	client.Watch = nil

	for _, unit := range units {

		watchCoordinate, watchUnit, err := PermissionCoordinates(client, unit, units)

		if err != nil {
			continue
		}

		for _, xLine := range watchUnit {
			for _, hostile := range xLine {
				if hostile.NameUser != client.Login {
					client.addHostileUnit(hostile)
				} else {
					client.addUnit(unit)
				}
			}
		}

		for _, coordinate := range watchCoordinate {
			client.addCoordinate(coordinate)
		}
	}
}

func (client *Clients) addCoordinate(coordinate *objects.Coordinate) {
	if client.Watch != nil {
		if client.Watch[coordinate.X] != nil {
			client.Watch[coordinate.X][coordinate.Y] = coordinate
		} else {
			client.Watch[coordinate.X] = make(map[int]*objects.Coordinate)
			client.addCoordinate(coordinate)
		}
	} else {
		client.Watch = make(map[int]map[int]*objects.Coordinate)
		client.addCoordinate(coordinate)
	}
}

func (client *Clients) addUnit(unit *objects.Unit) {
	if client.Units != nil {
		if client.Units[unit.X] != nil {
			client.Units[unit.X][unit.Y] = unit
		} else {
			client.Units[unit.X] = make(map[int]*objects.Unit)
			client.addUnit(unit)
		}
	} else {
		client.Units = make(map[int]map[int]*objects.Unit)
		client.addUnit(unit)
	}
}

func (client *Clients) addHostileUnit(hostile *objects.Unit) {
	if client.HostileUnits != nil {
		if client.HostileUnits[hostile.X] != nil {
			client.HostileUnits[hostile.X][hostile.Y] = hostile
		} else {
			client.HostileUnits[hostile.X] = make(map[int]*objects.Unit)
			client.addHostileUnit(hostile)
		}
	} else {
		client.HostileUnits = make(map[int]map[int]*objects.Unit)
		client.addHostileUnit(hostile)
	}
}