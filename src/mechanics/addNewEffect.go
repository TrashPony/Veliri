package mechanics

import (
	"./unit"
	"./effect"
	"./db"
	"./coordinate"
)

const maxLvl = 2

func AddNewUnitEffect(gameUnit *unit.Unit, newEffect *effect.Effect) {

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
// TODO эти 2 метода очень похожи возможно их стоить обьеденить через интерфейс
func AddNewCoordinateEffect(gameCoordinate *coordinate.Coordinate, newEffect *effect.Effect) {

	addAnimate := true

	for i, coordinateEffect := range gameCoordinate.Effects {
		if coordinateEffect.Type != "unit_always_animate" && coordinateEffect.Type != "animate" {
			if coordinateEffect.Name == newEffect.Name {
				if coordinateEffect.Level+newEffect.Level >= maxLvl {
					newLvl := maxLvl - coordinateEffect.Level
					gameCoordinate.Effects[i] = db.GetNewLvlEffect(coordinateEffect, newLvl)
				} else {
					gameCoordinate.Effects[i] = db.GetNewLvlEffect(coordinateEffect, newEffect.Level)
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