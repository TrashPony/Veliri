package detailUnit

import (
	"log"
	"../dbConnect"
)

type Weapon struct {
	Id             int    `json:"id"`
	Name		   string `json:"name"`
	Type           string `json:"type"`
	Weight         int    `json:"weight"`
	Damage         int    `json:"damage"`
	MinAttackRange int    `json:"min_attack_range"`
	Range		   int    `json:"range"`
	Accuracy       int    `json:"accuracy"`
	AreaCovers     int    `json:"area_covers"`
}

func GetWeapons() (weapons []Weapon) {
	weapons = make([]Weapon, 0)

	rows, err := dbConnect.GetDBConnect().Query("select * from weapon_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var weapon Weapon

	for rows.Next() {
		err := rows.Scan(&weapon.Id, &weapon.Name, &weapon.Type, &weapon.Weight, &weapon.Damage, &weapon.MinAttackRange, &weapon.Range, &weapon.Accuracy, &weapon.AreaCovers)
		if err != nil {
			log.Fatal("get weapons" + err.Error())
		}
		weapons = append(weapons, weapon)
	}

	return weapons
}

func GetWeapon(id int) (weapon *Weapon) {

	rows, err := dbConnect.GetDBConnect().Query("select * from weapon_type where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	weapon = &Weapon{}

	for rows.Next() {
		err := rows.Scan(&weapon.Id, &weapon.Name, &weapon.Type, &weapon.Weight, &weapon.Damage, &weapon.MinAttackRange, &weapon.Range, &weapon.Accuracy, &weapon.AreaCovers)
		if err != nil {
			log.Fatal("get weapon" + err.Error())
		}
	}

	return weapon
}

