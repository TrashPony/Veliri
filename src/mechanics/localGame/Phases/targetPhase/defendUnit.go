package targetPhase

import (
	"../../../gameObjects/unit"
	"../../../db/localGame/update"
	"../../../gameObjects/effect"
	"../../useEquip"
)

func DefendTarget(gameUnit *unit.Unit) {
	gameUnit.Target = nil
	gameUnit.Action = true

	defendEffect := effect.Effect{TypeID: 21, Name: "defend", Level: 1, Type: "enhances", StepsTime: 1, Parameter: "armor", Quantity: 10, Percentages: false, Forever: false}
	useEquip.AddNewUnitEffect(gameUnit, &defendEffect)

	update.Unit(gameUnit)
}
