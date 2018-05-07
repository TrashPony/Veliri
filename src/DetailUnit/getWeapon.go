package DetailUnit

import "log"

type Weapon struct {
	Id             int    `json:"id"`
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

	rows, err := db.Query("select * from weapon_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var weapon Weapon

	for rows.Next() {
		err := rows.Scan(&weapon.Id, &weapon.Type, &weapon.Weight, &weapon.Damage, &weapon.MinAttackRange, &weapon.Range, &weapon.Accuracy, &weapon.AreaCovers)
		if err != nil {
			log.Fatal("get weapons" + err.Error())
		}
		weapons = append(weapons, weapon)
	}

	return weapons
}

func GetWeapon(id int) (weapon *Weapon) {

	rows, err := db.Query("select * from weapon_type where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	weapon = &Weapon{}

	for rows.Next() {
		err := rows.Scan(&weapon.Id, &weapon.Type, &weapon.Weight, &weapon.Damage, &weapon.MinAttackRange, &weapon.Range, &weapon.Accuracy, &weapon.AreaCovers)
		if err != nil {
			log.Fatal("get weapon" + err.Error())
		}
	}

	return weapon
}

