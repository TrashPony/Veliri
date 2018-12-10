package squadInventory

import (
	"../db/get"
	"../db/updateSquad"
	"../gameObjects/ammo"
	"../gameObjects/detail"
	"../player"
	"errors"
)

func SetMSAmmo(user *player.Player, idAmmo, inventorySlot, numEquipSlot int) error {
	ammoItem := user.GetSquad().Inventory.Slots[inventorySlot]

	if ammoItem.ItemID == idAmmo {
		newAmmo := get.Ammo(idAmmo)

		ammoSlot, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]
		if ok {
			err := SetAmmo(ammoSlot, user, newAmmo, inventorySlot)
			if err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("wrong weapon slot")
		}
	} else {
		return errors.New("wrong inventory slot")
	}
}

func SetUnitAmmo(user *player.Player, idAmmo, inventorySlot, numEquipSlot, numberUnitSlot int) error {
	ammoItem := user.GetSquad().Inventory.Slots[inventorySlot]

	if ammoItem.ItemID == idAmmo {
		newAmmo := get.Ammo(idAmmo)

		unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
		if ok && unitSlot.Unit != nil {
			ammoSlot, ok := unitSlot.Unit.Body.Weapons[numEquipSlot]
			if ok {
				err := SetAmmo(ammoSlot, user, newAmmo, inventorySlot)
				if err != nil {
					return err
				}
				return nil
			} else {
				return errors.New("wrong weapon slot")
			}
		} else {
			return errors.New("wrong unit slot")
		}
	} else {
		return errors.New("wrong inventory slot")
	}
}

func SetAmmo(ammoSlot *detail.BodyWeaponSlot, user *player.Player, newAmmo *ammo.Ammo, inventorySlot int) error {
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
		user.GetSquad().Inventory.AddItem(ammoSlot.Ammo, "ammo", ammoSlot.Ammo.ID, ammoSlot.AmmoQuantity, 1, ammoSlot.Ammo.Size, 1)
	}

	ammoSlot.Ammo = newAmmo
	ammoSlot.AmmoQuantity = user.GetSquad().Inventory.Slots[inventorySlot].RemoveItem(ammoSlot.Weapon.AmmoCapacity)

	updateSquad.Squad(user.GetSquad())

	return nil
}
