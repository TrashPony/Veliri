package get

import (
	"../../../dbConnect"
	"../../gameObjects/detail"
	"log"
)

func BodiesType() map[int]detail.Body {
	rows, err := dbConnect.GetDBConnect().Query("SELECT " +
		"id, " +
		"name, " +
		"mother_ship, " +
		"speed, " +
		"initiative, " +
		"max_hp, " +
		"armor, " +
		"evasion_critical, " +
		"vulnerability_to_kinetics, " +
		"vulnerability_to_thermo, " +
		"vulnerability_to_explosion, " +
		"range_view, " +
		"accuracy, " +
		"max_power, " +
		"recovery_power, " +
		"wall_hack, " +
		"recovery_hp, " +
		"capacity_size, " +
		"standard_size, " +
		"standard_size_small, " +
		"standard_size_medium, " +
		"standard_size_big " +
		"" +
		"FROM body_type")
	if err != nil {
		log.Fatal("get all type body: " + err.Error())
	}
	defer rows.Close()

	allType := make(map[int]detail.Body)

	for rows.Next() {
		var body detail.Body

		err = rows.Scan(
			&body.ID,
			&body.Name,
			&body.MotherShip,
			&body.Speed,
			&body.Initiative,
			&body.MaxHP,
			&body.Armor,
			&body.EvasionCritical,
			&body.VulToKinetics,
			&body.VulToThermo,
			&body.VulToExplosion,
			&body.RangeView,
			&body.Accuracy,
			&body.MaxPower,
			&body.RecoveryPower,
			&body.WallHack,
			&body.RecoveryHP,
			&body.CapacitySize,
			&body.StandardSize,
			&body.StandardSizeSmall,
			&body.StandardSizeMedium,
			&body.StandardSizeBig,
		)
		if err != nil {
			log.Fatal("get scan all type body: " + err.Error())
		}

		BodySlots(&body)

		allType[body.ID] = body
	}

	return allType
}

func BodySlots(body *detail.Body) {
	rows, err := dbConnect.GetDBConnect().Query("SELECT type_slot, number_slot, weapon, weapon_type "+
		"FROM body_slots "+
		"WHERE id_body = $1", body.ID)
	if err != nil {
		log.Fatal("get body slot " + err.Error())
	}
	defer rows.Close()

	body.EquippingI = make(map[int]*detail.BodyEquipSlot) // чето как то пиздец
	body.EquippingII = make(map[int]*detail.BodyEquipSlot)
	body.EquippingIII = make(map[int]*detail.BodyEquipSlot)
	body.EquippingIV = make(map[int]*detail.BodyEquipSlot)
	body.EquippingV = make(map[int]*detail.BodyEquipSlot)

	body.Weapons = make(map[int]*detail.BodyWeaponSlot)

	for rows.Next() {
		var slotType int
		var slotNumber int
		var slotWeapon bool
		var slotWeaponType string

		err := rows.Scan(&slotType, &slotNumber, &slotWeapon, &slotWeaponType)
		if err != nil {
			log.Fatal("get body slot " + err.Error())
		}

		if slotWeapon {
			weaponSlot := detail.BodyWeaponSlot{Type: slotType, Number: slotNumber, WeaponType: slotWeaponType}
			body.Weapons[slotNumber] = &weaponSlot
		} else {
			if slotType == 1 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
				body.EquippingI[slotNumber] = &equipSlot
			}
			if slotType == 2 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
				body.EquippingII[slotNumber] = &equipSlot
			}
			if slotType == 3 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
				body.EquippingIII[slotNumber] = &equipSlot
			}
			if slotType == 4 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
				body.EquippingIV[slotNumber] = &equipSlot
			}
			if slotType == 5 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber}
				body.EquippingV[slotNumber] = &equipSlot
			}
		}
	}
}
