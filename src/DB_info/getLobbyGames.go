package DB_info

var openGames = make(map[string]LobbyGames)

func CreateNewLobbyGame (nameGame string, nameMap string, nameCreator string ) {

	openGames[nameGame] = LobbyGames{Name:nameGame, Map:nameMap, Creator:nameCreator, Users:make(map[string]bool)}
	openGames[nameGame].Users[nameCreator] = true
}

func JoinToLobbyGame(gameName string, userName string ) {
	for game := range openGames {
		if game == gameName {
			openGames[game].Users[userName] = false
		}
	}
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

func DelLobbyGame(nameCreator string)  {
	for game := range openGames {
		if openGames[game].Creator == nameCreator{
			delete(openGames,game)
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
