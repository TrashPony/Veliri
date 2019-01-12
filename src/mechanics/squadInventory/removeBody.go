package squadInventory

import (
	"../../mechanics/db/squad/update"
	"../factories/storages"
	"../gameObjects/detail"
	"../gameObjects/unit"
	"../player"
	"errors"
)

func RemoveMSBody(user *player.Player) error {
	if user.InBaseID > 0 {
		if user.GetSquad().MatherShip.Body != nil {

			for i, unitSlot := range user.GetSquad().MatherShip.Units {
				RemoveUnitBody(user, unitSlot.NumberSlot)
				delete(user.GetSquad().MatherShip.Units, i)
			}

			BodyRemove(user, user.GetSquad().MatherShip)
			user.GetSquad().MatherShip.Body = nil

			user.GetSquad().MatherShip.HP = 0 // обнулям статы т.к. юез тела их не может быть
			user.GetSquad().MatherShip.Power = 0

			user.GetSquad().MatherShip.CalculateParams()
			update.Squad(user.GetSquad(), true)
			return nil
		} else {
			return errors.New("no item")
		}
	} else {
		return errors.New("not in base")
	}
}

func RemoveUnitBody(user *player.Player, unitSlot int) error {
	if user.InBaseID > 0 {
		if user.GetSquad().MatherShip.Body != nil && user.GetSquad().MatherShip.Units[unitSlot].Unit != nil {
			if user.GetSquad().MatherShip.Units[unitSlot].Unit.Body != nil {

				BodyRemove(user, user.GetSquad().MatherShip.Units[unitSlot].Unit)
				user.GetSquad().MatherShip.Units[unitSlot].Unit = nil // если юниту убрали тело то юнит перестает существовать

				update.Squad(user.GetSquad(), true)
				return nil
			} else {
				return errors.New("unit no body")
			}
		} else {
			return errors.New("no unit")
		}
	} else {
		return errors.New("not in base")
	}
}

func BodyRemove(user *player.Player, unit *unit.Unit) {

	removeAllEquippingBody(user, unit, 1, unit.Body.EquippingI)
	removeAllEquippingBody(user, unit, 2, unit.Body.EquippingII)
	removeAllEquippingBody(user, unit, 3, unit.Body.EquippingIII)
	removeAllEquippingBody(user, unit, 4, unit.Body.EquippingIV)
	removeAllEquippingBody(user, unit, 5, unit.Body.EquippingV)

	for _, weaponSlot := range unit.Body.Weapons {
		if weaponSlot.Weapon != nil {
			RemoveWeapon(user, weaponSlot.Number, unit)
		}
		if weaponSlot.Ammo != nil {
			RemoveAmmo(user, weaponSlot.Number, unit)
		}
	}

	storages.Storages.AddItem(user.GetID(), user.InBaseID, unit.Body, "body", unit.Body.ID, 1, unit.HP,
		unit.Body.CapacitySize, unit.Body.MaxHP) // кидает боди в инвентарь
}

func removeAllEquippingBody(user *player.Player, unit *unit.Unit, typeSlot int, equipping map[int]*detail.BodyEquipSlot) {
	for _, equipSlot := range equipping {
		if equipSlot.Equip != nil {
			RemoveEquip(user, equipSlot.Number, typeSlot, unit)
		}
	}
}
