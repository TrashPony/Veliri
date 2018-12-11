package field

import (
	"../../mechanics/factories/games"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/localGame/Phases/targetPhase"
	"github.com/gorilla/websocket"
	"strconv"
)

func SetTarget(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := games.Games.Get(client.GetGameID())

	if findClient && findUnit && findGame && !client.GetReady() && !gameUnit.Defend && gameUnit.GetWeaponSlot() != nil && gameUnit.GetAmmoCount() > 0 {

		targetCoordinate := targetPhase.GetWeaponTargetCoordinate(gameUnit, activeGame, client, "GetTargets")
		_, find := targetCoordinate[strconv.Itoa(msg.ToQ)][strconv.Itoa(msg.ToR)]

		if find {
			targetPhase.SetTarget(gameUnit, activeGame, msg.ToQ, msg.ToR, client)
			ws.WriteJSON(Unit{Event: "UpdateUnit", Unit: gameUnit})
		} else {
			ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not allow"})
		}
	}
}

type Unit struct {
	Event    string     `json:"event"`
	UserName string     `json:"user_name"`
	GameID   int        `json:"game_id"`
	Unit     *unit.Unit `json:"unit"`
}
