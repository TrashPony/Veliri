package initGame

import (
	"../../localGame"
	"../../db/localGame/get"
)


func InitGame(idGAme int) (newGame *localGame.Game) {

	newGame = get.Game(idGAme)

	players := get.Players(newGame)
	newGame.SetPlayers(players) // добавляем параметры всех игроков к обьекту игры

	Map := get.Map(newGame)
	//units, unitStorage := get.AllUnits(idGAme)
	//matherShips := get.MatherShips(idGAme)

	newGame.SetMap(&Map)       // добавляем информацию об карте
	//newGame.SetUnits(units)    // добавляем имеющихся юнитов
	//newGame.SetUnitsStorage(unitStorage)
	//newGame.SetMatherShips(matherShips) // добавляем в игру все структуры на карте

	return
}