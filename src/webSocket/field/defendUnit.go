package field

import (
	"../../mechanics/factories/games"
	"../../mechanics/localGame/Phases/targetPhase"
	"github.com/gorilla/websocket"
)

func DefendTarget(msg Message, ws *websocket.Conn) {
	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	_, findGame := games.Games.Get(client.GetGameID())

	if findClient && findUnit && findGame && !client.GetReady() {
		targetPhase.DefendTarget(gameUnit, client)
		ws.WriteJSON(Unit{Event: "UpdateUnit", Unit: gameUnit})
	}
}
