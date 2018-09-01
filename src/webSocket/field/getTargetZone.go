package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/localGame/Phases/targetPhase"
)

func GetTargetZone(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := Games.Get(client.GetGameID())

	if !findUnit {
		gameUnit, findUnit = client.GetHostileUnit(msg.Q, msg.R)
	}

	if findClient && findUnit && findGame {

		tmpUnit := *gameUnit

		tmpUnit.SetQ(msg.ToX)
		tmpUnit.SetR(msg.ToY)

		ws.WriteJSON(TargetCoordinate{Event: "GetFirstTargets", Unit: gameUnit, Targets: targetPhase.GetWeaponTargetCoordinate(&tmpUnit, activeGame)})
	}
}
