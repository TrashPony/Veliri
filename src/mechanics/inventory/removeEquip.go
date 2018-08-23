package inventory

import (
	"../player"
	"../db/updateSquad"
)

func RemoveMSEquip(user *player.Player, numEquipSlot int, typeSlot int) {

	equipping := SelectType(typeSlot, user.GetSquad().MatherShip.Body)

	slotEquip, ok := equipping[numEquipSlot]

	if ok && slotEquip != nil && slotEquip.Equip != nil {

		AddItem(user.GetSquad().Inventory, slotEquip.Equip, "equip", slotEquip.Equip.ID, 1)
		slotEquip.Equip = nil

		updateSquad.Squad(user.GetSquad())
	}
}

func RemoveUnitEquip(user *player.Player, numEquipSlot, typeSlot, numberUnitSlot int) {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]

	if ok && unitSlot.Unit != nil {
		equipping := SelectType(typeSlot, unitSlot.Unit.Body)

		slotEquip, ok := equipping[numEquipSlot]

		if ok && slotEquip != nil && slotEquip.Equip != nil {

			AddItem(user.GetSquad().Inventory, slotEquip.Equip, "equip", slotEquip.Equip.ID, 1)
			slotEquip.Equip = nil

			updateSquad.Squad(user.GetSquad())
		}
	}
}