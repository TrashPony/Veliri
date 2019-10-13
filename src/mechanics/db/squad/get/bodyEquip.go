package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"log"
)

func BodyEquip(ship *unit.Unit) {
	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		"id_equipping, "+
		"slot_in_body, "+
		"type, type_slot, "+
		"quantity, "+
		"used, "+
		"steps_for_reload, "+
		"hp "+
		""+
		"FROM squad_units_equipping "+
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

	for rows.Next() {
		err := rows.Scan(&idEquip, &slot, &equipType, &slotType, &quantity, &used, &stepsForReload, &hp)
		if err != nil {
			log.Fatal("scan body equip " + err.Error())
		}

		if slotType == 1 {
			ParseTypeSlot(ship.GetBody().EquippingI, idEquip, slot, equipType, used, stepsForReload, hp)
		}
		if slotType == 2 {
			ParseTypeSlot(ship.GetBody().EquippingII, idEquip, slot, equipType, used, stepsForReload, hp)
		}
		if slotType == 3 {
			ParseTypeSlot(ship.GetBody().EquippingIII, idEquip, slot, equipType, used, stepsForReload, hp)
		}
		if slotType == 4 {
			ParseTypeSlot(ship.GetBody().EquippingIV, idEquip, slot, equipType, used, stepsForReload, hp)
		}
		if slotType == 5 {
			ParseTypeSlot(ship.GetBody().EquippingV, idEquip, slot, equipType, used, stepsForReload, hp)
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

func ParseTypeSlot(equipping map[int]*detail.BodyEquipSlot, idEquip int, slot int, equipType string, used bool, stepsForReload, hp int) {
	for i, bodyEquipSlot := range equipping {
		if bodyEquipSlot.Number == slot {
			if equipType == "equip" {
				bodyEquipSlot.Used = used
				bodyEquipSlot.StepsForReload = stepsForReload
				bodyEquipSlot.HP = hp

				equipping[i].Equip, _ = gameTypes.Equips.GetByID(idEquip)
			}
		}
	}
}
