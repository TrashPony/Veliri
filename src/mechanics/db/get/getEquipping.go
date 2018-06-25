package get

import (
	"log"
	"../../gameObjects/effect"
	"../../../dbConnect"
	"../../gameObjects/equip"
)

func TypeEquipping() []equip.Equip {

	rows, err := dbConnect.GetDBConnect().Query("SELECT * FROM equipping_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var equipping = make([]equip.Equip, 0)

	for rows.Next() {
		var equipType equip.Equip

		err := rows.Scan(&equipType.Id, &equipType.Type, &equipType.Specification, &equipType.Applicable, &equipType.Region)
		if err != nil {
			log.Fatal(err)
		}

		EffectsEquip(&equipType)

		equipping = append(equipping, equipType)
	}

	return equipping
}

func TypeEquip(id int) equip.Equip {

	rows, err := dbConnect.GetDBConnect().Query("SELECT * FROM equipping_type WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var equipType equip.Equip

	for rows.Next() {
		err := rows.Scan(&equipType.Id, &equipType.Type, &equipType.Specification,
			&equipType.Applicable, &equipType.Region)
		if err != nil {
			log.Fatal(err)
		}
	}

	EffectsEquip(&equipType)

	return equipType
}

func EffectsEquip(equipType *equip.Equip) {

	equipType.Effects = make([]*effect.Effect, 0)

	rows, err := dbConnect.GetDBConnect().Query(" SELECT et.id, et.name, et.level, et.type, et.steps_time, et.parameter, et.quantity, et.percentages, et.forever"+
		" FROM equip_effects, effects_type et WHERE equip_effects.id_equip=$1 AND equip_effects.id_effect=et.id;", equipType.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var equipEffect effect.Effect

		err := rows.Scan(&equipEffect.TypeID, &equipEffect.Name, &equipEffect.Level, &equipEffect.Type, &equipEffect.StepsTime, &equipEffect.Parameter,
			&equipEffect.Quantity, &equipEffect.Percentages, &equipEffect.Forever)
		if err != nil {
			log.Fatal(err)
		}

		equipType.Effects = append(equipType.Effects, &equipEffect)
	}
}
