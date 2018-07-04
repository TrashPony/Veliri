package inventory

import (
	"../player"
	"../gameObjects/matherShip"
	"../db/updateSquad"
	"../db/get"
)

func SetBody(user *player.Player, idBody, inventorySlot int) {
	body := user.GetSquad().Inventory[inventorySlot]

	if body.ItemID == idBody {
		body := get.Body(idBody)

		if user.GetSquad().MatherShip == nil {
			user.GetSquad().MatherShip = &matherShip.MatherShip{}
		} else {
			if user.GetSquad().MatherShip.Body != nil {
				BodyRemove(user.GetSquad().Inventory, user.GetSquad().MatherShip.Body)
			}
		}


		user.GetSquad().MatherShip.Body = body
		user.GetSquad().Inventory[inventorySlot].Item = nil // ставим итему nil что бы при обновление удалился слот

		delete(user.GetSquad().Inventory, inventorySlot)
	}

	updateSquad.Squad(user.GetSquad())
}
