package inventory

import (
	"../player"
	"../gameObjects/matherShip"
	"../gameObjects/unit"
	"../db/get"
	"../db/updateSquad"
)

func SetMSBody(user *player.Player, idBody, inventorySlot int) {
	body := user.GetSquad().Inventory[inventorySlot]

	if body != nil && body.ItemID == idBody {
		newBody := get.Body(idBody)

		if user.GetSquad().MatherShip == nil {
			user.GetSquad().MatherShip = &matherShip.MatherShip{}
		} else {
			if user.GetSquad().MatherShip.Body != nil {
				BodyRemove(user.GetSquad().Inventory, user.GetSquad().MatherShip.Body)
				user.GetSquad().MatherShip.Body = nil
			}
		}

		RemoveInventoryItem(1, user.GetSquad().Inventory[inventorySlot])
		user.GetSquad().MatherShip.Body = newBody

		user.GetSquad().MatherShip.Units = make(map[int]*matherShip.UnitSlot) // заполняем ячейки юнитов
		for _, slot := range user.GetSquad().MatherShip.Body.EquippingIV {
			unitSlot := matherShip.UnitSlot{}
			unitSlot.NumberSlot = slot.Number
			unitSlot.Unit = nil
			user.GetSquad().MatherShip.Units[slot.Number] = &unitSlot
		}

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

			RemoveInventoryItem(1, user.GetSquad().Inventory[inventorySlot])
			unitSlot.Unit.Body = newBody
		}
		updateSquad.Squad(user.GetSquad())
	}
}
