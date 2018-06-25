package initGame

import (
	"../../localGame"
	"../../db/getLocalGame"
)


func InitGame(idGAme int) (newGame *localGame.Game) {

	newGame = getLocalGame.Game(idGAme)
	Map := getLocalGame.Map(newGame)
	units, unitStorage := getLocalGame.AllUnits(idGAme)
	matherShips := getLocalGame.MatherShips(idGAme)

	newGame.SetMap(&Map)       // добавляем информацию об карте
	newGame.SetUnits(units)    // добавляем имеющихся юнитов
	newGame.SetUnitsStorage(unitStorage)
	newGame.SetMatherShips(matherShips) // добавляем в игру все структуры на карте

	players := getLocalGame.Player(newGame)
	newGame.SetPlayers(players) // добавляем параметры всех игроков к обьекту игры

	return
}