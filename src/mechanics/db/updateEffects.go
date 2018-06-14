package db

import (
	"../unit"
	"log"
)

func UpdateUnitEffects(unit *unit.Unit) error {
	for _, unitEffect := range unit.Effects {
		if unitEffect.StepsTime == 0 {

			_, err := db.Exec("DELETE FROM action_game_unit_effects WHERE id=$1", unitEffect.ID)
			// todo
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
