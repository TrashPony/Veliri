package game

import (
	"strconv"
)

type Player struct {
	// структура описывающая клиента ws соеденение
	login              string
	id                 int
	watch              map[string]map[string]*Coordinate // map[X]map[Y]
	unitStorage        []*Unit
	units              map[string]map[string]*Unit // map[X]map[Y]
	matherShip         *MatherShip
	hostileMatherShips map[string]map[string]*MatherShip // map[X]map[Y]
	hostileUnits       map[string]map[string]*Unit       // map[X]map[Y]
	respawn            *Coordinate
	createZone         []*Coordinate
	gameID             int
	equips             []*Equip
	ready              bool
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
	client.matherShip = nil
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
		if client.watch[strconv.Itoa(coordinate.X)] != nil {
			client.watch[strconv.Itoa(coordinate.X)][strconv.Itoa(coordinate.Y)] = coordinate
		} else {
			client.watch[strconv.Itoa(coordinate.X)] = make(map[string]*Coordinate)
			client.AddCoordinate(coordinate)
		}
	} else {
		client.watch = make(map[string]map[string]*Coordinate)
		client.AddCoordinate(coordinate)
	}
}

func (client *Player) AddUnit(unit *Unit) {
	if client.units != nil {
		if client.units[strconv.Itoa(unit.X)] != nil {
			client.units[strconv.Itoa(unit.X)][strconv.Itoa(unit.Y)] = unit
		} else {
			client.units[strconv.Itoa(unit.X)] = make(map[string]*Unit)
			client.AddUnit(unit)
		}
	} else {
		client.units = make(map[string]map[string]*Unit)
		client.AddUnit(unit)
	}
}

func (client *Player) AddHostileUnit(hostile *Unit) {
	if client.hostileUnits != nil {
		if client.hostileUnits[strconv.Itoa(hostile.X)] != nil {
			client.hostileUnits[strconv.Itoa(hostile.X)][strconv.Itoa(hostile.Y)] = hostile
		} else {
			client.hostileUnits[strconv.Itoa(hostile.X)] = make(map[string]*Unit)
			client.AddHostileUnit(hostile)
		}
	} else {
		client.hostileUnits = make(map[string]map[string]*Unit)
		client.AddHostileUnit(hostile)
	}
}

func (client *Player) AddMatherShips(matherShip *MatherShip) {
	client.matherShip = matherShip
}

func (client *Player) AddHostileMatherShip(matherShip *MatherShip) {
	if client.hostileMatherShips != nil {
		if client.hostileMatherShips[strconv.Itoa(matherShip.X)] != nil {
			client.hostileMatherShips[strconv.Itoa(matherShip.X)][strconv.Itoa(matherShip.Y)] = matherShip
		} else {
			client.hostileMatherShips[strconv.Itoa(matherShip.X)] = make(map[string]*MatherShip)
			client.AddHostileMatherShip(matherShip)
		}
	} else {
		client.hostileMatherShips = make(map[string]map[string]*MatherShip)
		client.AddHostileMatherShip(matherShip)
	}
}

func (client *Player) SetRespawn(respawn *Coordinate) {
	client.respawn = respawn
}

func (client *Player) GetCreateZone() ([]*Coordinate) {
	tmpCoordiantes := GetCoordinates(client.matherShip.X, client.matherShip.Y,client.matherShip.RangeView)

	client.createZone = make([]*Coordinate, 0)

	for _, coordinate := range tmpCoordiantes {
		if coordinate.X >= 0 && coordinate.Y >= 0 {
			client.createZone = append(client.createZone, coordinate)
		}
	}

	return client.createZone
}

func (client *Player) GetRespawn() (respawn *Coordinate) {
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

func (client *Player) GetUnits() (unit map[string]map[string]*Unit) {
	return client.units
}

func (client *Player) GetUnit(x, y int) (unit *Unit, find bool) {
	unit, find = client.units[strconv.Itoa(x)][strconv.Itoa(y)]
	return
}

func (client *Player) DelUnit(x, y int) {
	delete(client.units[strconv.Itoa(x)], strconv.Itoa(y))
}

func (client *Player) GetHostileUnits() (unit map[string]map[string]*Unit) {
	return client.hostileUnits
}

func (client *Player) GetHostileUnit(x, y int) (unit *Unit, find bool) {
	unit, find = client.hostileUnits[strconv.Itoa(x)][strconv.Itoa(y)]
	return
}

func (client *Player) DelHostileUnit(x, y int) {
	delete(client.hostileUnits[strconv.Itoa(x)], strconv.Itoa(y))
}

func (client *Player) GetMatherShip() (*MatherShip) {
	return client.matherShip
}

func (client *Player) GetHostileMatherShips() (matherShip map[string]map[string]*MatherShip) {
	return client.hostileMatherShips
}

func (client *Player) GetHostileMatherShip(x, y int) (matherShip *MatherShip, find bool) {
	matherShip, find = client.hostileMatherShips[strconv.Itoa(x)][strconv.Itoa(y)]
	return
}

func (client *Player) GetWatchCoordinates() (coordinate map[string]map[string]*Coordinate) {
	return client.watch
}

func (client *Player) GetWatchCoordinate(x, y int) (coordinate *Coordinate, find bool) {
	coordinate, find = client.watch[strconv.Itoa(x)][strconv.Itoa(y)]
	return
}

func (client *Player) DelWatchCoordinate(x, y int) {
	delete(client.watch[strconv.Itoa(x)], strconv.Itoa(y))
}

func (client *Player) SetEquip(equips []*Equip) {
	client.equips = equips
}

func (client *Player) GetEquip() []*Equip {
	return client.equips
}

func (client *Player) SetReady(ready bool) {
	client.ready = ready
}

func (client *Player) GetReady() (bool) {
	return client.ready
}

func (client *Player) SetUnitsStorage(units []*Unit) () {
	client.unitStorage = units
}

func (client *Player) GetUnitsStorage() (unit []*Unit) {
	return client.unitStorage
}

func (client *Player) GetUnitStorage(id int) (unit *Unit, find bool) {
	for _, unit := range client.GetUnitsStorage() {
		if id == unit.Id {
			return unit, true
		}
	}

	return
}
