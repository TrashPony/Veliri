package inventory

import (
	"../player"
	"../db/updateSquad"
)

func RemoveEquip(user *player.Player, numEquipSlot int, typeSlot int) {

	equipping := SelectType(typeSlot, user.GetSquad().MatherShip.Body)

	slotEquip, ok := equipping[numEquipSlot]

	if ok && slotEquip != nil && slotEquip.Equip != nil {

		AddItem(user.GetSquad().Inventory, slotEquip.Equip, "equip", slotEquip.Equip.ID, 1)
		slotEquip.Equip = nil

		updateSquad.Squad(user.GetSquad())
	}
}
