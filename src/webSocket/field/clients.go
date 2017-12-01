package field

import (
	"../../game"
	"strconv"
)

type Clients struct { // структура описывающая клиента ws соеденение
	Login        	  string
	Id           	  int
	Watch             map[int]map[int]*game.Coordinate // map[X]map[Y]
	Units        	  map[int]map[int]*game.Unit       // map[X]map[Y]
	Structure    	  map[int]map[int]*game.Structure  // map[X]map[Y]
	HostileStructure  map[int]map[int]*game.Structure  // map[X]map[Y]
	HostileUnits 	  map[int]map[int]*game.Unit       // map[X]map[Y]
	Respawn      	  *game.Structure
	CreateZone   	  map[string]*game.Coordinate
	GameID            int
}

func (client *Clients) getLogin() string  {
	return client.Login
}

func (client *Clients) getAllWatchObject(activeGame *ActiveGame) {

	for _, xLine := range activeGame.getUnits() {
		for _, unit := range xLine {
			watchCoordinate, watchUnit, watchStructure, err := game.Watch(unit, client.Login, activeGame.getUnits(), activeGame.getStructure())//PermissionCoordinates(client, unit, units)

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
					_, ok := activeGame.getMap().OneLayerMap[coordinate.X][coordinate.Y]
					if !ok {
						client.addCoordinate(coordinate)
					}
				}
			}
		}
	}

	for _, xLine := range activeGame.getStructure() {
		for _, structure := range xLine {

			watchCoordinate, watchUnit, watchStructure, err := game.Watch(structure, client.Login, activeGame.getUnits(), activeGame.getStructure())

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
					_, ok := activeGame.getMap().OneLayerMap[coordinate.X][coordinate.Y]
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

func (client *Clients) addCoordinate(coordinate *game.Coordinate) {
	if client.Watch != nil {
		if client.Watch[coordinate.X] != nil {
			client.Watch[coordinate.X][coordinate.Y] = coordinate
		} else {
			client.Watch[coordinate.X] = make(map[int]*game.Coordinate)
			client.addCoordinate(coordinate)
		}
	} else {
		client.Watch = make(map[int]map[int]*game.Coordinate)
		client.addCoordinate(coordinate)
	}
}

func (client *Clients) addUnit(unit *game.Unit) {
	if client.Units != nil {
		if client.Units[unit.X] != nil {
			client.Units[unit.X][unit.Y] = unit
		} else {
			client.Units[unit.X] = make(map[int]*game.Unit)
			client.addUnit(unit)
		}
	} else {
		client.Units = make(map[int]map[int]*game.Unit)
		client.addUnit(unit)
	}
}

func (client *Clients) addHostileUnit(hostile *game.Unit) {
	if client.HostileUnits != nil {
		if client.HostileUnits[hostile.X] != nil {
			client.HostileUnits[hostile.X][hostile.Y] = hostile
		} else {
			client.HostileUnits[hostile.X] = make(map[int]*game.Unit)
			client.addHostileUnit(hostile)
		}
	} else {
		client.HostileUnits = make(map[int]map[int]*game.Unit)
		client.addHostileUnit(hostile)
	}
}

func (client *Clients) addStructure(structure *game.Structure) {
	if client.Structure != nil {
		if client.Structure[structure.X] != nil {
			client.Structure[structure.X][structure.Y] = structure
		} else {
			client.Structure[structure.X] = make(map[int]*game.Structure)
			client.addStructure(structure)
		}
	} else {
		client.Structure = make(map[int]map[int]*game.Structure)
		client.addStructure(structure)
	}
}

func (client *Clients) addHostileStructure(structure *game.Structure) {
	if client.HostileStructure != nil {
		if client.HostileStructure[structure.X] != nil {
			client.HostileStructure[structure.X][structure.Y] = structure
		} else {
			client.HostileStructure[structure.X] = make(map[int]*game.Structure)
			client.addHostileStructure(structure)
		}
	} else {
		client.HostileStructure = make(map[int]map[int]*game.Structure)
		client.addHostileStructure(structure)
	}
}

func (client *Clients) setRespawn(respawn *game.Structure)  {
	PermCoordinates := game.GetCoordinates(respawn.X, respawn.Y, respawn.WatchZone)
	client.CreateZone = make(map[string]*game.Coordinate)
	for _, coordinate := range PermCoordinates {
		if !(coordinate.X == respawn.X && coordinate.Y == respawn.Y) {
			client.CreateZone[strconv.Itoa(coordinate.X)+":"+strconv.Itoa(coordinate.Y)] = coordinate
		}
	}

	client.Respawn = respawn
}
