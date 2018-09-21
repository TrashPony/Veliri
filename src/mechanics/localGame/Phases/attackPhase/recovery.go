package attackPhase

import (
	"../../../localGame"
	"../../../gameObjects/unit"
	"../../../gameObjects/detail"
)

func recovery(game *localGame.Game) {
	for _, qLine := range game.GetUnits() {
		for _, gameUnit := range qLine {

			if gameUnit.HP < 0 {
				game.DelUnit(gameUnit)
				continue
			}

			gameUnit.UseEquip = false
			gameUnit.Action = false
			gameUnit.Target = nil
			gameUnit.QueueAttack = 0

			for _, effect := range gameUnit.Effects {
				effect.StepsTime -= 1
			}

			recoveryEquips(gameUnit)
			recoveryPower(gameUnit)


			for _, player := range game.GetPlayers() {
				// т.к. на одной точке не может стоять 2х юнитов в 1 момент времени эта операция безопасна :)
				player.DelUnit(gameUnit, true)
				player.DelHostileUnit(gameUnit.ID)
			}
		}
	}
}

func recoveryPower(gameUnit *unit.Unit)  {
	// todo
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
			equipSlot.StepsForReload -= 1
		}
	}
}
