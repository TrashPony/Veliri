package get

import (
	"../../../dbConnect"
	"../../gameObjects/ammo"
	"log"
)

func Ammo(id int) (resultAmmo *ammo.Ammo) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT id, name, type, type_attack, min_damage, max_damage, area_covers, standard_size, size "+
		"FROM ammunition_type "+
		" WHERE id=$1", id)
	if err != nil {
		log.Fatal("get ammo " + err.Error())
	}
	defer rows.Close()

	resultAmmo = &ammo.Ammo{}

	for rows.Next() {
		err := rows.Scan(&resultAmmo.ID, &resultAmmo.Name, &resultAmmo.Type, &resultAmmo.TypeAttack, &resultAmmo.MinDamage,
			&resultAmmo.MaxDamage, &resultAmmo.AreaCovers, &resultAmmo.StandardSize, &resultAmmo.Size)
		if err != nil {
			log.Fatal("get ammo " + err.Error())
		}
	}

	return resultAmmo
}
