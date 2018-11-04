package get

import (
	"../../../dbConnect"
	"../../gameObjects/effect"
	"../../gameObjects/equip"
	"log"
)

func TypeEquip(id int) *equip.Equip {

	rows, err := dbConnect.GetDBConnect().Query("SELECT id,"+
		" name,"+
		" active,"+
		" specification,"+
		" applicable,"+
		" region,"+
		" radius,"+
		" type_slot,"+
		" reload,"+
		" power,"+
		" use_power,"+
		" max_hp,"+
		" steps_time,"+
		" size, "+
		" initiative "+
		""+
		"FROM equipping_type "+
		"WHERE id=$1", id)
	if err != nil {
		log.Fatal("get type equip " + err.Error())
	}
	defer rows.Close()

	var equipType equip.Equip

	for rows.Next() {
		err := rows.Scan(&equipType.ID, &equipType.Name, &equipType.Active, &equipType.Specification,
			&equipType.Applicable, &equipType.Region, &equipType.Radius, &equipType.TypeSlot, &equipType.Reload,
			&equipType.Power, &equipType.UsePower, &equipType.MaxHP, &equipType.StepsTime, &equipType.Size,
			&equipType.Initiative)
		if err != nil {
			log.Fatal("get scan type equip " + err.Error())
		}
	}

	EffectsEquip(&equipType)

	return &equipType
}

func EffectsEquip(equipType *equip.Equip) {

	equipType.Effects = make([]*effect.Effect, 0)

	rows, err := dbConnect.GetDBConnect().Query(" SELECT et.id, et.name, et.level, et.type, et.parameter, et.quantity, et.percentages, et.forever"+
		" FROM equip_effects, effects_type et WHERE equip_effects.id_equip=$1 AND equip_effects.id_effect=et.id;", equipType.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var equipEffect effect.Effect

		err := rows.Scan(&equipEffect.TypeID, &equipEffect.Name, &equipEffect.Level, &equipEffect.Type, &equipEffect.Parameter,
			&equipEffect.Quantity, &equipEffect.Percentages, &equipEffect.Forever)
		if err != nil {
			log.Fatal(err)
		}

		equipType.Effects = append(equipType.Effects, &equipEffect)
	}
}
