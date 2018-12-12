package get

import (
	"../../../../dbConnect"
	"../../../factories/gameTypes"
	"../../../gameObjects/detail"
	"../../../gameObjects/unit"
	"log"
)

func BodyEquip(ship *unit.Unit) {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		"id_equipping, " +
		"slot_in_body, " +
		"type, type_slot, " +
		"quantity, " +
		"used, " +
		"steps_for_reload, " +
		"hp, " +
		"target " +
		""+
		"FROM squad_units_equipping " +
		"WHERE id_squad_unit = $1", ship.GetID())
	if err != nil {
		log.Fatal("get body equip" + err.Error())
	}
	defer rows.Close()

	var idEquip int
	var slot int
	var slotType int
	var equipType string
	var quantity int
	var used bool
	var stepsForReload int
	var hp int
	var target string

	for rows.Next() {
		err := rows.Scan(&idEquip, &slot, &equipType, &slotType, &quantity, &used, &stepsForReload, &hp, &target)
		if err != nil {
			log.Fatal("scan body equip " + err.Error())
		}

		if slotType == 1 {
			ParseTypeSlot(ship.GetBody().EquippingI, idEquip, slot, equipType, used, stepsForReload, hp, target)
		}
		if slotType == 2 {
			ParseTypeSlot(ship.GetBody().EquippingII, idEquip, slot, equipType, used, stepsForReload, hp, target)
		}
		if slotType == 3 {
			ParseTypeSlot(ship.GetBody().EquippingIII, idEquip, slot, equipType, used, stepsForReload, hp, target)
		}
		if slotType == 4 {
			ParseTypeSlot(ship.GetBody().EquippingIV, idEquip, slot, equipType, used, stepsForReload, hp, target)
		}
		if slotType == 5 {
			ParseTypeSlot(ship.GetBody().EquippingV, idEquip, slot, equipType, used, stepsForReload, hp, target)
		}

		if equipType == "weapon" || equipType == "ammo" {
			for _, bodyWeaponSlot := range ship.GetBody().Weapons {
				if equipType == "weapon" {
					bodyWeaponSlot.HP = hp

					bodyWeaponSlot.Weapon, _ = gameTypes.Weapons.GetByID(idEquip)
				}
				if equipType == "ammo" {
					bodyWeaponSlot.Ammo, _ = gameTypes.Ammo.GetByID(idEquip)
					bodyWeaponSlot.AmmoQuantity = quantity
				}
			}
		}
	}
}

func ParseTypeSlot(equipping map[int]*detail.BodyEquipSlot, idEquip int, slot int, equipType string, used bool, stepsForReload, hp int, target string) {
	for i, bodyEquipSlot := range equipping {
		if bodyEquipSlot.Number == slot {
			if equipType == "equip" {
				bodyEquipSlot.Used = used
				bodyEquipSlot.StepsForReload = stepsForReload
				bodyEquipSlot.HP = hp
				bodyEquipSlot.Target = ParseTarget(target)

				equipping[i].Equip, _ = gameTypes.Equips.GetByID(idEquip)
			}
		}
	}
}
