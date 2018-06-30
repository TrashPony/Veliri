package get

import (
	"log"
	"../../../dbConnect"
	"../../gameObjects/ammo"
)

func Ammo(id int) (resultAmmo *ammo.Ammo) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT id, name, type, type_attack, damage, area_covers " +
		"FROM ammunition_type " +
		" WHERE id=$1", id)
	if err != nil {
		log.Fatal("get ammo " + err.Error())
	}
	defer rows.Close()

	resultAmmo = &ammo.Ammo{}

	for rows.Next() {
		err := rows.Scan(&resultAmmo.ID, &resultAmmo.Name, &resultAmmo.Type, &resultAmmo.TypeAttack, &resultAmmo.Damage, &resultAmmo.AreaCovers)
		if err != nil {
			log.Fatal("get ammo " + err.Error())
		}
	}

	return resultAmmo
}
