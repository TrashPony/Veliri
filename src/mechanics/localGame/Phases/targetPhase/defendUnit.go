package targetPhase

import (
	"../../../db/localGame/update"
	squadUpdate "../../../db/squad/update"
	"../../../gameObjects/effect"
	"../../../gameObjects/unit"
	"../../../player"
	"../../useEquip"
)

func DefendTarget(gameUnit *unit.Unit, client *player.Player) {
	gameUnit.Target = nil
	gameUnit.Defend = true

	defendEffect := effect.Effect{TypeID: 21, Name: "defend", Level: 1, Type: "enhances", StepsTime: 1, Parameter: "armor", Quantity: 10, Percentages: false, Forever: false}
	useEquip.AddNewUnitEffect(gameUnit, &defendEffect, 1)

	gameUnit.CalculateParams()
	squadUpdate.Squad(client.GetSquad())
	update.UnitEffects(gameUnit)
}
