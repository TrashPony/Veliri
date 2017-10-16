package DB_info

var openGames = make(map[string]LobbyGames)

func CreateNewLobbyGame (nameGame string, nameMap string, nameCreator string ) {
	respawns := GetRespawns(nameMap)
	openGames[nameGame] = LobbyGames{Name:nameGame, Map:nameMap, Creator:nameCreator, Users:make(map[string]bool), Respawns:respawns}
	openGames[nameGame].Users[nameCreator] = true
}

func JoinToLobbyGame(gameName string, userName string ) (bool, string) {
	for game := range openGames {
		if game == gameName {
			if len(openGames[game].Respawns) > len(openGames[game].Users) {
				openGames[game].Users[userName] = false
				return true, ""
			} else {
				return false, "lobby is full"
			}
		}
	}
	return false, "unknown error"
}

func UserReady(gameName string, userName string)  {
	for game := range openGames {
		if game == gameName {
			if openGames[game].Users[userName] == true {
				openGames[game].Users[userName] = false
			} else {
				openGames[game].Users[userName] = true
			}
		}
	}
}

func GetLobbyGames()(map[string]LobbyGames) {
	return openGames
}

func GetUserList(nameGame string)(map[string]bool)  {
	for game := range openGames {
		if openGames[game].Name == nameGame{
			return openGames[game].Users
		}
	}
	return nil
}

func DisconnectLobbyGame(userName string)(bool, string) {
	var success bool = false
	var nameGame string
	for game := range openGames {
		for client := range openGames[game].Users {
			if userName == client {
				nameGame = openGames[game].Name
				delete(openGames[game].Users, client)
				success = true
			}
		}
	}
	return success, nameGame
}


func DelLobbyGame(nameCreator string) (bool, map[string]bool)  {
	var success bool = false
	users := make(map[string]bool)
	for game := range openGames {
		if openGames[game].Creator == nameCreator{
			users = openGames[game].Users
			delete(openGames,game)
			success = true
		}
	}

	return success, users
}