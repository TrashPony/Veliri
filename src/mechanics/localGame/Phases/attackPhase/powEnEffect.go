package attackPhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/effect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func powEnEffect(effect *effect.Effect, toUseUnit *unit.Unit, targetsUnit *[]TargetUnit) {
	if effect.Type == "replenishes" && effect.Parameter == "hp" {

		if toUseUnit.HP+effect.Quantity <= toUseUnit.Body.MaxHP {
			toUseUnit.HP += effect.Quantity
		} else {
			toUseUnit.HP = toUseUnit.Body.MaxHP
		}

		*targetsUnit = append(*targetsUnit, TargetUnit{Unit: *toUseUnit, Heal: effect.Quantity})
	}

	if effect.Type == "takes_away" && effect.Parameter == "hp" {
		toUseUnit.HP -= effect.Quantity
		*targetsUnit = append(*targetsUnit, TargetUnit{Unit: *toUseUnit, Damage: effect.Quantity})
	}

	if effect.Type == "replenishes" && effect.Parameter == "power" {

		if toUseUnit.Power+effect.Quantity <= toUseUnit.Body.MaxPower {
			toUseUnit.Power += effect.Quantity
		} else {
			toUseUnit.Power = toUseUnit.Body.MaxPower
		}

		*targetsUnit = append(*targetsUnit, TargetUnit{Unit: *toUseUnit, Power: effect.Quantity})
	}

	if effect.Type == "takes_away" && effect.Parameter == "power" {
		toUseUnit.Power -= effect.Quantity

		if toUseUnit.Power < 0 {
			toUseUnit.Power = 0
		}

		*targetsUnit = append(*targetsUnit, TargetUnit{Unit: *toUseUnit, Power: -effect.Quantity})
	}
}
