package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/targetPhase"
	"github.com/gorilla/websocket"
)

func DefendTarget(msg Message, ws *websocket.Conn) {
	client := localGame.Clients.GetByWs(ws)

	if client != nil {

		gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
		activeGame, findGame := games.Games.Get(client.GetGameID())

		if findUnit && findGame && !client.GetReady() {
			targetPhase.DefendTarget(gameUnit, client)
			SendMessage(Unit{Event: "UpdateUnit", Unit: gameUnit}, client.GetID(), activeGame.Id)
		}
	}
}
