package player

import (
	"../equip"
	"../matherShip"
	"../unit"
	"../coordinate"
)

type Player struct {
	// структура описывающая клиента ws соеденение
	login              string
	id                 int
	watch              map[string]map[string]*coordinate.Coordinate // map[X]map[Y]
	unitStorage        []*unit.Unit
	units              map[string]map[string]*unit.Unit // map[X]map[Y]
	matherShip         *matherShip.MatherShip
	hostileMatherShips map[string]map[string]*matherShip.MatherShip // map[X]map[Y]
	hostileUnits       map[string]map[string]*unit.Unit       // map[X]map[Y]
	respawn            *coordinate.Coordinate
	createZone         []*coordinate.Coordinate
	gameID             int
	equips             []*equip.Equip
	ready              bool
}

func (client *Player) SetRespawn(respawn *coordinate.Coordinate) {
	client.respawn = respawn
}

func (client *Player) GetCreateZone() ([]*coordinate.Coordinate) {
	tmpCoordinates := coordinate.GetCoordinatesRadius(client.matherShip.X, client.matherShip.Y, client.matherShip.RangeView)

	client.createZone = make([]*coordinate.Coordinate, 0)

	for _, gameCoordinate := range tmpCoordinates {
		if gameCoordinate.X >= 0 && gameCoordinate.Y >= 0 {
			client.createZone = append(client.createZone, gameCoordinate)
		}
	}

	return client.createZone
}

func (client *Player) GetRespawn() (respawn *coordinate.Coordinate) {
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

func (client *Player) SetReady(ready bool) {
	client.ready = ready
}

func (client *Player) GetReady() (bool) {
	return client.ready
}
