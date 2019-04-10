package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases/targetPhase"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func SelectWeapon(msg Message, client *player.Player) {

	gameUnit, findUnit := client.GetUnit(msg.Q, msg.R)
	activeGame, findGame := games.Games.Get(client.GetGameID())

	if findUnit && findGame && gameUnit.GetWeaponSlot() != nil && gameUnit.GetAmmoCount() > 0 {
		if activeGame.Phase == "targeting" {
			SelectTarget(client, gameUnit, activeGame)
		}
	}
}

func SelectTarget(client *player.Player, gameUnit *unit.Unit, actionGame *localGame.Game) {
	if !client.GetReady() && !gameUnit.Defend {
		SendMessage(
			TargetCoordinate{
				Event:   "GetTargets",
				Unit:    gameUnit,
				Targets: targetPhase.GetWeaponTargetCoordinate(gameUnit, actionGame, client, "GetTargets"),
			},
			client.GetID(),
			actionGame.Id,
		)
	} else {
		SendMessage(ErrorMessage{Event: "Error", Error: "you ready"}, client.GetID(), actionGame.Id)
	}
}

type TargetCoordinate struct {
	Event   string                                       `json:"event"`
	Unit    *unit.Unit                                   `json:"unit"`
	Targets map[string]map[string]*coordinate.Coordinate `json:"targets"`
}
