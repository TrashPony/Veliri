package useEquip

import (
	"../../gameObjects/unit"
	"../../gameObjects/effect"
	"../../db/localGame/get"
	"../../gameObjects/coordinate"
)

func AddNewUnitEffect(gameUnit *unit.Unit, newEffect *effect.Effect) {

	addAnimate := true

	maxLvl := get.MaxLvlEffect(newEffect)

	for i, unitEffect := range gameUnit.Effects {
		if unitEffect.Type != "unit_always_animate" && unitEffect.Type != "animate" && unitEffect.Type != "zone_always_animate" {
			if unitEffect.Name == newEffect.Name {
				if unitEffect.Level+newEffect.Level >= maxLvl {
					newLvl := maxLvl - unitEffect.Level
					gameUnit.Effects[i] = get.NewLvlEffect(unitEffect, newLvl)
				} else {
					gameUnit.Effects[i] = get.NewLvlEffect(unitEffect, newEffect.Level)
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

func AddNewCoordinateEffect(gameCoordinate *coordinate.Coordinate, newEffect *effect.Effect) {

	addAnimate := true

	if newEffect.Type == "anchor" {
		return
	}

	maxLvl := get.MaxLvlEffect(newEffect)

	for i, coordinateEffect := range gameCoordinate.Effects {
		if coordinateEffect.Type != "unit_always_animate" && coordinateEffect.Type != "animate" && coordinateEffect.Type != "zone_always_animate" {
			if coordinateEffect.Name == newEffect.Name {
				if coordinateEffect.Level+newEffect.Level >= maxLvl {
					newLvl := maxLvl - coordinateEffect.Level
					gameCoordinate.Effects[i] = get.NewLvlEffect(coordinateEffect, newLvl)
				} else {
					gameCoordinate.Effects[i] = get.NewLvlEffect(coordinateEffect, newEffect.Level)
				}
				return
			}
		} else {
			if coordinateEffect.Name == newEffect.Name {
				addAnimate = false
			}
		}
	}

	if newEffect.Type == "unit_always_animate" || newEffect.Type == "animate" {
		if addAnimate {
			gameCoordinate.Effects = append(gameCoordinate.Effects, newEffect)
		}
	} else {
		gameCoordinate.Effects = append(gameCoordinate.Effects, newEffect)
	}
}