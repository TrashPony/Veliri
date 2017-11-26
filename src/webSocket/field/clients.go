package field

import (
	"../../game/objects"
	"strconv"
)

type Clients struct { // структура описывающая клиента ws соеденение
	Login        string
	Id           int
	Watch        map[int]map[int]*objects.Coordinate // map[X]map[Y]
	Units        map[int]map[int]*objects.Unit       // map[X]map[Y]
	Structure    map[int]map[int]*objects.Structure  // map[X]map[Y]
	HostileUnits map[int]map[int]*objects.Unit       // map[X]map[Y]
	Respawn      *objects.Structure
	CreateZone   map[string]*objects.Coordinate
	GameID       int
}

func (client *Clients) getAllWatchObject(units map[int]map[int]*objects.Unit, structures map[int]map[int]*objects.Structure) {
	for _, xLine := range units {
		for _, unit := range xLine {
			watchCoordinate, watchUnit, watchStructure, err := unit.Watch(client.Login, units, structures)//PermissionCoordinates(client, unit, units)

			if err != nil { // если крип не мой то пропускаем дальнейшее действие
				continue
			} else {
				client.addUnit(unit)

				for _, xLine := range watchUnit {
					for _, hostile := range xLine {
						if hostile.NameUser != client.Login {
							client.addHostileUnit(hostile)
						}
					}
				}

				for _, xLine := range watchStructure {
					for _, hostile := range xLine {
						if hostile.NameUser != client.Login {
							// TODO хостаил структуры
						}
					}
				}

				for _, coordinate := range watchCoordinate {
					client.addCoordinate(coordinate)
				}
			}
		}
	}

	for _, xLine := range structures {
		for _, structure := range xLine {
			watchCoordinate, watchUnit, watchStructure, err := structure.Watch(client.Login, units, structures)//PermissionCoordinates(client, unit, units)

			if err != nil { // если крип не мой то пропускаем дальнейшее действие
				continue
			} else {
				client.addStructure(structure)

				for _, xLine := range watchUnit {
					for _, hostile := range xLine {
						if hostile.NameUser != client.Login {
							client.addHostileUnit(hostile)
						}
					}
				}

				for _, xLine := range watchStructure {
					for _, hostile := range xLine {
						if hostile.NameUser != client.Login {
							// TODO хостаил структуры
						}
					}
				}

				for _, coordinate := range watchCoordinate {
					client.addCoordinate(coordinate)
				}
			}
		}
	}
}

// отправляем открытые ячейки, удаляем закрытые
func (client *Clients) updateWatchZone(units map[int]map[int]*objects.Unit, structures map[int]map[int]*objects.Structure) {

	oldWatchZone := client.Watch
	oldWatchUnit := client.HostileUnits

	client.Units = nil
	client.HostileUnits = nil
	client.Watch = nil

	client.getAllWatchObject(units, structures)

	updateOpenCoordinate(client, oldWatchZone)
	updateHostileUnit(client, oldWatchUnit)
}

func updateOpenCoordinate(client *Clients, oldWatchZone map[int]map[int]*objects.Coordinate) {
	for _, xLine := range client.Watch { // отправляем все новые координаты, и т.к. старая клетка юнита теперь тоже является координатой то и ее тоже обновляем
		for _, newCoordinate := range xLine {
			_, ok := oldWatchZone[newCoordinate.X][newCoordinate.Y]
			if !ok {
				client.addCoordinate(newCoordinate)
				openCoordinate(client.Login, newCoordinate.X, newCoordinate.Y)
			}
		}
	}

	for _, xLine := range oldWatchZone { // удаляем старые координаты из зоны видимости
		for _, oldCoordinate := range xLine {
			_, find := client.Watch[oldCoordinate.X][oldCoordinate.Y]
			if !find {
				delete(client.Watch[oldCoordinate.X], oldCoordinate.Y)
				closeCoordinate(client.Login, oldCoordinate.X, oldCoordinate.Y)
			}
		}
	}
}

func updateHostileUnit(client *Clients, oldWatchUnit map[int]map[int]*objects.Unit) {
	for _, xLine := range client.HostileUnits { // добавляем новые вражеские юниты которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchUnit[hostile.X][hostile.Y]
			if !ok {
				client.addHostileUnit(hostile)
				var unitsParameter InitUnit
				unitsParameter.initUnit(hostile, client.Login)
			}
		}
	}

	for _, xLine := range oldWatchUnit {
		for _, hostile := range xLine {
			_, find := client.HostileUnits[hostile.X][hostile.Y]
			if !find {
				closeCoordinate(client.Login, hostile.X, hostile.Y)
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

func (client *Clients) addStructure(structure *objects.Structure) {
	if client.Structure != nil {
		if client.Structure[structure.X] != nil {
			client.Structure[structure.X][structure.Y] = structure
		} else {
			client.Structure[structure.X] = make(map[int]*objects.Structure)
			client.addStructure(structure)
		}
	} else {
		client.Structure = make(map[int]map[int]*objects.Structure)
		client.addStructure(structure)
	}
}

func (client *Clients) setRespawn(respawn *objects.Structure)  {
	PermCoordinates := objects.GetCoordinates(respawn.X, respawn.Y, respawn.WatchZone)
	client.CreateZone = make(map[string]*objects.Coordinate)
	for _, coordinate := range PermCoordinates {
		if !(coordinate.X == respawn.X && coordinate.Y == respawn.Y) {
			client.CreateZone[strconv.Itoa(coordinate.X)+":"+strconv.Itoa(coordinate.Y)] = coordinate
		}
	}

	client.Respawn = respawn
}
