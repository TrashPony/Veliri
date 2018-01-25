package lobby

import "log"

type Weapon struct {
	Id          int    `json:"id"`
	Type        string `json:"type"`
	Damage      int    `json:"damage"`
	RangeAttack int    `json:"range_attack"`
	RangeView   int    `json:"range_view"`
	AreaAttack  int    `json:"area_attack"`
	TypeAttack  string `json:"type_attack"`
	Size        int    `json:"size"`
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
		err := rows.Scan(&weapon.Id, &weapon.Type, &weapon.Damage, &weapon.RangeAttack, &weapon.RangeView, &weapon.AreaAttack, &weapon.TypeAttack, &weapon.Size)
		if err != nil {
			log.Fatal(err)
		}
		weapons = append(weapons, weapon)
	}

	return weapons
}
