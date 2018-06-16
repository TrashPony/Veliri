package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/Phases/targetPhase"
)

func DefendTarget(msg Message, ws *websocket.Conn) {
	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.X, msg.Y)
	activeGame, findGame :=  Games.Get(client.GetGameID())

	if findClient && findUnit && findGame && !client.GetReady() {
		targetPhase.DefendTarget(gameUnit)
		ws.WriteJSON(Unit{Event: "UpdateUnit", Unit: gameUnit})
		updateUnitHostileUser(client, activeGame, gameUnit)
	}
}
