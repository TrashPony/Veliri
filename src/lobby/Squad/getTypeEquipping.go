package Squad

import "log"

func GetTypeEquipping() []Equipping {

	rows, err := db.Query("SELECT * FROM equipping_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var equipping = make([]Equipping, 0)
	var equip Equipping

	for rows.Next() {
		err := rows.Scan(&equip.Id, &equip.Type, &equip.Specification)
		if err != nil {
			log.Fatal(err)
		}
		equipping = append(equipping, equip)
	}

	return equipping
}

func GetTypeEquip(id int) Equipping {

	rows, err := db.Query("SELECT * FROM equipping_type WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var equip Equipping

	for rows.Next() {
		err := rows.Scan(&equip.Id, &equip.Type, &equip.Specification)
		if err != nil {
			log.Fatal(err)
		}
	}

	return equip
}

type Equipping struct {
	Id            int    `json:"id"`
	Type          string `json:"type"`
	Specification string `json:"specification"`
}
