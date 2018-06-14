package mechanics

import (
	"./unit"
	"./effect"
	"./db"
)

const maxLvl = 2

func AddNewEffect(gameUnit *unit.Unit, newEffect *effect.Effect) {

	addAnimate := true

	for i, unitEffect := range gameUnit.Effects {
		if unitEffect.Type != "unit_always_animate" && unitEffect.Type != "animate" {
			if unitEffect.Name == newEffect.Name {
				if unitEffect.Level+newEffect.Level >= maxLvl {
					newLvl := maxLvl - unitEffect.Level
					gameUnit.Effects[i] = db.GetNewLvlEffect(unitEffect, newLvl)
				} else {
					gameUnit.Effects[i] = db.GetNewLvlEffect(unitEffect, newEffect.Level)
				}
				return
			}
		} else {
			if unitEffect.Name == newEffect.Name {
				addAnimate = false
			}
		}
	}

	if newEffect.Type == "unit_always_animate" || newEffect.Type == "animate" {
		if addAnimate {
			gameUnit.Effects = append(gameUnit.Effects, newEffect)
		}
	} else {
		gameUnit.Effects = append(gameUnit.Effects, newEffect)
	}
}
