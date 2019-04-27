package squadInventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/remove"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func RemoveMSBody(user *player.Player) error {
	// при удаление мс происходит удаление отряда

	if user.InBaseID > 0 {
		if user.GetSquad() != nil && user.GetSquad().MatherShip.Body != nil {

			for i, unitSlot := range user.GetSquad().MatherShip.Units {
				RemoveUnitBody(user, unitSlot.NumberSlot, false)
				delete(user.GetSquad().MatherShip.Units, i)
			}

			BodyRemove(user, user.GetSquad().MatherShip, false)

			for _, inventorySlot := range user.GetSquad().Inventory.Slots {
				storages.Storages.AddSlot(user.GetID(), user.InBaseID, inventorySlot)
			}

			remove.Squad(user.GetSquad())
			user.RemoveSquadsByID(user.GetSquad().ID)
			user.SetSquad(nil)

			return nil
		} else {
			return errors.New("no item")
		}
	} else {
		return errors.New("not in base")
	}
}

func RemoveUnitBody(user *player.Player, unitSlot int, updateDB bool) error {
	if user.InBaseID > 0 {
		if user.GetSquad().MatherShip.Body != nil && user.GetSquad().MatherShip.Units[unitSlot].Unit != nil {
			if user.GetSquad().MatherShip.Units[unitSlot].Unit.Body != nil {

				BodyRemove(user, user.GetSquad().MatherShip.Units[unitSlot].Unit, updateDB)
				user.GetSquad().MatherShip.Units[unitSlot].Unit = nil // если юниту убрали тело то юнит перестает существовать

				if updateDB {
					update.Squad(user.GetSquad(), true)
				}
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

func BodyRemove(user *player.Player, unit *unit.Unit, updateDB bool) {

	removeAllEquippingBody(user, unit, 1, unit.Body.EquippingI, updateDB)
	removeAllEquippingBody(user, unit, 2, unit.Body.EquippingII, updateDB)
	removeAllEquippingBody(user, unit, 3, unit.Body.EquippingIII, updateDB)
	removeAllEquippingBody(user, unit, 4, unit.Body.EquippingIV, updateDB)
	removeAllEquippingBody(user, unit, 5, unit.Body.EquippingV, updateDB)

	for _, weaponSlot := range unit.Body.Weapons {
		if weaponSlot.Weapon != nil {
			RemoveWeapon(user, weaponSlot.Number, unit, "storage", updateDB)
		}
		if weaponSlot.Ammo != nil {
			RemoveAmmo(user, weaponSlot.Number, unit, "storage", updateDB)
		}
	}

	for _, thoriumSlot := range unit.Body.ThoriumSlots {
		if thoriumSlot.Count != 0 {
			RemoveThorium(user, thoriumSlot.Number, updateDB)
		}
	}

	storages.Storages.AddItem(user.GetID(), user.InBaseID, unit.Body, "body", unit.Body.ID, 1, unit.HP,
		unit.Body.CapacitySize, unit.Body.MaxHP) // кидает боди в инвентарь
}

func removeAllEquippingBody(user *player.Player, unit *unit.Unit, typeSlot int, equipping map[int]*detail.BodyEquipSlot, updateDB bool) {
	for _, equipSlot := range equipping {
		if equipSlot.Equip != nil {
			RemoveEquip(user, equipSlot.Number, typeSlot, unit, "storage", updateDB)
		}
	}
}
