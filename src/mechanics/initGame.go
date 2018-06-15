package mechanics

import (
	"./game"
	"./db"
)


func InitGame(idGAme int) (newGame *game.Game) {

	newGame = db.GetGame(idGAme)
	Map := db.GetMap(newGame)
	units, unitStorage := db.GetAllUnits(idGAme)
	matherShips := db.GetMatherShips(idGAme)

	newGame.SetMap(&Map)       // добавляем информацию об карте
	newGame.SetUnits(units)    // добавляем имеющихся юнитов
	newGame.SetUnitsStorage(unitStorage)
	newGame.SetMatherShips(matherShips) // добавляем в игру все структуры на карте

	players := db.GetPlayer(newGame)
	newGame.SetPlayers(players) // добавляем параметры всех игроков к обьекту игры

	return
}