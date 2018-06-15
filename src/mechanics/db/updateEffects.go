package db

import (
	"../unit"
	"../coordinate"
	"log"
)

func UpdateUnitEffects(unit *unit.Unit) error {
	for _, unitEffect := range unit.Effects {
		if unitEffect.StepsTime == 0 {

			_, err := db.Exec("DELETE FROM action_game_unit_effects WHERE id=$1", unitEffect.ID)

			if err != nil {
				println("Error delete unit effect")
				log.Fatal(err)
				return err
			}

		} else {
			if unitEffect.ID != 0 {

				_, err := db.Exec("UPDATE action_game_unit_effects SET left_steps=$1, id_effect=$3 WHERE id=$2", unitEffect.StepsTime, unitEffect.ID, unitEffect.TypeID)

				if err != nil {
					println("Error update unit effect")
					log.Fatal(err)
					return err
				}

			} else {
				id := 0
				err := db.QueryRow("INSERT INTO action_game_unit_effects (id_unit, id_effect, left_steps) "+
					"VALUES ($1, $2, $3) RETURNING id", unit.Id, unitEffect.TypeID, unitEffect.StepsTime).Scan(&id)

				if err != nil {
					println("Error add new unit effect")
					log.Fatal(err)
					return err
				}

				unitEffect.ID = id
			}
		}
	}

	return nil
}

func UpdateCoordinateEffects(mapCoordinate *coordinate.Coordinate) error {
	for _, coordinateEffect := range mapCoordinate.Effects {
		if coordinateEffect.StepsTime == 0 {

			_, err := db.Exec("DELETE FROM action_game_zone_effects WHERE id=$1", coordinateEffect.ID)

			if err != nil {
				println("Error delete coordinate effect")
				log.Fatal(err)
				return err
			}

		} else {
			if coordinateEffect.ID != 0 {

				_, err := db.Exec("UPDATE action_game_zone_effects SET left_steps=$1, x=$2, y=$3 id_effect=$4 WHERE id=$5",
					coordinateEffect.StepsTime, mapCoordinate.X, mapCoordinate.Y, coordinateEffect.TypeID, coordinateEffect.ID)

				if err != nil {
					println("Error update coordinate effect")
					log.Fatal(err)
					return err
				}

			} else {

				id := 0
				err := db.QueryRow("INSERT INTO action_game_zone_effects (id_game, id_effect, x, y, left_steps) "+
					"VALUES ($1, $2, $3, $4, $5) RETURNING id", mapCoordinate.GameID, coordinateEffect.TypeID, mapCoordinate.X, mapCoordinate.Y, coordinateEffect.StepsTime).Scan(&id)

				if err != nil {
					println("Error add new coordinate effect")
					log.Fatal(err)
					return err
				}

				coordinateEffect.ID = id
			}
		}
	}

	return nil
}