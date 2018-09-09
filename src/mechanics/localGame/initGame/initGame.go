package initGame

import (
	"../../db/localGame/get"
	"../../localGame"
	"../../localGame/map/watchZone"
)

func InitGame(idGAme int) (newGame *localGame.Game) {

	newGame = get.Game(idGAme)

	players := get.Players(newGame)
	newGame.SetPlayers(players) // добавляем параметры всех игроков к обьекту игры

	Map := get.Map(newGame)
	units, unitStorage := get.AllUnits(newGame)

	newGame.SetMap(&Map)    // добавляем информацию об карте
	newGame.SetUnits(units) // добавляем имеющихся юнитов
	newGame.SetUnitsStorage(unitStorage)

	GetWatchPlayers(newGame)
	return
}

func GetWatchPlayers(game *localGame.Game) {
	for _, client := range game.GetPlayers() {
		watchZone.UpdateWatchZone(game, client)
	}
}
