package DB_info

func ConnectGame(nameGame string, userName string) (string, bool)  {
	for game := range openGames {
		if game.nameGame == nameGame{
			game.nameNewPlayer = userName
			StartNewGame(game.nameMap, game)
			DelLobbyGame(game.nameCreator)
			return game.nameCreator, true
		}
	}
	return "", false
}

