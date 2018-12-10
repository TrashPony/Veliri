package squadInventory

import (
	"../db/updateSquad"
	"../player"
)

func RemoveMSEquip(user *player.Player, numEquipSlot int, typeSlot int) {

	equipping := SelectType(typeSlot, user.GetSquad().MatherShip.Body)

	slotEquip, ok := equipping[numEquipSlot]

	if ok && slotEquip != nil && slotEquip.Equip != nil {

		user.GetSquad().Inventory.AddItem(slotEquip.Equip, "equip", slotEquip.Equip.ID, 1, slotEquip.HP, slotEquip.Equip.Size, slotEquip.Equip.MaxHP)
		slotEquip.Equip = nil

		user.GetSquad().MatherShip.CalculateParams()

		updateSquad.Squad(user.GetSquad())
	}
}

func RemoveUnitEquip(user *player.Player, numEquipSlot, typeSlot, numberUnitSlot int) {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]

	if ok && unitSlot.Unit != nil {
		equipping := SelectType(typeSlot, unitSlot.Unit.Body)

		slotEquip, ok := equipping[numEquipSlot]

		if ok && slotEquip != nil && slotEquip.Equip != nil {

			user.GetSquad().Inventory.AddItem(slotEquip.Equip, "equip", slotEquip.Equip.ID, 1, slotEquip.HP, slotEquip.Equip.Size, slotEquip.Equip.MaxHP)
			slotEquip.Equip = nil

			unitSlot.Unit.CalculateParams()

			updateSquad.Squad(user.GetSquad())
		}
	}
}
