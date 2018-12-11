package field

import (
	"../../mechanics/factories/games"
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/localGame/Phases/movePhase"
	"github.com/gorilla/websocket"
)

func GetPreviewPath(msg Message, ws *websocket.Conn) {
	var findStart, findEnd bool
	var startCoordinate, endCoordinate *coordinate.Coordinate
	var event string

	client, findClient := usersFieldWs[ws]
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

	if findClient && findUnit && findGame && findStart && findEnd {
		err, path := movePhase.FindPath(client, activeGame.GetMap(), startCoordinate, endCoordinate, gameUnit, event)
		if err != nil {
			ws.WriteJSON(ErrorMessage{Event: "Error", Error: err.Error()})
		} else {
			ws.WriteJSON(previewPath{Event: "PreviewPath", Path: path})
		}
	}
}

type previewPath struct {
	Event string                   `json:"event"`
	Path  []*coordinate.Coordinate `json:"path"`
}
