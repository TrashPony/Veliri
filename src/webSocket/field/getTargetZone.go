package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/Phases/targetPhase"
)

func GetTargetZone(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.X, msg.Y)
	activeGame, findGame := Games.Get(client.GetGameID())

	if !findUnit {
		gameUnit, findUnit = client.GetHostileUnit(msg.X, msg.Y)
	}

	if findClient && findUnit && findGame {

		tmpUnit := *gameUnit

		tmpUnit.SetX(msg.ToX)
		tmpUnit.SetY(msg.ToY)

		ws.WriteJSON(TargetCoordinate{Event: "GetFirstTargets", Unit: gameUnit, Targets: targetPhase.GetTargetCoordinate(&tmpUnit, client, activeGame)})
	}
}
