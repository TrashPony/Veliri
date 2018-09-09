package player

import (
	"../gameObjects/coordinate"
	"../gameObjects/squad"
	"../gameObjects/unit"
)

type Player struct {
	// структура описывающая клиента ws соеденение
	login string
	id    int

	watch        map[string]map[string]*coordinate.Coordinate // map[X]map[Y]
	unitStorage  []*unit.Unit
	units        map[string]map[string]*unit.Unit // map[X]map[Y]
	squad        *squad.Squad
	Ready        bool
	hostileUnits map[string]map[string]*unit.Unit // map[X]map[Y]
	gameID       int

	LobbyReady bool
	Respawn    *coordinate.Coordinate
	squads     []*squad.Squad
}

func (client *Player) SetRespawn(respawn *coordinate.Coordinate) {
	client.Respawn = respawn
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

func (client *Player) GetReady() bool {
	return client.Ready
}

func (client *Player) SetLobbyReady(ready bool) {
	client.LobbyReady = ready
}

func (client *Player) GetLobbyReady() bool {
	return client.LobbyReady
}

func (client *Player) GetSquad() *squad.Squad {
	return client.squad
}

func (client *Player) SetSquad(squad *squad.Squad) {
	client.squad = squad
}

func (client *Player) GetSquads() []*squad.Squad {
	return client.squads
}

func (client *Player) SetSquads(squads []*squad.Squad) {
	client.squads = squads
}
