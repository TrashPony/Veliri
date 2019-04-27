package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/movePhase"
)

func SelectUnit(msg Message, client *player.Player) {
	var findUnit bool
	var gameUnit *unit.Unit

	activeGame, findGame := games.Games.Get(client.GetGameID())

	if msg.Event == "SelectStorageUnit" {
		// т.к. юнит в корпусе берем координаты мс и присваиваем их юниту.
		gameUnit, findUnit = client.GetUnitStorage(msg.UnitID)
		if findUnit {
			gameUnit.Q = client.GetSquad().MatherShip.Q
			gameUnit.R = client.GetSquad().MatherShip.R
			SelectMove(client, gameUnit, activeGame, msg.Event)
		}
	} else {
		gameUnit, findUnit = client.GetUnit(msg.Q, msg.R)
	}

	if findUnit && findGame {
		if activeGame.Phase == "move" {

			// если юнит не в трюме добавляем в путь МС для возвращения в трюм
			if gameUnit.OnMap {
				msg.Event = "ToMC"
			}

			SelectMove(client, gameUnit, activeGame, msg.Event)
		}
	}
}

func SelectMove(client *player.Player, gameUnit *unit.Unit, actionGame *localGame.Game, event string) {
	if !client.GetReady() {
		if gameUnit.ActionPoints > 0 && gameUnit.Move {
			SendMessage(
				MoveCoordinate{
					Event: "SelectMoveUnit",
					Unit:  gameUnit,
					Move:  movePhase.GetMoveCoordinate(gameUnit, client, actionGame, event),
				},
				client.GetID(),
				actionGame.Id,
			)

		} else {
			SendMessage(ErrorMessage{Event: "Error", Error: "unit already move"}, client.GetID(), actionGame.Id)
		}
	} else {
		SendMessage(ErrorMessage{Event: "Error", Error: "you ready"}, client.GetID(), actionGame.Id)
	}
}

type MoveCoordinate struct {
	Event string                                       `json:"event"`
	Unit  *unit.Unit                                   `json:"unit"`
	Move  map[string]map[string]*coordinate.Coordinate `json:"move"`
}
