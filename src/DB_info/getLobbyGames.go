package DB_info

var openGames = make(map[string]Games)

func CreateNewLobbyGame (nameGame string, nameMap string, nameCreator string ) {
	openGames[nameGame] = Games{nameGame, nameMap, nameCreator, ""}
}

func JoinToLobbyGame(nameGame string, nameUser string ) (string)  {

	for game := range openGames {
		if game == nameGame && openGames[game].NewPlayer == ""{
			openGames[game] = Games{nameGame, openGames[game].Map, openGames[game].Creator, nameUser}
			return openGames[game].Creator
		}
	}
	return ""
}

func DelLobbyGame(nameCreator string)  {
	for game := range openGames {
		if openGames[game].Creator == nameCreator{
			delete(openGames,game)
		}
	}
}

func GetLobbyGames()(map[string]Games) {
	return openGames
}

func GetUserList(nameGame string)([]string)  {
	playerList := make([]string, 0)
	for game := range openGames {
		if openGames[game].Name == nameGame{
			playerList = append(playerList, openGames[game].Creator)
			playerList = append(playerList, openGames[game].NewPlayer)
		}
	}
	return playerList
}
