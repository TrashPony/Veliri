package attackPhase

import (
	"../../../../mechanics/gameObjects/coordinate"
	"../../../../mechanics/gameObjects/unit"
	"../../../localGame"
	"math/rand"
)

// TODO влияние защиты, хащитный бафов

func InitAttack(attacking *unit.Unit, target *coordinate.Coordinate, game *localGame.Game) *ResultBattle {

	if attacking.Body.Weapons[0].AmmoQuantity > 0 {

		attacking.Body.Weapons[0].AmmoQuantity -= 1

		if attacking.Body.Weapons[0].Ammo.AreaCovers == 0 {
			return TargetAttack(attacking, target, game) // тоесть у пухи радиус атаки 0
		} else {
			return MapAttack(attacking, target, game) // тоесть у пухи радиус атаки больше 0ля и оружие бьет по площади
		}

	} else {
		return &ResultBattle{Error: "no ammo"}
	}
}

func TargetAttack(attacking *unit.Unit, target *coordinate.Coordinate, game *localGame.Game) *ResultBattle {

	targetUnit, findUnit := game.GetUnit(target.Q, target.R)

	if findUnit {

		maxDamage := attacking.Body.Weapons[0].Ammo.MaxDamage
		minDamage := attacking.Body.Weapons[0].Ammo.MinDamage
		targetUnit.HP -= rand.Intn(maxDamage-minDamage) + minDamage

		return &ResultBattle{Map: false, AttackUnit: *attacking, TargetUnit: *targetUnit}
	} else {
		return &ResultBattle{Map: false, AttackUnit: *attacking, TargetUnit: nil}
	}
}

func MapAttack(attacking *unit.Unit, target *coordinate.Coordinate, game *localGame.Game) *ResultBattle {
	attackZone := coordinate.GetCoordinatesRadius(target, attacking.Body.Weapons[0].Ammo.AreaCovers)
	targetsUnit := make([]unit.Unit, 0)

	for _, attackCoordinate := range attackZone {
		targetUnit, find := game.GetUnit(attackCoordinate.Q, attackCoordinate.R)
		if find {

			maxDamage := attacking.Body.Weapons[0].Ammo.MaxDamage
			minDamage := attacking.Body.Weapons[0].Ammo.MinDamage
			targetUnit.HP -= rand.Intn(maxDamage-minDamage) + minDamage

			targetsUnit = append(targetsUnit, *targetUnit)
		}
	}

	return &ResultBattle{Map: true, AttackUnit: *attacking, TargetsUnit: targetsUnit}
}
