package get

import (
	"../../gameObjects/detail"
	"../../gameObjects/ammo"
	"../../gameObjects/equip"
	"dbConnect"
	"log"
)

func GetBodiesType() []*detail.Body {
	rows, err := dbConnect.GetDBConnect().Query("SELECT "+
		"id, "+
		"name, "+
		"mother_ship, "+
		"speed, "+
		"initiative, "+
		"max_hp, "+
		"armor, "+
		"evasion_critical, "+
		"vulnerability_to_kinetics, "+
		"vulnerability_to_thermo, "+
		"vulnerability_to_explosion, "+
		"range_view, "+
		"accuracy, "+
		"max_power, "+
		"recovery_power, "+
		"wall_hack, "+
		"recovery_hp, "+
		"capacity_size, "+
		"standard_size, "+
		"standard_size_small, "+
		"standard_size_medium, "+
		"standard_size_big "+
		""+
		"FROM body_type")
	if err != nil {
		log.Fatal("get all type body: " + err.Error())
	}
	defer rows.Close()

	allType := make([]*detail.Body, 0)

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

		allType = append(allType, &body)

		BodySlots(&body)
	}

	return allType
}

func GetWeaponsType() []*detail.Weapon {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id,"+
		" name,"+
		" min_attack_range,"+
		" range_attack,"+
		" accuracy,"+
		" ammo_capacity,"+
		" artillery,"+
		" power,"+
		" max_hp,"+
		" type,"+
		" standard_size,"+
		" size, "+
		" initiative "+
		""+
		"FROM weapon_type")
	if err != nil {
		log.Fatal("get all type weapon: " + err.Error())
	}
	defer rows.Close()

	allType := make([]*detail.Weapon, 0)

	for rows.Next() {
		var weapon detail.Weapon

		err := rows.Scan(&weapon.ID, &weapon.Name, &weapon.MinAttackRange, &weapon.Range, &weapon.Accuracy,
			&weapon.AmmoCapacity, &weapon.Artillery, &weapon.Power, &weapon.MaxHP, &weapon.Type, &weapon.StandardSize,
			&weapon.Size, &weapon.Initiative)
		if err != nil {
			log.Fatal("get all type scan weapon: " + err.Error())
		}

		allType = append(allType, &weapon)
	}

	return allType
}

func GetAmmoType() []*ammo.Ammo {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" name," +
		" type," +
		" type_attack," +
		" min_damage," +
		" max_damage," +
		" area_covers," +
		" standard_size," +
		" size " +
		""+
		"FROM ammunition_type")
	if err != nil {
		log.Fatal("get all type ammo " + err.Error())
	}
	defer rows.Close()

	allType := make([]*ammo.Ammo, 0)

	for rows.Next() {
		var ammoType ammo.Ammo

		err := rows.Scan(&ammoType.ID, &ammoType.Name, &ammoType.Type, &ammoType.TypeAttack, &ammoType.MinDamage,
			&ammoType.MaxDamage, &ammoType.AreaCovers, &ammoType.StandardSize, &ammoType.Size)
		if err != nil {
			log.Fatal("get scan all type ammo " + err.Error())
		}

		allType = append(allType, &ammoType)
	}

	return allType
}

func GetEquipsType() []*equip.Equip {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id,"+
		" name,"+
		" active,"+
		" specification,"+
		" applicable,"+
		" region,"+
		" radius,"+
		" type_slot,"+
		" reload,"+
		" power,"+
		" use_power,"+
		" max_hp,"+
		" steps_time,"+
		" size, "+
		" initiative "+
		""+
		"FROM equipping_type ")
	if err != nil {
		log.Fatal("get all type equip " + err.Error())
	}
	defer rows.Close()

	allType := make([]*equip.Equip, 0)

	for rows.Next() {
		var equipType equip.Equip

		err := rows.Scan(&equipType.ID, &equipType.Name, &equipType.Active, &equipType.Specification,
			&equipType.Applicable, &equipType.Region, &equipType.Radius, &equipType.TypeSlot, &equipType.Reload,
			&equipType.Power, &equipType.UsePower, &equipType.MaxHP, &equipType.StepsTime, &equipType.Size,
			&equipType.Initiative)
		if err != nil {
			log.Fatal("get scan all type equip " + err.Error())
		}

		EffectsEquip(&equipType)

		allType = append(allType, &equipType)
	}

	return allType
}