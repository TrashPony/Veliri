package field

import (
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/localGame"
	"../../mechanics/localGame/Phases/targetPhase"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
)

func SelectWeapon(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := Games.Get(client.GetGameID())

	if findClient && findUnit && findGame {
		if activeGame.Phase == "targeting" {
			SelectTarget(client, gameUnit, activeGame, ws)
		}
	}
}

func SelectTarget(client *player.Player, gameUnit *unit.Unit, actionGame *localGame.Game, ws *websocket.Conn) {
	if !client.GetReady() && !gameUnit.Defend {
		ws.WriteJSON(TargetCoordinate{Event: "GetTargets", Unit: gameUnit, Targets: targetPhase.GetWeaponTargetCoordinate(gameUnit, actionGame)})
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "you ready"})
	}
}

type TargetCoordinate struct {
	Event   string                                       `json:"event"`
	Unit    *unit.Unit                                   `json:"unit"`
	Targets map[string]map[string]*coordinate.Coordinate `json:"targets"`
}
