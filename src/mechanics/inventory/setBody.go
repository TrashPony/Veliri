package inventory

import (
	"../db/get"
	"../db/updateSquad"
	"../gameObjects/unit"
	"../player"
)

func SetMSBody(user *player.Player, idBody, inventorySlot int) {
	body := user.GetSquad().Inventory[inventorySlot]

	if body != nil && body.ItemID == idBody {
		newBody := get.Body(idBody)

		if user.GetSquad().MatherShip == nil {
			user.GetSquad().MatherShip = &unit.Unit{}
		} else {
			if user.GetSquad().MatherShip.Body != nil {
				BodyRemove(user.GetSquad().Inventory, user.GetSquad().MatherShip.Body, user.GetSquad().MatherShip.HP)
				user.GetSquad().MatherShip.Body = nil
			}
		}

		user.GetSquad().MatherShip.HP = body.HP                 // устанавливает колво хп как у тела
		user.GetSquad().MatherShip.Power = newBody.MaxPower     // устанавливаем мощьность как у тела
		user.GetSquad().MatherShip.ActionPoints = newBody.Speed // устанавливаем скорость как у тела

		RemoveInventoryItem(1, user.GetSquad().Inventory[inventorySlot])
		user.GetSquad().MatherShip.Body = newBody

		user.GetSquad().MatherShip.Units = make(map[int]*unit.Slot) // заполняем ячейки юнитов

		for _, slot := range user.GetSquad().MatherShip.Body.EquippingIV {
			unitSlot := unit.Slot{}
			unitSlot.NumberSlot = slot.Number
			unitSlot.Unit = nil
			user.GetSquad().MatherShip.Units[slot.Number] = &unitSlot
		}

		user.GetSquad().MatherShip.CalculateParams()

		updateSquad.Squad(user.GetSquad())
	}
}

func SetUnitBody(user *player.Player, idBody, inventorySlot, numberUnitSlot int) {
	body := user.GetSquad().Inventory[inventorySlot]

	if body != nil && body.ItemID == idBody {
		newBody := get.Body(idBody)

		unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]

		if ok {
			if unitSlot.Unit == nil {
				unitSlot.Unit = &unit.Unit{}
			} else {
				RemoveUnitBody(user, numberUnitSlot)
				unitSlot.Unit = &unit.Unit{}
			}

			unitSlot.Unit.HP = body.HP                 // устанавливает колво хп как у тела
			unitSlot.Unit.Power = newBody.MaxPower     // устанавливаем мощьность как у тела
			unitSlot.Unit.ActionPoints = newBody.Speed // устанавливаем скорость как у тела

			RemoveInventoryItem(1, user.GetSquad().Inventory[inventorySlot])
			unitSlot.Unit.Body = newBody
		}

		unitSlot.Unit.CalculateParams()

		updateSquad.Squad(user.GetSquad())
	}
}
