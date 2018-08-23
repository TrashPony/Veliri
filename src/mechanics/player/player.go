package player

import (
	"../gameObjects/equip"
	"../gameObjects/matherShip"
	"../gameObjects/unit"
	"../gameObjects/squad"
	"../gameObjects/coordinate"
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
	hostileUnits       map[string]map[string]*unit.Unit             // map[X]map[Y]
	Respawn            *coordinate.Coordinate
	createZone         map[string]map[string]*coordinate.Coordinate
	gameID             int
	equips             []*equip.Equip
	Ready              bool

	LobbyReady		   bool
	squad  			   *squad.Squad
	squads 			   []*squad.Squad
}

func (client *Player) SetRespawn(respawn *coordinate.Coordinate) {
	client.Respawn = respawn
}

func (client *Player) SetCreateZone(zone map[string]map[string]*coordinate.Coordinate) () {
	client.createZone = zone
}

func (client *Player) GetCreateZone() (map[string]map[string]*coordinate.Coordinate) {
	return client.createZone
}

func (client *Player) GetRespawn() (respawn *coordinate.Coordinate) {
	return client.Respawn
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
	client.Ready = ready
}

func (client *Player) GetReady() (bool) {
	return client.Ready
}

func (client *Player) SetLobbyReady(ready bool) {
	client.LobbyReady = ready
}

func (client *Player) GetLobbyReady() (bool) {
	return client.LobbyReady
}

func (client *Player) GetSquad() (*squad.Squad) {
	return client.squad
}

func (client *Player) SetSquad(squad *squad.Squad) () {
	client.squad = squad
}

func (client *Player) GetSquads() ([]*squad.Squad) {
	return client.squads
}

func (client *Player) SetSquads(squads []*squad.Squad) () {
	client.squads = squads
}
