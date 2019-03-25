package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/targetPhase"
	"github.com/gorilla/websocket"
	"strconv"
)

func SetTarget(msg Message, ws *websocket.Conn) {

	client := localGame.Clients.GetByWs(ws)
	if client != nil {

		gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
		activeGame, findGame := games.Games.Get(client.GetGameID())

		if findUnit && findGame && !client.GetReady() && !gameUnit.Defend && gameUnit.GetWeaponSlot() != nil && gameUnit.GetAmmoCount() > 0 {

			targetCoordinate := targetPhase.GetWeaponTargetCoordinate(gameUnit, activeGame, client, "GetTargets")
			_, find := targetCoordinate[strconv.Itoa(msg.ToQ)][strconv.Itoa(msg.ToR)]

			if find {
				targetPhase.SetTarget(gameUnit, activeGame, msg.ToQ, msg.ToR, client)
				SendMessage(Unit{Event: "UpdateUnit", Unit: gameUnit}, client.GetID(), activeGame.Id)
			} else {
				SendMessage(ErrorMessage{Event: "Error", Error: "not allow"}, client.GetID(), activeGame.Id)
			}
		}
	}
}

type Unit struct {
	Event    string     `json:"event"`
	UserName string     `json:"user_name"`
	GameID   int        `json:"game_id"`
	Unit     *unit.Unit `json:"unit"`
}
