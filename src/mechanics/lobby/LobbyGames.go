package lobby

import (
	"errors"
	"../player"
	"../gameObjects/coordinate"
	LocalMap "../gameObjects/map"
	"../db/get"
)

type Game struct {
	ID       int
	Name     string
	Map      LocalMap.Map
	Creator  *player.Player
	Respawns []*coordinate.Coordinate
	Users    []*player.Player
}

func CreateNewLobbyGame(nameGame string, mapID int, creator *player.Player, id int) Game {

	respawns := get.Respawns(mapID)
	mp := get.Map(mapID)

	game := Game{ID: id, Name: nameGame, Map: mp, Creator: creator, Users: make([]*player.Player, 0), Respawns: respawns}
	game.Users = append(game.Users, creator)
	return game
}

func (game *Game) JoinToLobbyGame(user *player.Player) error {

	if len(game.Respawns) > len(game.Users) {
		game.Users = append(game.Users, user)
		return nil
	} else {
		return errors.New("lobby is full")
	}

	return errors.New("unknown error")
}

func (game *Game) UserReady(user *player.Player, respawn *coordinate.Coordinate) {
	if user.GetReady() == true {
		user.SetReady(false)
		game.DelRespawnUser(user)
	} else {
		user.SetReady(true)
		user.SetRespawn(respawn)
	}
}

func (game *Game) SetRespawnUser(user *player.Player, respawnID int) (*coordinate.Coordinate, error) {

	for _, user := range game.Users { // смотрим что бы респ не был занят
		if user.GetRespawn() == nil {
			continue
		} else {
			if user.GetRespawn().ID == respawnID {
				return nil, errors.New("respawn busy")
			}
		}
	}

	for _, respawn := range game.Respawns { // устанавливаем юзеру респаун
		if respawn.ID == respawnID {
			user.SetRespawn(respawn)
			return respawn, nil
		}
	}

	return nil, errors.New("respawn not find")
}

func (game *Game) DelRespawnUser(user *player.Player) {
	for _, respawn := range game.Respawns {
		if respawn.ID == user.GetRespawn().ID {
			user.SetRespawn(nil)
		}
	}
}

func (game *Game) RemoveUser(user *player.Player) {
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
