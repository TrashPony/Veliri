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
				// todo замена тела
			}

			RemoveInventoryItem(1, user.GetSquad().Inventory[inventorySlot])
			unitSlot.Unit.Body = newBody
		}
		updateSquad.Squad(user.GetSquad())
	}
}
