package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"log"
)

func WeaponsType() map[int]detail.Weapon {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" name," +
		" min_attack_range," +
		" range_attack," +
		" accuracy," +
		" ammo_capacity," +
		" artillery," +
		" power," +
		" max_hp," +
		" type," +
		" standard_size," +
		" size, " +
		" initiative " +
		"" +
		"FROM weapon_type")
	if err != nil {
		log.Fatal("get all type weapon: " + err.Error())
	}
	defer rows.Close()

	allType := make(map[int]detail.Weapon)

	for rows.Next() {
		var weapon detail.Weapon

		err := rows.Scan(&weapon.ID, &weapon.Name, &weapon.MinAttackRange, &weapon.Range, &weapon.Accuracy,
			&weapon.AmmoCapacity, &weapon.Artillery, &weapon.Power, &weapon.MaxHP, &weapon.Type, &weapon.StandardSize,
			&weapon.Size, &weapon.Initiative)
		if err != nil {
			log.Fatal("get all type scan weapon: " + err.Error())
		}

		allType[weapon.ID] = weapon
	}

	return allType
}
