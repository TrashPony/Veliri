package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/targetPhase"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
)

func SelectWeapon(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := games.Games.Get(client.GetGameID())

	if findClient && findUnit && findGame && gameUnit.GetWeaponSlot() != nil && gameUnit.GetAmmoCount() > 0 {
		if activeGame.Phase == "targeting" {
			SelectTarget(client, gameUnit, activeGame, ws)
		}
	}
}

func SelectTarget(client *player.Player, gameUnit *unit.Unit, actionGame *localGame.Game, ws *websocket.Conn) {
	if !client.GetReady() && !gameUnit.Defend {
		ws.WriteJSON(TargetCoordinate{Event: "GetTargets", Unit: gameUnit, Targets: targetPhase.GetWeaponTargetCoordinate(gameUnit, actionGame, client, "GetTargets")})
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "you ready"})
	}
}

type TargetCoordinate struct {
	Event   string                                       `json:"event"`
	Unit    *unit.Unit                                   `json:"unit"`
	Targets map[string]map[string]*coordinate.Coordinate `json:"targets"`
}
