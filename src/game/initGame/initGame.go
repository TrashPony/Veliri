package initGame

import (
	"strconv"
)

func InitGame(idGame string, idUser int)([]string, int) {
	game := GetGame(idGame)
	userStat := GetUserStat(idGame, idUser)

	var idMap = game.idMap
	var playerParams = make([]string,0)

	if game.winner == "" {
		playerParams = []string{strconv.Itoa(userStat.price), strconv.Itoa(game.step), game.phase, userStat.ready}
	}

	return playerParams, idMap
}