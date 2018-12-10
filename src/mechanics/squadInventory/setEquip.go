package squadInventory

import (
	"../db/get"
	"../db/updateSquad"
	"../gameObjects/detail"
	"../gameObjects/equip"
	"../player"
	"errors"
)

func SetMSEquip(user *player.Player, idEquip, inventorySlot, numEquipSlot, typeEquipSlot int) error {
	equipItem := user.GetSquad().Inventory.Slots[inventorySlot]

	msBody := user.GetSquad().MatherShip.Body

	if equipItem != nil && equipItem.ItemID == idEquip {
		newEquip := get.TypeEquip(idEquip)
		equipping := SelectType(typeEquipSlot, msBody)
		if equipping != nil {
			equipSlot, ok := equipping[numEquipSlot]
			if ok && equipSlot.Type == typeEquipSlot {

				// писос, но тут смотрить можно ли поставить из расчета свободной энергии, или в замену текущему эквипу
				if (equipSlot.Equip != nil && msBody.MaxPower-msBody.GetUsePower()+equipSlot.Equip.Power >= newEquip.Power) ||
					(equipSlot.Equip == nil && msBody.MaxPower-msBody.GetUsePower() >= newEquip.Power) {

					SetEquip(equipSlot, user, newEquip, inventorySlot, equipItem.HP)

					user.GetSquad().MatherShip.CalculateParams()
				} else {
					return errors.New("lacking power")
				}
			}
		}
	}
	return nil
}

func SetUnitEquip(user *player.Player, idEquip, inventorySlot, numEquipSlot, typeEquipSlot, numberUnitSlot int) error {
	equipItem := user.GetSquad().Inventory.Slots[inventorySlot]

	if equipItem.ItemID == idEquip {
		newEquip := get.TypeEquip(idEquip)
		unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
		if ok && unitSlot.Unit != nil {
			equipping := SelectType(typeEquipSlot, unitSlot.Unit.Body)
			if equipping != nil {
				equipSlot, ok := equipping[numEquipSlot]
				if ok && equipSlot.Type == typeEquipSlot {

					// писос, но тут смотрить можно ли поставить из расчета свободной энергии, или в замену текущему эквипу
					if (equipSlot.Equip != nil && unitSlot.Unit.Body.MaxPower-unitSlot.Unit.Body.GetUsePower()+equipSlot.Equip.Power >= newEquip.Power) ||
						(equipSlot.Equip == nil && unitSlot.Unit.Body.MaxPower-unitSlot.Unit.Body.GetUsePower() >= newEquip.Power) {

						if unitSlot.Unit.Body.GetUseCapacitySize()+newEquip.Size <= unitSlot.Unit.Body.CapacitySize {
							SetEquip(equipSlot, user, newEquip, inventorySlot, equipItem.HP)
							unitSlot.Unit.CalculateParams()
							return nil
						} else {
							return errors.New("lacking size")
						}
					} else {
						return errors.New("lacking power")
					}
				} else {
					return errors.New("wrong type slot")
				}
			} else {
				return errors.New("wrong equip slot")
			}
		} else {
			return errors.New("wrong unit")
		}
	} else {
		return errors.New("wrong inventory slot")
	}
}

func SetEquip(equipSlot *detail.BodyEquipSlot, user *player.Player, newEquip *equip.Equip, inventorySlot int, hp int) {

	if equipSlot.Equip != nil {
		user.GetSquad().Inventory.AddItem(equipSlot.Equip, "equip", equipSlot.Equip.ID, 1, equipSlot.HP, equipSlot.Equip.Size, equipSlot.Equip.MaxHP)
		equipSlot.Equip = nil
	}

	equipSlot.HP = hp
	user.GetSquad().Inventory.Slots[inventorySlot].RemoveItem(1)

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
