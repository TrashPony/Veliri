package squadInventory

import (
	"../db/squad/update"
	"../player"
)

func RemoveMSEquip(user *player.Player, numEquipSlot int, typeSlot int) {

	equipping := SelectType(typeSlot, user.GetSquad().MatherShip.Body)

	slotEquip, ok := equipping[numEquipSlot]

	if ok && slotEquip != nil && slotEquip.Equip != nil {

		user.GetSquad().Inventory.AddItem(slotEquip.Equip, "equip", slotEquip.Equip.ID, 1, slotEquip.HP, slotEquip.Equip.Size, slotEquip.Equip.MaxHP)
		slotEquip.Equip = nil

		user.GetSquad().MatherShip.CalculateParams()

		go update.Squad(user.GetSquad(), true)
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

			go update.Squad(user.GetSquad(), true)
		}
	}
}
