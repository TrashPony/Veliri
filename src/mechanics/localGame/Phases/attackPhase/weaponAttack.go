package attackPhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"math/rand"
)

func InitAttack(attacking *unit.Unit, target *coordinate.Coordinate, game *localGame.Game) *ResultBattle {

	if attacking.GetWeaponSlot() != nil && attacking.GetAmmoCount() > 0 && attacking.GetWeaponSlot().HP > 0 {
		attacking.GetWeaponSlot().AmmoQuantity -= 1
		return MapAttack(attacking, target, game, attacking.GetWeaponSlot())
	} else {
		if attacking.GetAmmoCount() == 0 {
			return &ResultBattle{Error: "no ammo"}
		}
		if attacking.GetWeaponSlot().HP == 0 {
			return &ResultBattle{Error: "weapon breaking"}
		}
		if attacking.GetWeaponSlot() == nil {
			return &ResultBattle{Error: "no find weapon"}
		}
	}

	return nil
}

func MapAttack(attacking *unit.Unit, target *coordinate.Coordinate, game *localGame.Game, weapon *detail.BodyWeaponSlot) *ResultBattle {
	unitCoordinate, _ := game.Map.GetCoordinate(attacking.Q, attacking.R)

	attackZone := coordinate.GetCoordinatesRadius(target, weapon.Ammo.AreaCovers)
	targetsUnit := make([]TargetUnit, 0)

	for _, attackCoordinate := range attackZone {
		targetUnit, find := game.GetUnit(attackCoordinate.Q, attackCoordinate.R)
		if find {

			if targetUnit.HP <= 0 {
				continue
			}

			var damage int

			if targetUnit.Q == target.Q && targetUnit.R == target.R {
				damage = calculateDamage(targetUnit, weapon.Ammo.MaxDamage, weapon.Ammo.MinDamage)
			} else {
				//т.к. не эпицентер атаки юниты получают только 50% урона
				damage = calculateDamage(targetUnit, weapon.Ammo.MaxDamage/2, weapon.Ammo.MinDamage/2)
			}

			targetUnit.HP -= damage

			// наносим урон по эквипу, если какойто экип сломался говорим об этом
			broken := breakingEquip(
				targetUnit,
				weapon.Weapon.EquipDamage+weapon.Ammo.EquipDamage,
				weapon.Weapon.EquipCriticalDamage+weapon.Ammo.EquipCriticalDamage,
			)

			targetsUnit = append(targetsUnit, TargetUnit{Unit: *targetUnit, Damage: damage, BreakingEquip: broken})
		}
	}

	return &ResultBattle{AttackUnit: *attacking, Target: *target, RotateTower: rotateTower(unitCoordinate, target), TargetUnits: targetsUnit, WeaponSlot: *weapon}
}

func calculateDamage(targetUnit *unit.Unit, maxDamage, minDamage int) int {

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

func breakingEquip(targetUnit *unit.Unit, damage int, chanceCrit int) bool { // если хотя бы 1 эквип сломался надо оповестить игроков об этом )

	weaponDamage := damage
	countDamage := 0

	if chanceCrit < rand.Intn(100) {
		// есть некий шанс нанести урон по снаряжение х2
		weaponDamage = damage * 2
	}

	// 25 < rand.Intn(100) сломать эквип 1 к 4, наверное...

	var breakingWeapon bool

	if targetUnit.GetWeaponSlot() != nil && targetUnit.GetWeaponSlot().HP-damage > 0 && 25 < rand.Intn(100) {
		targetUnit.GetWeaponSlot().HP -= weaponDamage
		countDamage++
		breakingWeapon = false
	} else {
		if targetUnit.GetWeaponSlot() != nil && 25 < rand.Intn(100) {
			targetUnit.GetWeaponSlot().HP = 0
			countDamage++
			breakingWeapon = true
		}
	}

	return breaking(targetUnit.Body.EquippingI, damage, chanceCrit, &countDamage) ||
		breaking(targetUnit.Body.EquippingII, damage, chanceCrit, &countDamage) ||
		breaking(targetUnit.Body.EquippingIII, damage, chanceCrit, &countDamage) ||
		breaking(targetUnit.Body.EquippingIV, damage, chanceCrit, &countDamage) ||
		breaking(targetUnit.Body.EquippingV, damage, chanceCrit, &countDamage) || breakingWeapon
}

func breaking(equip map[int]*detail.BodyEquipSlot, damage, chanceCrit int, countDamage *int) bool {

	if *countDamage > 2 {
		return false
	}

	equipDamage := damage

	if chanceCrit < rand.Intn(100) {
		// есть некий шанс нанести урон по снаряжение х2
		equipDamage = damage * 2
	}

	for _, equipSlot := range equip {

		if equipSlot.Equip != nil {

			if equipSlot.HP-damage > 0 && 25 < rand.Intn(100) {
				equipSlot.HP -= equipDamage
				*countDamage++
				return false
			} else {
				if 25 < rand.Intn(100) {
					equipSlot.HP = 0
					*countDamage++
					return true
				}
			}
		}

	}

	return false
}
