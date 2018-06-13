package Squad

import (
	"log"
	"../../mechanics/effect"
)

func GetTypeEquipping() []Equipping {

	rows, err := db.Query("SELECT * FROM equipping_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var equipping = make([]Equipping, 0)

	for rows.Next() {
		var equip Equipping

		err := rows.Scan(&equip.Id, &equip.Type, &equip.Specification, &equip.Applicable, &equip.Region)
		if err != nil {
			log.Fatal(err)
		}

		GetEffectsEquip(&equip)

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
		err := rows.Scan(&equip.Id, &equip.Type, &equip.Specification, &equip.Applicable, &equip.Region)
		if err != nil {
			log.Fatal(err)
		}
	}

	GetEffectsEquip(&equip)

	return equip
}

func GetEffectsEquip(equip *Equipping) {

	equip.Effects = make([]effect.Effect, 0)

	rows, err := db.Query(" SELECT et.id, et.name, et.type, et.steps_time, et.parameter, et.quantity, et.percentages, et.forever"+
		" FROM equip_effects, effects_type et WHERE equip_effects.id_equip=$1 AND equip_effects.id_effect=et.id;", equip.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var equipEffect effect.Effect

		err := rows.Scan(&equipEffect.TypeID, &equipEffect.Name, &equipEffect.Type, &equipEffect.StepsTime, &equipEffect.Parameter,
			&equipEffect.Quantity, &equipEffect.Percentages, &equipEffect.Forever)
		if err != nil {
			log.Fatal(err)
		}

		equip.Effects = append(equip.Effects, equipEffect)
	}
}

type Equipping struct {
	Id            int             `json:"id"`
	Type          string          `json:"type"`
	Specification string          `json:"specification"`
	Effects       []effect.Effect `json:"effects"`
	Applicable    string          `json:"applicable"`
	Region        int             `json:"region"`
}
