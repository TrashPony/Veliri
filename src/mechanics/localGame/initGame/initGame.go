package initGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame/get"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/movePhase"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/watchZone"
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

	if newGame.Phase == "move" {
		checkMoveUnit(newGame)
	}

	return
}

func checkMoveUnit(game *localGame.Game) {
	move := false
	for _, q := range game.GetUnits() {
		for _, unit := range q {
			if unit.Move {
				move = true
			}
		}
	}

	if !move {
		movePhase.QueueMove(game)
	}
}

func GetWatchPlayers(game *localGame.Game) {
	for _, client := range game.GetPlayers() {
		watchZone.UpdateWatchZone(game, client)
	}
}
