package initGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame/get"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/movePhase"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/watchZone"
)

func InitGame(userID int) (newGame *localGame.Game) {

	newGame = get.Game(userID)

	newGame.SetPlayers(get.Players(newGame)) // добавляем параметры всех игроков к обьекту игры

	// берем копию карты что бы она не влияла на другие сессии, и накладываем на нее эффекты текущий игры ))
	gameMap, _ := maps.Maps.GetCopyByID(newGame.MapID)
	for _, qLine := range gameMap.OneLayerMap {
		for _, gameCoordinate := range qLine {
			gameCoordinate.GameID = newGame.Id
			get.CoordinateEffects(gameCoordinate)
		}
	}

	// берем юнитов из пользвателей и добавляем в игру
	units, unitStorage := get.AllUnits(newGame)

	newGame.SetMap(gameMap) // добавляем информацию об карте
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
