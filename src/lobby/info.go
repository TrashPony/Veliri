package lobby


var openGames = make(map[Games]bool)

func OpenGameList()([]string) {
	var responseNameGame = ""
	var responseNameMap = ""
	var responseNameUser = ""

	openGames[Games{"ws", "test", "login"}] = true // тест
	openGames[Games{"ws2", "test2", "login2"}] = true // тест

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

func MapList() {

}

func CreateNewGame () {

}

type Games struct{
	nameGame string
	nameMap string
	nameCreator string
}