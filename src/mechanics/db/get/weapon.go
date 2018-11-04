package get

import (
	"../../../dbConnect"
	"../../gameObjects/detail"
	"log"
)

func Weapon(id int) (weapon *detail.Weapon) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT id,"+
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
		"FROM weapon_type "+
		"WHERE id=$1", id)
	if err != nil {
		log.Fatal("get weapon: " + err.Error())
	}
	defer rows.Close()

	weapon = &detail.Weapon{}

	for rows.Next() {
		err := rows.Scan(&weapon.ID, &weapon.Name, &weapon.MinAttackRange, &weapon.Range, &weapon.Accuracy,
			&weapon.AmmoCapacity, &weapon.Artillery, &weapon.Power, &weapon.MaxHP, &weapon.Type, &weapon.StandardSize,
			&weapon.Size, &weapon.Initiative)
		if err != nil {
			log.Fatal("get scan weapon: " + err.Error())
		}
	}

	return weapon
}
