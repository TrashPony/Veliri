package attackPhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/detail"
	"../../../gameObjects/unit"
	"../../../localGame"
	"math/rand"
	"time"
)

func InitAttack(attacking *unit.Unit, target *coordinate.Coordinate, game *localGame.Game) *ResultBattle {

	if attacking.GetWeaponSlot() != nil && attacking.GetAmmoCount() > 0 {
		attacking.GetWeaponSlot().AmmoQuantity -= 1
		return MapAttack(attacking, target, game, attacking.GetWeaponSlot())
	} else {
		return &ResultBattle{Error: "no ammo"}
	}

	return nil
}

func MapAttack(attacking *unit.Unit, target *coordinate.Coordinate, game *localGame.Game, weapon *detail.BodyWeaponSlot) *ResultBattle {
	attackZone := coordinate.GetCoordinatesRadius(target, weapon.Ammo.AreaCovers)
	targetsUnit := make([]TargetUnit, 0)

	for _, attackCoordinate := range attackZone {
		targetUnit, find := game.GetUnit(attackCoordinate.Q, attackCoordinate.R)
		if find {

			var damage int

			if targetUnit.Q == target.Q && targetUnit.R == target.R {
				damage = calculateDamage(targetUnit, weapon.Ammo.MaxDamage, weapon.Ammo.MinDamage)
			} else {
				//т.к. не эпичентер атаки юниты получают только 50% урона
				damage = calculateDamage(targetUnit, weapon.Ammo.MaxDamage/2, weapon.Ammo.MinDamage/2)
			}

			targetUnit.HP -= damage

			broken := breakingEquip(targetUnit, damage)

			targetsUnit = append(targetsUnit, TargetUnit{Unit: *targetUnit, Damage: damage, BreakingEquip: broken})
		}
	}

	return &ResultBattle{AttackUnit: *attacking, TargetUnits: targetsUnit}
}

func calculateDamage(targetUnit *unit.Unit, maxDamage, minDamage int) int {
	//Генератор случайных чисел обычно нужно рандомизировать перед использованием, иначе, он, действительно,
	// будет выдавать одну и ту же последовательность.
	rand.Seed(time.Now().UnixNano())
	damage := rand.Intn(maxDamage-minDamage) + minDamage

	armor := targetUnit.Armor

	for _, effect := range targetUnit.Effects {
		if effect != nil && effect.Parameter == "armor" {
			if effect.Type == "enhances" {
				if effect.Percentages {
					armor += armor / 100 * effect.Quantity
				} else {
					armor += effect.Quantity
				}
			}
			if effect.Type == "reduced" {
				if effect.Percentages {
					armor -= armor / 100 * effect.Quantity
				} else {
					armor -= effect.Quantity
				}
			}
		}
	}

	damage -= armor

	if damage < 0 {
		damage = 0
	}

	return damage
}

func breakingEquip(targetUnit *unit.Unit, damage int) bool {

	return breaking(targetUnit.Body.EquippingI, damage) ||
		breaking(targetUnit.Body.EquippingII, damage) ||
		breaking(targetUnit.Body.EquippingIII, damage) ||
		breaking(targetUnit.Body.EquippingIV, damage) ||
		breaking(targetUnit.Body.EquippingV, damage)
}

func breaking(equip map[int]*detail.BodyEquipSlot, damage int) bool {
	for _, equipSlot := range equip {
		if equipSlot.Equip != nil {

			// дамаг в 20%, в итоге должен зависеть от скила игрока
			if equipSlot.HP-damage/5 > 0 {
				equipSlot.HP -= damage / 5
				return false
			} else {
				equipSlot.HP = 0
				return true
			}
		}
	}

	return false
}
