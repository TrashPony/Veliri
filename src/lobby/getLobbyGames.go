package lobby

import (
	"errors"
	"strconv"
)

var openGames = make(map[string]LobbyGames)

func CreateNewLobbyGame(nameGame string, mapID int, nameCreator string) {

	respawns := GetRespawns(mapID)
	mp := GetMap(mapID)
	respswnsUser := make(map[Respawn]string)

	for i := 0; i < len(respawns); i++ {
		respswnsUser[respawns[i]] = ""
	}
	
	openGames[nameGame] = LobbyGames{Name: nameGame, Map: mp, Creator: nameCreator, Users: make(map[string]bool), Respawns: respswnsUser}
	openGames[nameGame].Users[nameCreator] = false
}

func JoinToLobbyGame(gameName string, userName string) error {
	for game := range openGames {
		if game == gameName {
			if len(openGames[game].Respawns) > len(openGames[game].Users) {
				openGames[game].Users[userName] = false
				return nil
			} else {
				return errors.New("lobby is full")
			}
		}
	}
	return errors.New("unknown error")
}

func UserReady(gameName string, userName string) {
	for game := range openGames {
		if game == gameName {
			if openGames[game].Users[userName] == true {
				openGames[game].Users[userName] = false
				DelRespawnUser(gameName, userName)
			} else {
				openGames[game].Users[userName] = true
			}
		}
	}
}

func GetLobbyGames() map[string]LobbyGames {
	return openGames
}

func GetGame(nameGame string) (LobbyGames, error) {
	var getGame LobbyGames
	for game := range openGames {
		if openGames[game].Name == nameGame {
			getGame = openGames[game]
			return getGame, nil
		}
	}
	return getGame, errors.New("no found this game")
}

func SetRespawnUser(gameName string, userName string, respawnId string) (string, error) {
	for game := range openGames {
		if game == gameName {
			for respawn := range openGames[game].Respawns {
				if strconv.Itoa(respawn.Id) == respawnId && (openGames[game].Respawns[respawn] == "" || openGames[game].Respawns[respawn] == userName) {
					if openGames[game].Respawns[respawn] == userName {
						openGames[game].Respawns[respawn] = ""
						return "", nil
					} else {
						openGames[game].Respawns[respawn] = userName
						return respawn.Name, nil
					}
				}
			}
		}
	}
	return "", errors.New("respawn busy")
}

func DisconnectLobbyGame(userName string) (game *LobbyGames) {
	for gameName := range openGames {
		for client, ready := range openGames[gameName].Users {
			if ready {
				for respawns := range openGames[gameName].Respawns {
					if openGames[gameName].Respawns[respawns] == userName {
						openGames[gameName].Respawns[respawns] = ""
					}
				}
			}

			if userName == client {
				tmpGame, find := openGames[gameName]
				delete(openGames[gameName].Users, client)
				if find {
					game = &tmpGame
					return
				}
			}
		}
	}
	return nil
}

func DelRespawnUser(gameName string, userName string) {
	for game := range openGames {
		if game == gameName {
			for respawn := range openGames[game].Respawns {
				if openGames[game].Respawns[respawn] == userName {
					openGames[game].Respawns[respawn] = ""
				}
			}
		}
	}
}

func DelLobbyGame(nameCreator string) (game *LobbyGames) {
	for gameName := range openGames {
		if openGames[gameName].Creator == nameCreator {
			tmpGame, find := openGames[gameName]
			if find {
				game = &tmpGame
				delete(game.Users, nameCreator)
				delete(openGames, gameName)
				return
			}
		}
	}
	return
}
