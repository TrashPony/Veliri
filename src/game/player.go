package game

import (
	"strconv"
)

type Player struct {
	// структура описывающая клиента ws соеденение
	login              string
	id                 int
	watch              map[int]map[int]*Coordinate // map[X]map[Y]
	units              map[int]map[int]*Unit       // map[X]map[Y]
	MatherShip         *MatherShip
	hostileMatherShips map[int]map[int]*MatherShip // map[X]map[Y]
	hostileUnits       map[int]map[int]*Unit       // map[X]map[Y]
	respawn            *MatherShip
	createZone         map[string]*Coordinate
	gameID             int
}

func (client *Player) GetAllWatchObject(activeGame *Game) {

	for _, xLine := range activeGame.GetUnits() {
		for _, unit := range xLine {
			watchCoordinate, watchUnit, watchMatherShip, err := Watch(unit, client.login, activeGame) //PermissionCoordinates(client, unit, units)

			if err != nil { // если крип не мой то пропускаем дальнейшее действие
				continue
			} else {
				client.AddUnit(unit)

				for _, xLine := range watchUnit {
					for _, hostile := range xLine {
						if hostile.Owner != client.login {
							client.AddHostileUnit(hostile)
						}
					}
				}

				for _, xLine := range watchMatherShip {
					for _, hostile := range xLine {
						if hostile.Owner != client.login {
							client.AddHostileMatherShip(hostile)
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

	for _, xLine := range activeGame.GetMatherShips() {
		for _, matherShip := range xLine {

			watchCoordinate, watchUnit, watchMatherShip, err := Watch(matherShip, client.login, activeGame)

			if err != nil { // если структура не моя то пропускаем дальнейшее действие
				continue
			} else {
				client.AddMatherShips(matherShip)

				for _, xLine := range watchUnit {
					for _, hostile := range xLine {
						if hostile.Owner != client.login {
							client.AddHostileUnit(hostile)
						}
					}
				}

				for _, xLine := range watchMatherShip {
					for _, hostile := range xLine {
						if hostile.Owner != client.login {
							client.AddHostileMatherShip(hostile)
						}
					}
				}

				for _, coordinate := range watchCoordinate {
					//_, ok := activeGame.GetMap().OneLayerMap[coordinate.X][coordinate.Y]
					//if !ok { TODO инициализировать всю карту...
					client.AddCoordinate(coordinate)
					//
				}
			}
		}
	}
}

type UpdaterWatchZone struct {
	CloseCoordinate []*Coordinate `json:"close_coordinate"`
	OpenCoordinate  []*Coordinate `json:"open_coordinate"`
	OpenUnit        []*Unit       `json:"open_unit"`
	OpenMatherShip  []*MatherShip `json:"open_mather_ship"`
}

// отправляем открытые ячейки, удаляем закрытые
func (client *Player) UpdateWatchZone(game *Game) (*UpdaterWatchZone) {
	var updaterWatchZone UpdaterWatchZone

	oldWatchZone := client.GetWatchCoordinates()
	oldWatchHostileUnits := client.GetHostileUnits()
	oldWatchHostileMatherShips := client.GetHostileMatherShips()

	client.units = nil
	client.MatherShip = nil
	client.hostileUnits = nil
	client.hostileMatherShips = nil
	client.watch = nil

	client.GetAllWatchObject(game)

	openCoordinate, closeCoordinate := updateOpenCoordinate(client, oldWatchZone)
	openUnit, closeUnit := updateHostileUnit(client, oldWatchHostileUnits)
	openMatherShip, closeMatherShip := updateHostileStrcuture(client, oldWatchHostileMatherShips)

	sendCloseCoordinate := parseCloseCoordinate(closeCoordinate, closeUnit, closeMatherShip, game)

	updaterWatchZone.CloseCoordinate = sendCloseCoordinate
	updaterWatchZone.OpenCoordinate = openCoordinate
	updaterWatchZone.OpenUnit = openUnit
	updaterWatchZone.OpenMatherShip = openMatherShip

	return &updaterWatchZone
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

func (client *Player) AddMatherShips(matherShip *MatherShip) {
	client.MatherShip = matherShip
}

func (client *Player) AddHostileMatherShip(matherShip *MatherShip) {
	if client.hostileMatherShips != nil {
		if client.hostileMatherShips[matherShip.X] != nil {
			client.hostileMatherShips[matherShip.X][matherShip.Y] = matherShip
		} else {
			client.hostileMatherShips[matherShip.X] = make(map[int]*MatherShip)
			client.AddHostileMatherShip(matherShip)
		}
	} else {
		client.hostileMatherShips = make(map[int]map[int]*MatherShip)
		client.AddHostileMatherShip(matherShip)
	}
}

func (client *Player) SetRespawn(respawn *MatherShip) {
	PermCoordinates := GetCoordinates(respawn.X, respawn.Y, respawn.RangeView)
	client.createZone = make(map[string]*Coordinate)
	for _, coordinate := range PermCoordinates {
		if !(coordinate.X == respawn.X && coordinate.Y == respawn.Y) {
			client.createZone[strconv.Itoa(coordinate.X)+":"+strconv.Itoa(coordinate.Y)] = coordinate
		}
	}

	client.respawn = respawn
}

func (client *Player) GetCreateZone() (map[string]*Coordinate) {
	return client.createZone
}

func (client *Player) GetRespawn() (respawn *MatherShip) {
	return client.respawn
}

func (client *Player) SetLogin(login string) {
	client.login = login
}

func (client *Player) GetLogin() (login string) {
	return client.login
}

func (client *Player) SetID(id int) {
	client.id = id
}

func (client *Player) GetID() (id int) {
	return client.id
}

func (client *Player) SetGameID(id int) {
	client.gameID = id
}

func (client *Player) GetGameID() (id int) {
	return client.gameID
}

func (client *Player) GetUnits() (unit map[int]map[int]*Unit) {
	return client.units
}

func (client *Player) GetUnit(x, y int) (unit *Unit, find bool) {
	unit, find = client.units[x][y]
	return
}

func (client *Player) DelUnit(x, y int) {
	delete(client.units[x], y)
}

func (client *Player) GetHostileUnits() (unit map[int]map[int]*Unit) {
	return client.hostileUnits
}

func (client *Player) GetHostileUnit(x, y int) (unit *Unit, find bool) {
	unit, find = client.hostileUnits[x][y]
	return
}

func (client *Player) DelHostileUnit(x, y int) {
	delete(client.hostileUnits[x], y)
}

func (client *Player) GetMatherShip() (*MatherShip) {
	return client.MatherShip
}

func (client *Player) GetHostileMatherShips() (matherShip map[int]map[int]*MatherShip) {
	return client.hostileMatherShips
}

func (client *Player) GetHostileMatherShip(x, y int) (matherShip *MatherShip, find bool) {
	matherShip, find = client.hostileMatherShips[x][y]
	return
}

func (client *Player) GetWatchCoordinates() (coordinate map[int]map[int]*Coordinate) {
	return client.watch
}

func (client *Player) GetWatchCoordinate(x, y int) (coordinate *Coordinate, find bool) {
	coordinate, find = client.watch[x][y]
	return
}

func (client *Player) DelWatchCoordinate(x, y int) {
	delete(client.watch[x], y)
}
