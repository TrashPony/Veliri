package squadInventory

import (
	"../db/squad/update"
	"../factories/storages"
	"../gameObjects/unit"
	"../player"
	"errors"
)

func RemoveEquip(user *player.Player, numEquipSlot int, typeSlot int, unit *unit.Unit) error {
	if user.InBaseID > 0 {

		equipping := SelectType(typeSlot, unit.Body)
		slot, ok := equipping[numEquipSlot]

		if ok && slot != nil && slot.Equip != nil {
			okAddItem := storages.Storages.AddItem(user.GetID(), user.InBaseID, slot.Equip, "equip",
				slot.Equip.ID, 1, slot.HP, slot.Equip.Size, slot.Equip.MaxHP)
			if okAddItem {
				slot.Equip = nil
				user.GetSquad().MatherShip.CalculateParams()
				go update.Squad(user.GetSquad(), true)
				return nil
			} else {
				return errors.New("add item error")
			}
		} else {
			return errors.New("no item")
		}
	} else {
		return errors.New("not in base")
	}
}

func RemoveUnitEquip(user *player.Player, numEquipSlot, typeSlot, numberUnitSlot int) error {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
	if ok && unitSlot.Unit != nil {
		return RemoveEquip(user, numEquipSlot, typeSlot, unitSlot.Unit)
	} else {
		return errors.New("no unit")
	}
}
