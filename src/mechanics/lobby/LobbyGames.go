package lobby

import (
	"errors"
	"../player"
	"../localGame/map/coordinate"
)

type LobbyGames struct {
	ID       int
	Name     string
	Map      Map
	Creator  *player.Player
	Respawns []*coordinate.Coordinate
	Users    []*player.Player
}

func CreateNewLobbyGame(nameGame string, mapID int, creator *player.Player) LobbyGames {

	respawns := GetRespawns(mapID)
	mp := GetMap(mapID)

	game := LobbyGames{Name: nameGame, Map: mp, Creator: creator, Users: make([]*player.Player, 0), Respawns: respawns}
	game.Users = append(game.Users, creator)
	return game
}

func (game *LobbyGames) JoinToLobbyGame(user *player.Player) error {

	if len(game.Respawns) > len(game.Users) {
		game.Users = append(game.Users, user)
		return nil
	} else {
		return errors.New("lobby is full")
	}

	return errors.New("unknown error")
}

func (game *LobbyGames) UserReady(user *player.Player, respawn *coordinate.Coordinate) {
	if user.GetReady() == true {
		user.SetReady(false)
		game.DelRespawnUser(user)
	} else {
		user.SetReady(true)
		user.SetRespawn(respawn)
	}
}

func (game *LobbyGames)SetRespawnUser(user *player.Player, respawnID int) (*coordinate.Coordinate, error) {

	/*for _, respawn := range game.Respawns {
		if respawn.Id == respawnID && (respawn.UserName == "" || respawn.UserName == user.GetLogin()) {
			if respawn.UserName == user.GetLogin() {
				respawn.UserName = ""
				return nil, nil
			} else {
				respawn.UserName = user.GetLogin()
				return respawn, nil
			}
		}
	}*/ // todo адаптировать под новую логику

	return nil, errors.New("respawn busy")
}

func (game *LobbyGames) DelRespawnUser(user *player.Player) {
	/*for _, respawn := range game.Respawns {
		if respawn.UserName == user.GetLogin() {
			respawn.UserName = ""
		}
	}*/ // todo адаптировать под новую логику
}

func (game *LobbyGames) RemoveUser(user *player.Player)  {
	game.DelRespawnUser(user)
	for i, gameUser := range game.Users {
		if gameUser.GetLogin() == user.GetLogin() {
			game.Users = remove(game.Users, i)
		}
	}
}

func remove(users []*player.Player, i int) []*player.Player {
	users[len(users)-1], users[i] = users[i], users[len(users)-1]
	return users[:len(users)-1]
}
