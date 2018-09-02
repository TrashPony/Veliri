package get

import (
	"../../../gameObjects/unit"
	"../../../gameObjects/effect"
	"../../../gameObjects/coordinate"
	"../../../../dbConnect"
	"log"
)

func MaxLvlEffect(gameEffect *effect.Effect) int {
	var maxLvl int

	row := dbConnect.GetDBConnect().QueryRow("SELECT COUNT(*) FROM effects_type WHERE name = $1 AND type = $2;", gameEffect.Name, gameEffect.Type)

	err := row.Scan(&maxLvl)

	if err != nil {
		println("get max lvl effect")
		log.Fatal(err)
	}
	return maxLvl
}

func NewLvlEffect(oldEffect *effect.Effect, up int) *effect.Effect {
	newLevel := oldEffect.Level + up

	rows, err := dbConnect.GetDBConnect().Query("SELECT * FROM effects_type WHERE level=$1 AND name=$2 AND type=$3 AND parameter=$4;",
		newLevel, oldEffect.Name, oldEffect.Type, oldEffect.Parameter)

	if err != nil {
		println("get new lvl effect")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var newEffect effect.Effect

		err := rows.Scan(&newEffect.TypeID, &newEffect.Name, &newEffect.Level, &newEffect.Type,
			&newEffect.Parameter, &newEffect.Quantity, &newEffect.Percentages, &newEffect.Forever)
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

func UnitEffects(unit *unit.Unit) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT age.id, et.id, et.name, et.level, et.type, age.left_steps, et.parameter, et.quantity, et.percentages, et.forever "+
		"FROM action_game_unit_effects age, effects_type et "+
		"WHERE age.id_unit=$1 AND age.id_effect=et.id;", unit.ID)
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

func CoordinateEffects(mapCoordinate *coordinate.Coordinate) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT age.id, et.id, et.name, et.level, et.type, age.left_steps, et.parameter, et.quantity, et.percentages, et.forever "+
		"FROM action_game_zone_effects age, effects_type et "+
		"WHERE age.id_game = $1 AND age.q = $2 AND age.r = $3 AND et.id = age.id_effect;", mapCoordinate.GameID, mapCoordinate.Q, mapCoordinate.R)
	if err != nil {
		println("get coordinate effects")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var coordinateEffect effect.Effect

		err := rows.Scan(&coordinateEffect.ID, &coordinateEffect.TypeID, &coordinateEffect.Name, &coordinateEffect.Level, &coordinateEffect.Type,
			&coordinateEffect.StepsTime, &coordinateEffect.Parameter, &coordinateEffect.Quantity, &coordinateEffect.Percentages, &coordinateEffect.Forever)
		if err != nil {
			println("get coordinate effects")
			log.Fatal(err)
		}

		mapCoordinate.Effects = append(mapCoordinate.Effects, &coordinateEffect)
	}
}