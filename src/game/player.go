package game

import (
	"./"
	"strconv"
)

type Player struct { // структура описывающая клиента ws соеденение
	login        	  string
	id           	  int
	watch             map[int]map[int]*Coordinate // map[X]map[Y]
	units        	  map[int]map[int]*Unit       // map[X]map[Y]
	structure    	  map[int]map[int]*Structure  // map[X]map[Y]
	hostileStructure  map[int]map[int]*Structure  // map[X]map[Y]
	hostileUnits 	  map[int]map[int]*Unit       // map[X]map[Y]
	respawn      	  *Structure
	createZone   	  map[string]*Coordinate
	gameID            int
}

func (client *Player) GetAllWatchObject(activeGame *Game) {

	for _, xLine := range activeGame.GetUnits() {
		for _, unit := range xLine {
			watchCoordinate, watchUnit, watchStructure, err := Watch(unit, client.login, activeGame.GetUnits(), activeGame.GetStructure())//PermissionCoordinates(client, unit, units)

			if err != nil { // если крип не мой то пропускаем дальнейшее действие
				continue
			} else {
				client.AddUnit(unit)

				for _, xLine := range watchUnit {
					for _, hostile := range xLine {
						if hostile.NameUser != client.login {
							client.AddHostileUnit(hostile)
						}
					}
				}

				for _, xLine := range watchStructure {
					for _, hostile := range xLine {
						if hostile.NameUser != client.login {
							client.AddHostileStructure(hostile)
						}
					}
				}

				for _, coordinate := range watchCoordinate {
					_, ok := activeGame.GetMap().OneLayerMap[coordinate.X][coordinate.Y]
					if !ok {
						client.AddCoordinate(coordinate)
					}
				}
			}
		}
	}

	for _, xLine := range activeGame.GetStructure() {
		for _, structure := range xLine {

			watchCoordinate, watchUnit, watchStructure, err := Watch(structure, client.login, activeGame.GetUnits(), activeGame.GetStructure())

			if err != nil { // если структура не моя то пропускаем дальнейшее действие
				continue
			} else {
				client.AddStructure(structure)

				for _, xLine := range watchUnit {
					for _, hostile := range xLine {
						if hostile.NameUser != client.login {
							client.AddHostileUnit(hostile)
						}
					}
				}

				for _, xLine := range watchStructure {
					for _, hostile := range xLine {
						if hostile.NameUser != client.login {
							client.AddHostileStructure(hostile)
						}
					}
				}

				for _, coordinate := range watchCoordinate {
					_, ok := activeGame.GetMap().OneLayerMap	[coordinate.X][coordinate.Y]
					if !ok {
						client.AddCoordinate(coordinate)
					}
				}
			}
		}
	}
}

// отправляем открытые ячейки, удаляем закрытые
func (client *Player) updateWatchZone(game *Game) {

	//oldWatchZone := client.Watch
	//oldWatchHostileUnits := client.HostileUnits
	//oldWatchHostileStructure := client.HostileStructure
	// TODO
	client.units = nil
	client.structure = nil
	client.hostileUnits = nil
	client.hostileStructure = nil
	client.watch = nil

	client.GetAllWatchObject(game)

	//updateMyUnit(client)
	//updateMyStructure(client)
	//updateHostileUnit(client, oldWatchHostileUnits)
	//updateHostileStrcuture(client, oldWatchHostileStructure)
	//updateOpenCoordinate(client, oldWatchZone)
}

func (client *Player) AddCoordinate(coordinate *Coordinate) {
	if client.watch != nil {
		if client.watch[coordinate.X] != nil {
			client.watch[coordinate.X][coordinate.Y] = coordinate
		} else {
			client.watch[coordinate.X] = make(map[int]*Coordinate)
			client.AddCoordinate(coordinate)
		}
	} else {
		client.watch = make(map[int]map[int]*Coordinate)
		client.AddCoordinate(coordinate)
	}
}

func (client *Player) AddUnit(unit *Unit) {
	if client.units != nil {
		if client.units[unit.X] != nil {
			client.units[unit.X][unit.Y] = unit
		} else {
			client.units[unit.X] = make(map[int]*Unit)
			client.AddUnit(unit)
		}
	} else {
		client.units = make(map[int]map[int]*Unit)
		client.AddUnit(unit)
	}
}

func (client *Player) AddHostileUnit(hostile *Unit) {
	if client.hostileUnits != nil {
		if client.hostileUnits[hostile.X] != nil {
			client.hostileUnits[hostile.X][hostile.Y] = hostile
		} else {
			client.hostileUnits[hostile.X] = make(map[int]*Unit)
			client.AddHostileUnit(hostile)
		}
	} else {
		client.hostileUnits = make(map[int]map[int]*Unit)
		client.AddHostileUnit(hostile)
	}
}

func (client *Player) AddStructure(structure *Structure) {
	if client.structure != nil {
		if client.structure[structure.X] != nil {
			client.structure[structure.X][structure.Y] = structure
		} else {
			client.structure[structure.X] = make(map[int]*Structure)
			client.AddStructure(structure)
		}
	} else {
		client.structure = make(map[int]map[int]*Structure)
		client.AddStructure(structure)
	}
}

func (client *Player) AddHostileStructure(structure *Structure) {
	if client.hostileStructure != nil {
		if client.hostileStructure[structure.X] != nil {
			client.hostileStructure[structure.X][structure.Y] = structure
		} else {
			client.hostileStructure[structure.X] = make(map[int]*Structure)
			client.AddHostileStructure(structure)
		}
	} else {
		client.hostileStructure = make(map[int]map[int]*Structure)
		client.AddHostileStructure(structure)
	}
}

func (client *Player) SetRespawn(respawn *Structure)  {
	PermCoordinates := GetCoordinates(respawn.X, respawn.Y, respawn.WatchZone)
	client.createZone = make(map[string]*Coordinate)
	for _, coordinate := range PermCoordinates {
		if !(coordinate.X == respawn.X && coordinate.Y == respawn.Y) {
			client.createZone[strconv.Itoa(coordinate.X)+":"+strconv.Itoa(coordinate.Y)] = coordinate
		}
	}

	client.respawn = respawn
}

func (client *Player) SetLogin (login string) {
	client.login = login
}

func (client *Player) GetLogin()(login string) {
	return client.login
}

func (client *Player) SetID (id int) {
	client.id = id
}

func (client *Player) GetID () (id int) {
	return client.id
}

func (client *Player) SetGameID (id int) {
	client.gameID = id
}

func (client *Player) GetGameID () (id int) {
	return client.gameID
}

func (client *Player) GetUnits() (unit map[int]map[int]*Unit)  {
	return client.units
}

func (client *Player) GetUnit(x,y int) (unit *Unit, find bool)  {
	unit, find = client.units[x][y]
	return
}