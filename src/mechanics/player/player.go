package player

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

type Player struct {
	id              int
	login           string
	email           string
	credits         int
	experiencePoint int

	squad  *squad.Squad   // отряд игрока
	squads []*squad.Squad // не активные отряды которые ждут игрока на безах

	unitStorage []*unit.Unit // юниты которы находяться не на поле игры в трюме мса

	watch              map[string]map[string]*coordinate.Coordinate // map[X]map[Y] координаты которые видит пользватель
	units              map[string]map[string]*unit.Unit             // map[X]map[Y] свои юниты представленные ввиде карты на поле
	hostileUnits       map[string]map[string]*unit.Unit             // map[X]map[Y] вражеские юниты которы видно в настоящее время
	memoryHostileUnits map[string]unit.Unit                         // Юниты которые видел и запомнил пользователь за всю игру

	gameID int
	Ready  bool

	LobbyReady bool
	Respawn    *coordinate.Coordinate
	InBaseID   int // ид базы в которой сидит игрок
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

func (client *Player) GetSquadsByID(ID int) *squad.Squad {
	for _, userSquad := range client.squads {
		if userSquad != nil && userSquad.ID == ID {
			return userSquad
		}
	}

	return nil
}

func (client *Player) RemoveSquadsByID(ID int) {
	for i, userSquad := range client.squads {
		if userSquad.ID == ID {
			client.squads[i] = nil
		}
	}
}

func (client *Player) GetSquadsByBaseID(BaseID int) []*squad.Squad {
	squads := make([]*squad.Squad, 0)

	for _, userSquad := range client.squads {
		if client.GetSquad() != nil {
			if userSquad != nil && userSquad.BaseID == BaseID && client.GetSquad().ID != userSquad.ID {
				squads = append(squads, userSquad)
			}
		} else {
			if userSquad != nil && userSquad.BaseID == BaseID {
				squads = append(squads, userSquad)
			}
		}
	}

	return squads
}

func (client *Player) SetSquads(squads []*squad.Squad) {
	client.squads = squads
}

func (client *Player) SetEmail(email string) {
	client.email = email
}

func (client *Player) GetEmail() string {
	return client.email
}

func (client *Player) SetCredits(credits int) {
	client.credits = credits
}

func (client *Player) GetCredits() int {
	return client.credits
}

func (client *Player) SetExperiencePoint(experiencePoint int) {
	client.experiencePoint = experiencePoint
}

func (client *Player) GetExperiencePoint() int {
	return client.experiencePoint
}
