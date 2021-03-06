package squad_inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func SetEquip(user *player.Player, idEquip, inventorySlot, numEquipSlot, typeEquipSlot int, unit *unit.Unit, source string) error {
	slot := getSlotBySource(user, inventorySlot, source)

	body := unit.Body

	if slot != nil && slot.ItemID == idEquip && slot.Type == "equip" {

		if slot.HP <= 0 {
			return errors.New("equip broken")
		}

		newEquip, _ := gameTypes.Equips.GetByID(idEquip)
		equipping := SelectType(typeEquipSlot, body)

		if equipping != nil && newEquip.TypeSlot == typeEquipSlot {

			equipSlot, ok := equipping[numEquipSlot]

			if (newEquip.Applicable == "ore" || newEquip.Applicable == "digger") && !equipSlot.Mining {
				return errors.New("wrong equip slot")
			}

			if ok && equipSlot.Type == typeEquipSlot {
				// писос, но тут смотрить можно ли поставить из расчета свободной энергии, или в замену текущему эквипу
				if (equipSlot.Equip != nil && body.MaxPower-body.GetUsePower()+equipSlot.Equip.Power >= newEquip.Power) ||
					(equipSlot.Equip == nil && body.MaxPower-body.GetUsePower() >= newEquip.Power) {

					if (unit.Body.GetUseCapacitySize()+newEquip.Size <= unit.Body.CapacitySize) || unit.Body.MotherShip {
						setEquip(equipSlot, user, newEquip, inventorySlot, slot.HP, unit, typeEquipSlot, source)
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
		return errors.New("wrong inventory slot")
	}
}

func SetUnitEquip(user *player.Player, idEquip, inventorySlot, numEquipSlot, typeEquipSlot, numberUnitSlot int, source string) error {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
	if ok && unitSlot.Unit != nil {
		return SetEquip(user, idEquip, inventorySlot, numEquipSlot, typeEquipSlot, unitSlot.Unit, source)
	} else {
		return errors.New("no unit")
	}
}

func setEquip(equipSlot *detail.BodyEquipSlot, user *player.Player, newEquip *equip.Equip, inventorySlot, hp int, unit *unit.Unit, typeSlot int, source string) {

	defer unit.CalculateParams()

	if equipSlot.Equip != nil {
		RemoveEquip(user, equipSlot.Number, typeSlot, unit, "storage")
	}

	equipSlot.HP = hp
	RemoveSlotBySource(user, inventorySlot, source, 1)

	update.Squad(user.GetSquad(), true) // без этого если в слоте есть снаряжение то оно не заменяется, а добавляется в бд

	equipSlot.Equip = newEquip
	equipSlot.InsertToDB = true
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
