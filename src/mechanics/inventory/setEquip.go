package inventory

import (
	"../player"
	"../db/get"
	"../gameObjects/detail"
	"../db/updateSquad"
)

func SetEquip(user *player.Player, idEquip, inventorySlot, numEquipSlot, typeEquipSlot int)  {
	equip := user.GetSquad().Inventory[inventorySlot]

	if equip.ItemID == idEquip {
		newEquip := get.TypeEquip(idEquip)

		equipping := SelectType(typeEquipSlot, user.GetSquad().MatherShip.Body)

		if equipping != nil {
			equipSlot, ok := equipping[numEquipSlot]
			if ok {
				if equipSlot.Equip != nil {
					AddItem(user.GetSquad().Inventory, equipSlot.Equip, "equip", equipSlot.Equip.ID, 1)
				} else {
					equipSlot.InsertToDB = true
				}

				user.GetSquad().Inventory[inventorySlot].Item = nil
				equipSlot.Equip = newEquip

				updateSquad.Squad(user.GetSquad()) //todo для теста опустил обновления в бд
			}
		}
	}
}

func SelectType(typeEquipSlot int, body *detail.Body) map[int]*detail.BodyEquipSlot {
	if typeEquipSlot == 1 {
		return body.EquippingI
	}

	if typeEquipSlot == 2 {
		return body.EquippingII
	}

	if typeEquipSlot == 3 {
		return body.EquippingIII
	}

	if typeEquipSlot == 4 {
		return body.EquippingIV
	}

	if typeEquipSlot == 5 {
		return body.EquippingV
	}

	return nil
}
