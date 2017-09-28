package DB_info

var openGames = make(map[Games]bool)

func OpenLobbyGameList()([]string) {
	var responseNameGame = ""
	var responseNameMap = ""
	var responseNameUser = ""

	for game := range openGames {
		responseNameGame = responseNameGame + game.nameGame + ":" ;
		responseNameMap = responseNameMap + game.nameMap + ":";
		responseNameUser = responseNameUser + game.nameCreator + ":";
	}

	var games []string
	games = append(games, responseNameGame)
	games = append(games, responseNameMap)
	games = append(games, responseNameUser)

	return games
}

func CreateNewLobbyGame (nameGame string, nameMap string, nameCreator string ) {
	openGames[Games{nameGame, nameMap, nameCreator, ""}] = true
}

func DelLobbyGame(nameCreator string)  {
	for game := range openGames {
		if game.nameCreator == nameCreator{
			delete(openGames,game)
		}
	}
}