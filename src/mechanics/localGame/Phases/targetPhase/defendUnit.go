package targetPhase

import (
	"../../../gameObjects/unit"
	"../../../gameObjects/effect"
	"../../useEquip"
	"../../../player"
	"../../../db/updateSquad"
	"../../../db/localGame/update"
)

func DefendTarget(gameUnit *unit.Unit, client *player.Player) {
	gameUnit.Target = nil
	gameUnit.Action = true

	defendEffect := effect.Effect{TypeID: 21, Name: "defend", Level: 1, Type: "enhances", StepsTime: 1, Parameter: "armor", Quantity: 10, Percentages: false, Forever: false}
	useEquip.AddNewUnitEffect(gameUnit, &defendEffect)

	updateSquad.Squad(client.GetSquad())
	update.UnitEffects(gameUnit)
}
