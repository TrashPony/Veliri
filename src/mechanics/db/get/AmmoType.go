package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/ammo"
	"log"
)

func AmmoType() map[int]ammo.Ammo {
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
		"" +
		"FROM ammunition_type")
	if err != nil {
		log.Fatal("get all type ammo " + err.Error())
	}
	defer rows.Close()

	allType := make(map[int]ammo.Ammo)

	for rows.Next() {
		var ammoType ammo.Ammo

		err := rows.Scan(&ammoType.ID, &ammoType.Name, &ammoType.Type, &ammoType.TypeAttack, &ammoType.MinDamage,
			&ammoType.MaxDamage, &ammoType.AreaCovers, &ammoType.StandardSize, &ammoType.Size)
		if err != nil {
			log.Fatal("get scan all type ammo " + err.Error())
		}

		allType[ammoType.ID] = ammoType
	}

	return allType
}
