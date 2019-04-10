package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/movePhase"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func GetPreviewPath(msg Message, client *player.Player) {
	var findStart, findEnd bool
	var startCoordinate, endCoordinate *coordinate.Coordinate
	var event string

	activeGame, findGame := games.Games.Get(client.GetGameID())

	gameUnit, findUnit := client.GetUnitStorage(msg.UnitID)
	if findUnit {
		// т.к. юнит в корпусе берем координаты мс и присваиваем их юниту.
		startCoordinate, findStart = activeGame.GetMap().GetCoordinate(client.GetSquad().MatherShip.Q, client.GetSquad().MatherShip.R)
		endCoordinate, findEnd = activeGame.GetMap().GetCoordinate(msg.ToQ, msg.ToR)
		gameUnit.Q = startCoordinate.Q
		gameUnit.R = startCoordinate.R
		event = "SelectStorageUnit"
	} else {
		gameUnit, findUnit = client.GetUnit(msg.Q, msg.R)
		if findUnit {
			startCoordinate, findStart = activeGame.GetMap().GetCoordinate(gameUnit.Q, gameUnit.R)
			endCoordinate, findEnd = activeGame.GetMap().GetCoordinate(msg.ToQ, msg.ToR)
		}
	}

	if findUnit && findGame && findStart && findEnd {
		err, path := movePhase.FindPath(client, activeGame.GetMap(), startCoordinate, endCoordinate, gameUnit, event)
		if err != nil {
			SendMessage(ErrorMessage{Event: "Error", Error: err.Error()}, client.GetID(), activeGame.Id)
		} else {
			SendMessage(previewPath{Event: "PreviewPath", Path: path}, client.GetID(), activeGame.Id)
		}
	}
}

type previewPath struct {
	Event string                   `json:"event"`
	Path  []*coordinate.Coordinate `json:"path"`
}
