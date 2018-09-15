package field

import (
	"../../mechanics/localGame/Phases/targetPhase"
	"github.com/gorilla/websocket"
)

func GetTargetZone(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	activeGame, findGame := Games.Get(client.GetGameID())

	gameUnit, findUnit := client.GetUnitStorage(msg.UnitID)
	if !findUnit {
		gameUnit, findUnit = client.GetHostileUnit(msg.Q, msg.R)
		if !findUnit {
			gameUnit, findUnit = client.GetUnit(msg.Q, msg.R)
		}
	}

	if findClient && findUnit && findGame {

		tmpUnit := *gameUnit

		tmpUnit.SetQ(msg.ToQ)
		tmpUnit.SetR(msg.ToR)

		ws.WriteJSON(TargetCoordinate{Event: "GetFirstTargets", Unit: gameUnit, Targets: targetPhase.GetWeaponTargetCoordinate(&tmpUnit, activeGame)})
	}
}
