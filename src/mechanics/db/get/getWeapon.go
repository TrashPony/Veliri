package get

import (
	"../../../dbConnect"
	"../../gameObjects/detail"
	"log"
)

func Weapons() (weapons []detail.Weapon) {
	weapons = make([]detail.Weapon, 0)

	rows, err := dbConnect.GetDBConnect().Query("select * from weapon_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var weapon detail.Weapon

	for rows.Next() {
		err := rows.Scan(&weapon.Id, &weapon.Name, &weapon.Type, &weapon.Weight, &weapon.Damage, &weapon.MinAttackRange, &weapon.Range, &weapon.Accuracy, &weapon.AreaCovers)
		if err != nil {
			log.Fatal("get weapons" + err.Error())
		}
		weapons = append(weapons, weapon)
	}

	return weapons
}

func Weapon(id int) (weapon *detail.Weapon) {

	rows, err := dbConnect.GetDBConnect().Query("select * from weapon_type where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	weapon = &detail.Weapon{}

	for rows.Next() {
		err := rows.Scan(&weapon.Id, &weapon.Name, &weapon.Type, &weapon.Weight, &weapon.Damage, &weapon.MinAttackRange, &weapon.Range, &weapon.Accuracy, &weapon.AreaCovers)
		if err != nil {
			log.Fatal("get weapon" + err.Error())
		}
	}

	return weapon
}
