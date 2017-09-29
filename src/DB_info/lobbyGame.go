package DB_info

var openGames = make(map[string]Games)

func CreateNewLobbyGame (nameGame string, nameMap string, nameCreator string ) {
	openGames[nameGame] = Games{nameGame, nameMap, nameCreator, ""}
}

func JoinToLobbyGame(nameGame string, nameUser string ) (string)  {

	for game := range openGames {
		if game == nameGame && openGames[game].nameNewPlayer == ""{
			openGames[game] = Games{nameGame, openGames[game].nameMap, openGames[game].nameCreator, nameUser}
			return openGames[game].nameCreator
		}
	}
	return ""
}

func DelLobbyGame(nameCreator string)  {
	for game := range openGames {
		if openGames[game].nameCreator == nameCreator{
			delete(openGames,game)
		}
	}
}

func OpenLobbyGameList()([]string) {
	var responseNameGame = ""
	var responseNameMap = ""
	var responseNameUser = ""

	for game := range openGames {
		responseNameGame = responseNameGame + openGames[game].nameGame + ":" ;
		responseNameMap = responseNameMap + openGames[game].nameMap + ":";
		responseNameUser = responseNameUser + openGames[game].nameCreator + ":";
	}

	var games []string
	games = append(games, responseNameGame)
	games = append(games, responseNameMap)
	games = append(games, responseNameUser)

	return games
}

func GetUserList(nameGame string)([]string)  {
	playerList := make([]string, 0)
	for game := range openGames {
		if openGames[game].nameGame == nameGame{
			playerList = append(playerList, openGames[game].nameCreator)
			playerList = append(playerList, openGames[game].nameNewPlayer)
		}
	}
	return playerList
}
