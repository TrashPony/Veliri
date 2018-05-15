package mechanics

import (
	"./game"
	"./dbo"
)


func InitGame(idGAme int) (newGame *game.Game) {

	newGame = dbo.GetGame(idGAme)
	Map := dbo.GetMap(newGame.MapID)
	units, unitStorage := dbo.GetAllUnits(idGAme)
	matherShips := dbo.GetMatherShips(idGAme)

	newGame.SetMap(&Map)       // добавляем информацию об карте
	newGame.SetUnits(units)    // добавляем имеющихся юнитов
	newGame.SetUnitsStorage(unitStorage)
	newGame.SetMatherShips(matherShips) // добавляем в игру все структуры на карте

	players := dbo.GetPlayer(newGame)
	newGame.SetPlayers(players) // добавляем параметры всех игроков к обьекту игры

	return
}