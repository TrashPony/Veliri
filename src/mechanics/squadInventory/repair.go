package squadInventory

import (
	"../db/squad/update"
	"../factories/gameTypes"
	"../gameObjects/detail"
	"../player"
)

func ItemsRepair(user *player.Player) error {
	// todo высчитывание денег и рем комплектов необходимых, отнимание денег или рем комплектов )

	//if user.GetCredits() > 100 {
	for _, slot := range user.GetSquad().Inventory.Slots {
		if slot.Type == "body" {
			body, _ := gameTypes.Bodies.GetByID(slot.ItemID)
			slot.HP = body.MaxHP
		}

		if slot.Type == "equip" {
			equip, _ := gameTypes.Equips.GetByID(slot.ItemID)
			slot.HP = equip.MaxHP
		}

		if slot.Type == "weapon" {
			weapon, _ := gameTypes.Weapons.GetByID(slot.ItemID)
			slot.HP = weapon.MaxHP
		}
	}

	update.Squad(user.GetSquad(), true)
	return nil
	//} else {
	//	return errors.New("no credits")
	//}
}

func EquipRepair(user *player.Player) error {
	// todo высчитывание денег и рем комплектов необходимых, отнимание денег или рем комплектов )

	//if user.GetCredits() < 100 {

	var repairEquip = func(equips map[int]*detail.BodyEquipSlot) {
		for _, equip := range equips {
			if equip.Equip != nil {
				equip.HP = equip.Equip.MaxHP
			}
		}
	}

	var repairWeapon = func(weapons map[int]*detail.BodyWeaponSlot) {
		for _, weapon := range weapons {
			if weapon.Weapon != nil {
				weapon.HP = weapon.Weapon.MaxHP
			}
		}
	}

	user.GetSquad().MatherShip.HP = user.GetSquad().MatherShip.Body.MaxHP

	repairEquip(user.GetSquad().MatherShip.Body.EquippingI)
	repairEquip(user.GetSquad().MatherShip.Body.EquippingII)
	repairEquip(user.GetSquad().MatherShip.Body.EquippingIII)
	repairEquip(user.GetSquad().MatherShip.Body.EquippingIV)
	repairEquip(user.GetSquad().MatherShip.Body.EquippingV)

	repairWeapon(user.GetSquad().MatherShip.Body.Weapons)

	for _, unit := range user.GetSquad().MatherShip.Units {
		if unit.Unit != nil {
			unit.Unit.HP = unit.Unit.Body.MaxHP

			repairEquip(unit.Unit.Body.EquippingI)
			repairEquip(unit.Unit.Body.EquippingII)
			repairEquip(unit.Unit.Body.EquippingIII)
			repairEquip(unit.Unit.Body.EquippingIV)
			repairEquip(unit.Unit.Body.EquippingV)

			repairWeapon(unit.Unit.Body.Weapons)
		}
	}

	update.Squad(user.GetSquad(), true)
	return nil
	//} else {
	//	return errors.New("no credits")
	//}
}

func AllRepair(user *player.Player) error {
	// todo высчитывание денег и рем комплектов необходимых, отнимание денег или рем комплектов )

	//if user.GetCredits() < 100 {
	ItemsRepair(user)
	EquipRepair(user)

	return nil
	//} else {
	//return errors.New("no credits")
	//}
}
