package targetPhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	squadUpdate "github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/effect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/useEquip"
)

func DefendTarget(gameUnit *unit.Unit, client *player.Player) {
	gameUnit.Target = nil
	gameUnit.Defend = true

	defendEffect := effect.Effect{TypeID: 21, Name: "defend", Level: 1, Type: "enhances", StepsTime: 1, Parameter: "armor", Quantity: 10, Percentages: false, Forever: false}
	useEquip.AddNewUnitEffect(gameUnit, &defendEffect, 1)

	gameUnit.CalculateParams()
	squadUpdate.Squad(client.GetSquad(), true)
	update.UnitEffects(gameUnit)
}
