package lobby

import "errors"

type LobbyGames struct {
	Name     string
	Map      Map
	Creator  *User
	Respawns []*Respawn
	Users    []*User
}

func CreateNewLobbyGame(nameGame string, mapID int, creator *User) LobbyGames {

	respawns := GetRespawns(mapID)
	mp := GetMap(mapID)

	game := LobbyGames{Name: nameGame, Map: mp, Creator: creator, Users: make([]*User, 0), Respawns: respawns}
	game.Users = append(game.Users, creator)
	return game
}

func (game *LobbyGames) JoinToLobbyGame(user *User) error {

	if len(game.Respawns) > len(game.Users) {
		game.Users = append(game.Users, user)
		return nil
	} else {
		return errors.New("lobby is full")
	}

	return errors.New("unknown error")
}

func (game *LobbyGames) UserReady(user *User, respawn *Respawn) {
	if user.Ready == true {
		user.Ready = false
		game.DelRespawnUser(user)
	} else {
		user.Ready = true
		user.Respawn = respawn
	}
}

func (game *LobbyGames)SetRespawnUser(user *User, respawnID int) (*Respawn, error) {

	for _, respawn := range game.Respawns {
		if respawn.Id == respawnID && (respawn.UserName == "" || respawn.UserName == user.Name) {
			if respawn.UserName == user.Name {
				respawn.UserName = ""
				return nil, nil
			} else {
				respawn.UserName = user.Name
				return respawn, nil
			}
		}
	}

	return nil, errors.New("respawn busy")
}

func (game *LobbyGames) DelRespawnUser(user *User) {
	for _, respawn := range game.Respawns {
		if respawn.UserName == user.Name {
			respawn.UserName = ""
		}
	}
}

func (game *LobbyGames) RemoveUser(user *User)  {
	game.DelRespawnUser(user)
	for i, gameUser := range game.Users {
		if gameUser.Name == user.Name {
			game.Users = remove(game.Users, i)
		}
	}
}

func remove(users []*User, i int) []*User {
	users[len(users)-1], users[i] = users[i], users[len(users)-1]
	return users[:len(users)-1]
}
