package squad_inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func RemoveEquip(user *player.Player, numEquipSlot int, typeSlot int, unit *unit.Unit, dst string) error {
	if user.InBaseID > 0 {

		defer unit.CalculateParams()

		equipping := SelectType(typeSlot, unit.Body)
		slot, ok := equipping[numEquipSlot]

		if ok && slot != nil && slot.Equip != nil {
			if dst == "squadInventory" {
				// TODO
				return nil
			}

			if dst == "storage" {
				okAddItem := storages.Storages.AddItem(user.GetID(), user.InBaseID, slot.Equip, "equip",
					slot.Equip.ID, 1, slot.HP, slot.Equip.Size, slot.Equip.MaxHP, false)
				if okAddItem {
					slot.Equip = nil
					user.GetSquad().MatherShip.CalculateParams()

					return nil
				} else {
					return errors.New("add item error")
				}
			}
		} else {
			return errors.New("no item")
		}
	} else {
		return errors.New("not in base")
	}

	return errors.New("unknown error")
}

func RemoveUnitEquip(user *player.Player, numEquipSlot, typeSlot, numberUnitSlot int, dst string) error {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
	if ok && unitSlot.Unit != nil {
		return RemoveEquip(user, numEquipSlot, typeSlot, unitSlot.Unit, dst)
	} else {
		return errors.New("no unit")
	}
}
