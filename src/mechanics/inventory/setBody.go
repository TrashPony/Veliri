package inventory

import (
	"../player"
	"../gameObjects/matherShip"
	"../db/get"
	"../db/updateSquad"
)

func SetBody(user *player.Player, idBody, inventorySlot int) {
	body := user.GetSquad().Inventory[inventorySlot]

	if body.ItemID == idBody {
		newBody := get.Body(idBody)

		if user.GetSquad().MatherShip == nil {
			user.GetSquad().MatherShip = &matherShip.MatherShip{}
		} else {
			if user.GetSquad().MatherShip.Body != nil {
				BodyRemove(user.GetSquad().Inventory, user.GetSquad().MatherShip.Body)
			}
		}

		user.GetSquad().Inventory[inventorySlot].Item = nil // ставим итему nil что бы при обновление удалился слот из бд
		user.GetSquad().MatherShip.Body = newBody

		updateSquad.Squad(user.GetSquad()) // todo для теста опустил обновления в бд
	}
}
