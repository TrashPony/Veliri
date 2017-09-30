package initGame

import "strconv"

func InitGame(idGame string)([]string, int) {
	game := GetGame(idGame)
	var idMap = game.idMap
	var playerParams = make([]string,0)

	if !game.gameend {
		playerParams = []string{strconv.Itoa(game.price1), strconv.Itoa(game.step), game.phase}
	}

	return playerParams, idMap
}


