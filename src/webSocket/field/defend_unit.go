package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/targetPhase"
)

func DefendTarget(msg Message, client *player.Player) {

	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := games.Games.Get(client.GetGameID())

	if findUnit && findGame && !client.GetReady() {
		targetPhase.DefendTarget(gameUnit, client)
		SendMessage(Unit{Event: "UpdateUnit", Unit: gameUnit}, client.GetID(), activeGame.Id)
	}
}
