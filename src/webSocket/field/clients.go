package field

import (
	"../../game/objects"
	"strconv"
)

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
	client.Units = nil // TODO переделать на работу с сылками
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

// отправляем открытые ячейки, удаляем закрытые
func (client *Clients) updateWatchZone(units map[string]*objects.Unit) {

	oldWatchZone := client.Watch
	oldWatchUnit := client.HostileUnits

	client.getAllWatchObject(units)

	updateOpenCoordinate(client, oldWatchZone)
	updateHostileUnit(client, oldWatchUnit)
}

func updateOpenCoordinate(client *Clients, oldWatchZone map[int]map[int]*objects.Coordinate)  {
	for _, xLine := range client.Watch { // отправляем все новые координаты, и т.к. старая клетка юнита теперь тоже является координатой то и ее тоже обновляем
		for _, newCoordinate := range xLine {
			_, ok := oldWatchZone[newCoordinate.X][newCoordinate.Y]
			if !ok {
				client.addCoordinate(newCoordinate)
				resp := Coordinate{Event: "OpenCoordinate", UserName: client.Login, X: newCoordinate.X, Y: newCoordinate.Y}
				coordiante <- resp
			}
		}
	}

	for _, xLine := range oldWatchZone { // удаляем старые координаты из зоны видимости
		for _, oldCoordinate := range xLine {
			_, find := client.Watch[oldCoordinate.X][oldCoordinate.Y]
			if !find {
				delete(client.Watch[oldCoordinate.X], oldCoordinate.Y)
				resp := Coordinate{Event: "DellCoordinate", UserName: client.Login, X: oldCoordinate.X, Y: oldCoordinate.Y} // удаляем старое поле доступа
				coordiante <- resp
			}
		}
	}
}

func updateHostileUnit(client *Clients, oldWatchUnit map[int]map[int]*objects.Unit)  {
	for _, xLine := range client.HostileUnits { // добавляем новые вражеские юниты которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchUnit[hostile.X][hostile.Y]
			if !ok {
				client.addHostileUnit(hostile)
				var unitsParametr = InitUnit{Event: "InitUnit", UserName: client.Login, TypeUnit: hostile.NameType, UserOwned: hostile.NameUser,
					HP: hostile.Hp, UnitAction: strconv.FormatBool(hostile.Action), Target: hostile.Target, X: hostile.X, Y: hostile.Y}
				initUnit <- unitsParametr
			}
		}
	}

	for _, xLine := range oldWatchUnit {
		for _, hostile := range xLine {
			_, find := client.HostileUnits[hostile.X][hostile.Y]
			if !find {
				resp := Coordinate{Event: "DellCoordinate", UserName: client.Login, X: hostile.X, Y: hostile.Y} // удаляем старое поле доступа
				coordiante <- resp
			}
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