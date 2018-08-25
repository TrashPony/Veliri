package inventory

import (
	"../player"
	"../db/get"
	"../gameObjects/detail"
	"../gameObjects/equip"
	"../db/updateSquad"
)

func SetMSEquip(user *player.Player, idEquip, inventorySlot, numEquipSlot, typeEquipSlot int) {
	equipItem := user.GetSquad().Inventory[inventorySlot]

	if equipItem.ItemID == idEquip {
		newEquip := get.TypeEquip(idEquip)

		equipping := SelectType(typeEquipSlot, user.GetSquad().MatherShip.Body)

		if equipping != nil {
			equipSlot, ok := equipping[numEquipSlot]
			if ok && equipSlot.Type == typeEquipSlot {
				SetEquip(equipSlot, user, newEquip, inventorySlot)
			}
		}
	}
}

func SetUnitEquip(user *player.Player, idEquip, inventorySlot, numEquipSlot, typeEquipSlot, numberUnitSlot int) {
	equipItem := user.GetSquad().Inventory[inventorySlot]

	if equipItem.ItemID == idEquip {
		newEquip := get.TypeEquip(idEquip)
		unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
		if ok && unitSlot.Unit != nil {
			equipping := SelectType(typeEquipSlot, unitSlot.Unit.Body)
			if equipping != nil {
				equipSlot, ok := equipping[numEquipSlot]
				if ok && equipSlot.Type == typeEquipSlot{
					SetEquip(equipSlot, user, newEquip, inventorySlot)
				}
			}
		}
	}
}

func SetEquip(equipSlot *detail.BodyEquipSlot, user *player.Player, newEquip *equip.Equip, inventorySlot int)  {
	if equipSlot.Equip != nil {
		AddItem(user.GetSquad().Inventory, equipSlot.Equip, "equip", equipSlot.Equip.ID, 1)
		equipSlot.Equip = nil
	}

	RemoveInventoryItem(1, user.GetSquad().Inventory[inventorySlot])

	updateSquad.Squad(user.GetSquad()) // без этого если в слоте есть снаряжение то оно не заменяется, а добавляется в бд

	equipSlot.Equip = newEquip
	equipSlot.InsertToDB = true
	updateSquad.Squad(user.GetSquad())
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
