package inventory

import (
	"../gameObjects/detail"
	"../gameObjects/squad"
)

func BodyRemove(inventory map[int]*squad.InventorySlot, Body *detail.Body) {
	removeAllEquippingBody(inventory, Body.EquippingI)
	removeAllEquippingBody(inventory, Body.EquippingII)
	removeAllEquippingBody(inventory, Body.EquippingIII)
	removeAllEquippingBody(inventory, Body.EquippingIV)
	removeAllEquippingBody(inventory, Body.EquippingV)

	for _, weaponSlot := range Body.Weapons {

		if weaponSlot.Weapon != nil {
			AddItem(inventory, weaponSlot.Weapon, "weapon", weaponSlot.Weapon.ID, 1)
		}

		if weaponSlot.Ammo != nil {
			AddItem(inventory, weaponSlot.Ammo, "ammo", weaponSlot.Ammo.ID, weaponSlot.AmmoQuantity)
		}

		weaponSlot.Ammo = nil
		weaponSlot.Weapon = nil
	}

	AddItem(inventory, Body, "body", Body.ID, 1) // кидает боди в инвентарь
}

func removeAllEquippingBody(inventory map[int]*squad.InventorySlot, equipping map[int]*detail.BodyEquipSlot) {
	for _, equipSlot := range equipping {
		if equipSlot.Equip != nil{
			AddItem(inventory, equipSlot.Equip, "equip", equipSlot.Equip.ID, 1)
			equipSlot.Equip = nil
		}
	}
}
