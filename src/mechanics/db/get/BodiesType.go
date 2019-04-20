package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
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
		"standard_size_big, " +
		"body_front_radius, " +
		"body_left_front_angle, " +
		"body_right_front_angle, " +
		"body_back_radius, " +
		"body_left_back_angle, " +
		"body_right_back_angle, " +
		"body_side_radius," +
		"height," +
		"width " +
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
			&body.FrontRadius,
			&body.LeftFrontAngle,
			&body.RightFrontAngle,
			&body.BackRadius,
			&body.LeftBackAngle,
			&body.RightBackAngle,
			&body.SideRadius,
			&body.Height,
			&body.Width,
		)
		if err != nil {
			log.Fatal("get scan all type body: " + err.Error())
		}

		BodySlots(&body)
		BodyThoriumSlots(&body)

		allType[body.ID] = body
	}

	return allType
}

func BodyThoriumSlots(body *detail.Body) {
	rows, err := dbConnect.GetDBConnect().Query("SELECT number_slot, max_thorium "+
		"FROM body_thorium_slots "+
		"WHERE id_body = $1", body.ID)
	if err != nil {
		log.Fatal("get body thorium slot " + err.Error())
	}
	defer rows.Close()

	body.ThoriumSlots = make(map[int]*detail.ThoriumSlot)

	for rows.Next() {
		var thoriumSlot detail.ThoriumSlot

		err := rows.Scan(&thoriumSlot.Number, &thoriumSlot.MaxCount)
		if err != nil {
			log.Fatal("get body thorium slot " + err.Error())
		}

		body.ThoriumSlots[thoriumSlot.Number] = &thoriumSlot
	}
}

func BodySlots(body *detail.Body) {
	rows, err := dbConnect.GetDBConnect().Query("SELECT type_slot, number_slot, weapon, weapon_type, standard_size, x_attach, y_attach, mining "+
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
		var slotStandardSize int
		var xAttach, yAttach int
		var mining bool

		err := rows.Scan(&slotType, &slotNumber, &slotWeapon, &slotWeaponType, &slotStandardSize, &xAttach, &yAttach, &mining)
		if err != nil {
			log.Fatal("get body slot " + err.Error())
		}

		if slotWeapon {
			weaponSlot := detail.BodyWeaponSlot{Type: slotType, Number: slotNumber, WeaponType: slotWeaponType, XAttach: xAttach, YAttach: yAttach}
			body.Weapons[slotNumber] = &weaponSlot
		} else {
			if slotType == 1 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber, XAttach: xAttach, YAttach: yAttach, Mining: mining}
				body.EquippingI[slotNumber] = &equipSlot
			}
			if slotType == 2 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber, XAttach: xAttach, YAttach: yAttach, Mining: mining}
				body.EquippingII[slotNumber] = &equipSlot
			}
			if slotType == 3 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber, XAttach: xAttach, YAttach: yAttach, Mining: mining}
				body.EquippingIII[slotNumber] = &equipSlot
			}
			if slotType == 4 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber, StandardSize: slotStandardSize, XAttach: xAttach, YAttach: yAttach, Mining: mining}
				body.EquippingIV[slotNumber] = &equipSlot
			}
			if slotType == 5 {
				equipSlot := detail.BodyEquipSlot{Type: slotType, Number: slotNumber, XAttach: xAttach, YAttach: yAttach, Mining: mining}
				body.EquippingV[slotNumber] = &equipSlot
			}
		}
	}
}
