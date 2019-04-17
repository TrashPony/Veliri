package get

import (
	"encoding/json"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"log"
)

func AllUnits(game *localGame.Game) (map[int]map[int]*unit.Unit, []*unit.Unit) {
	var units = make(map[int]map[int]*unit.Unit)
	var unitStorage = make([]*unit.Unit, 0)

	for _, gamePlayer := range game.GetPlayers() {

		if gamePlayer.Leave {
			getLeaveUnit(game, gamePlayer, &units)
		} else {

			gamePlayer.GetSquad().MatherShip.Owner = gamePlayer.GetLogin()

			UnitEffects(gamePlayer.GetSquad().MatherShip)          // берем эфекты ms
			addUnitToMap(&units, gamePlayer.GetSquad().MatherShip) // и кладем на карту, ms на карте с начала игры
			gamePlayer.GetSquad().MatherShip.CalculateParams()     // пересчитываем статы со всем эффектами

			// удаяем юнитов если они там есть
			gamePlayer.RemoveUnitsStorage()

			for _, playerUnit := range gamePlayer.GetSquad().MatherShip.Units {
				if playerUnit.Unit != nil {
					if playerUnit.Unit.OnMap {

						playerUnit.Unit.Owner = gamePlayer.GetLogin()
						UnitEffects(playerUnit.Unit)          // берем эфекты юнита
						addUnitToMap(&units, playerUnit.Unit) // и кладем на карту

					} else {
						playerUnit.Unit.Owner = gamePlayer.GetLogin()
						unitStorage = append(unitStorage, playerUnit.Unit)
						gamePlayer.AddUnitStorage(playerUnit.Unit)
					}

					playerUnit.Unit.CalculateParams() // пересчитываем статы со всем эффектами
				}
			}
		}
	}
	return units, unitStorage
}

func getLeaveUnit(game *localGame.Game, gamePlayer *player.Player, units *map[int]map[int]*unit.Unit) {
	rows, err := dbConnect.GetDBConnect().Query(
		"SELECT unit, id_user FROM game_leave_unit WHERE id_user = $1 AND id_game = $2", gamePlayer.GetID(), game.Id)
	if err != nil {
		log.Fatal("get game_leave_unit", err)
	}
	defer rows.Close()

	for rows.Next() {
		var jsonUnit []byte
		var memoryUnit unit.Unit
		var ownerID int

		err := rows.Scan(&jsonUnit, &ownerID)
		if err != nil {
			log.Fatal("scan game_leave_unit", err)
		}
		json.Unmarshal(jsonUnit, &memoryUnit)

		memoryUnit.Leave = true
		memoryUnit.OwnerID = ownerID
		addUnitToMap(units, &memoryUnit) // и кладем на карту
	}
}

func addUnitToMap(units *map[int]map[int]*unit.Unit, gameUnit *unit.Unit) {
	if (*units)[gameUnit.Q] != nil { // кладем юнита в матрицу
		(*units)[gameUnit.Q][gameUnit.R] = gameUnit
	} else {
		(*units)[gameUnit.Q] = make(map[int]*unit.Unit)
		(*units)[gameUnit.Q][gameUnit.R] = gameUnit
	}
}
