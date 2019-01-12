package squadInventory

import (
	"../db/squad/update"
	"../factories/gameTypes"
	"../gameObjects/unit"
	"../player"
	"errors"
)

func SetAmmo(user *player.Player, idAmmo, inventorySlot, numEquipSlot int, unit *unit.Unit, source string) error {
	if user.InBaseID > 0 {

		slot := getSlotBySource(user, inventorySlot, source)

		if slot != nil && slot.Item != nil && slot.ItemID == idAmmo && slot.Type == "ammo" {
			newAmmo, _ := gameTypes.Ammo.GetByID(idAmmo)

			ammoSlot, ok := unit.Body.Weapons[numEquipSlot]
			if ok {

				if ammoSlot.Weapon == nil {
					return errors.New("no weapon") // если нет оружия ему нельзя поставить боеприпас
				}

				if ammoSlot.Weapon.StandardSize != newAmmo.StandardSize {
					return errors.New("wrong standard size weapon")
				}

				if ammoSlot.Weapon.Type != newAmmo.Type {
					return errors.New("wrong type weapon")
				}

				if ammoSlot.Ammo != nil {
					RemoveAmmo(user, numEquipSlot, unit)
				}

				ammoSlot.Ammo = newAmmo
				ammoSlot.AmmoQuantity = slot.RemoveItemBySlot(ammoSlot.Weapon.AmmoCapacity)

				go update.Squad(user.GetSquad(), true)

				return nil

			} else {
				return errors.New("wrong weapon slot")
			}
		} else {
			return errors.New("wrong inventory slot")
		}
	} else {
		return errors.New("not in base")
	}
}

func SetUnitAmmo(user *player.Player, idAmmo, inventorySlot, numEquipSlot, numberUnitSlot int, source string) error {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
	if ok && unitSlot.Unit != nil {
		return SetAmmo(user, idAmmo, inventorySlot, numEquipSlot, unitSlot.Unit, source)
	} else {
		return errors.New("no unit")
	}
}
