package db

import (
	"log"
	"../unit"
	"../effect"
	"../equip"
)

func GetNewLvlEffect(oldEffect *effect.Effect, up int) *effect.Effect {
	newLevel := oldEffect.Level + up

	rows, err := db.Query("SELECT * FROM effects_type WHERE level=$1 AND name=$2 AND type=$3 AND parameter=$4;",
		newLevel, oldEffect.Name, oldEffect.Type, oldEffect.Parameter)

	if err != nil {
		println("get new lvl effect")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var newEffect effect.Effect

		err := rows.Scan(&newEffect.TypeID, &newEffect.Name, &newEffect.Level, &newEffect.Type,
			&newEffect.StepsTime, &newEffect.Parameter, &newEffect.Quantity, &newEffect.Percentages, &newEffect.Forever)
		if err != nil {
			println("get new lvl effect")
			log.Fatal(err)
		}

		newEffect.StepsTime = oldEffect.StepsTime
		newEffect.ID = oldEffect.ID

		return &newEffect
	}

	return nil
}

func GetUnitEffects(unit *unit.Unit) {

	rows, err := db.Query("SELECT age.id, et.id, et.name, et.level, et.type, age.left_steps, et.parameter, et.quantity, et.percentages, et.forever "+
		"FROM action_game_unit_effects age, effects_type et "+
		"WHERE age.id_unit=$1 AND age.id_effect=et.id;", unit.Id)
	if err != nil {
		println("get unit effects")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var unitEffect effect.Effect

		err := rows.Scan(&unitEffect.ID, &unitEffect.TypeID, &unitEffect.Name, &unitEffect.Level, &unitEffect.Type,
			&unitEffect.StepsTime, &unitEffect.Parameter, &unitEffect.Quantity, &unitEffect.Percentages, &unitEffect.Forever)
		if err != nil {
			println("get unit effects")
			log.Fatal(err)
		}

		unit.Effects = append(unit.Effects, &unitEffect)
	}
}

func GetEffectsEquip(equip *equip.Equip) {

	equip.Effects = make([]*effect.Effect, 0)

	rows, err := db.Query(" SELECT et.id, et.name, et.level, et.type, et.steps_time, et.parameter, et.quantity, " +
		" et.percentages, et.forever "+
		" FROM action_game_equipping age, equip_effects ee, effects_type et "+
		" WHERE age.id = $1 AND age.id_type = ee.id_equip AND ee.id_effect = et.id;", equip.Id)

	if err != nil {
		println("get user equip effects")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var equipEffect effect.Effect

		err := rows.Scan(&equipEffect.TypeID, &equipEffect.Name, &equipEffect.Level,&equipEffect.Type, &equipEffect.StepsTime,
			&equipEffect.Parameter, &equipEffect.Quantity, &equipEffect.Percentages, &equipEffect.Forever)
		if err != nil {
			println("get user equip effects")
			log.Fatal(err)
		}

		equip.Effects = append(equip.Effects, &equipEffect)
	}
}