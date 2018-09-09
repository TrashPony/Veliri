package field

import (
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/localGame/Phases/movePhase"
	"github.com/gorilla/websocket"
)

func GetPreviewPath(msg Message, ws *websocket.Conn) {
	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := Games.Get(client.GetGameID())

	startCoordinate, findStart := activeGame.GetMap().GetCoordinate(gameUnit.Q, gameUnit.R)
	endCoordinate, findEnd := activeGame.GetMap().GetCoordinate(msg.ToQ, msg.ToR)

	if findClient && findUnit && findGame && findStart && findEnd {
		err, path := movePhase.FindPath(client, activeGame.GetMap(), startCoordinate, endCoordinate, gameUnit)
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
