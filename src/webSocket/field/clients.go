package field

import (
	"../../game/objects"
	"strconv"
)

type Clients struct { // структура описывающая клиента ws соеденение
	Login        	  string
	Id           	  int
	Watch             map[int]map[int]*objects.Coordinate // map[X]map[Y]
	Units        	  map[int]map[int]*objects.Unit       // map[X]map[Y]
	Structure    	  map[int]map[int]*objects.Structure  // map[X]map[Y]
	HostileStructure  map[int]map[int]*objects.Structure  // map[X]map[Y]
	HostileUnits 	  map[int]map[int]*objects.Unit       // map[X]map[Y]
	Respawn      	  *objects.Structure
	CreateZone   	  map[string]*objects.Coordinate
	GameID            int
}

func (client *Clients) getAllWatchObject(game *ActiveGame) {
	for _, xLine := range game.getUnits() {
		for _, unit := range xLine {
			watchCoordinate, watchUnit, watchStructure, err := unit.Watch(client.Login, game.getUnits(), game.getStructure())//PermissionCoordinates(client, unit, units)

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
							client.addHostileStructure(hostile)
						}
					}
				}

				for _, coordinate := range watchCoordinate {
					_, ok := game.coordinate[coordinate.X][coordinate.Y]
					if !ok {
						client.addCoordinate(coordinate)
					}
				}
			}
		}
	}

	for _, xLine := range game.getStructure() {
		for _, structure := range xLine {
			watchCoordinate, watchUnit, watchStructure, err := structure.Watch(client.Login, game.getUnits(), game.getStructure())

			if err != nil { // если структура не моя то пропускаем дальнейшее действие
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
							client.addHostileStructure(hostile)
						}
					}
				}

				for _, coordinate := range watchCoordinate {
					_, ok := game.coordinate[coordinate.X][coordinate.Y]
					if !ok {
						client.addCoordinate(coordinate)
					}
				}
			}
		}
	}
}

// отправляем открытые ячейки, удаляем закрытые
func (client *Clients) updateWatchZone(game *ActiveGame) {

	oldWatchZone := client.Watch
	oldWatchHostileUnits := client.HostileUnits
	oldWatchHostileStructure := client.HostileStructure

	client.Units = nil
	client.Structure = nil
	client.HostileUnits = nil
	client.HostileStructure = nil
	client.Watch = nil

	client.getAllWatchObject(game)

	updateMyUnit(client)
	updateMyStructure(client)
	updateHostileUnit(client, oldWatchHostileUnits)
	updateHostileStrcuture(client, oldWatchHostileStructure)
	updateOpenCoordinate(client, oldWatchZone)
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

func (client *Clients) addHostileStructure(structure *objects.Structure) {
	if client.HostileStructure != nil {
		if client.HostileStructure[structure.X] != nil {
			client.HostileStructure[structure.X][structure.Y] = structure
		} else {
			client.HostileStructure[structure.X] = make(map[int]*objects.Structure)
			client.addHostileStructure(structure)
		}
	} else {
		client.HostileStructure = make(map[int]map[int]*objects.Structure)
		client.addHostileStructure(structure)
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
