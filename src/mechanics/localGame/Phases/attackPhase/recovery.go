package attackPhase

import (
	"../../../db/localGame/update"
	"../../../gameObjects/detail"
	"../../../gameObjects/unit"
	"../../../localGame"
)

func recovery(game *localGame.Game) {
	for _, qLine := range game.GetUnits() {
		for _, gameUnit := range qLine {

			if gameUnit.HP < 0 {
				game.DelUnit(gameUnit)

				for _, player := range game.GetPlayers() {
					// т.к. на одной точке не может стоять 2х юнитов в 1 момент времени эта операция безопасна :)
					player.DelUnit(gameUnit, true)
					player.DelHostileUnit(gameUnit.ID)
				}
				continue
			}

			// удаляем патроны если они кончились
			for _, weaponSlot := range gameUnit.Body.Weapons {
				if weaponSlot.Weapon != nil {
					if weaponSlot.AmmoQuantity <= 0 {
						weaponSlot.Ammo = nil
					}
				}
			}

			gameUnit.Target = nil
			gameUnit.QueueAttack = 0
			gameUnit.ActionPoints = gameUnit.Body.Speed
			gameUnit.Defend = false

			for _, effect := range gameUnit.Effects {
				if effect != nil {
					effect.StepsTime -= 1
				}
			}

			recoveryEquips(gameUnit)
			recoveryPower(gameUnit)

			update.UnitEffects(gameUnit)
		}
	}
}

func recoveryPower(gameUnit *unit.Unit) {
	if gameUnit.Power+gameUnit.RecoveryPower > gameUnit.Body.MaxPower {
		gameUnit.Power = gameUnit.Body.MaxPower
	} else {
		gameUnit.Power += gameUnit.RecoveryPower
	}
}

func recoveryEquips(gameUnit *unit.Unit) {
	recoveryEquip(gameUnit.Body.EquippingI)
	recoveryEquip(gameUnit.Body.EquippingII)
	recoveryEquip(gameUnit.Body.EquippingIII)
	recoveryEquip(gameUnit.Body.EquippingIV)
	recoveryEquip(gameUnit.Body.EquippingV)
}

func recoveryEquip(equip map[int]*detail.BodyEquipSlot) {
	for _, equipSlot := range equip {
		if equipSlot.Equip != nil {
			if equipSlot.StepsForReload-1 == 0 {
				equipSlot.StepsForReload = 0
				equipSlot.Used = false
			}
			if equipSlot.StepsForReload-1 > 0 {
				equipSlot.StepsForReload -= 1
			}
		}
	}
}