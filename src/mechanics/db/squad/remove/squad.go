package remove

import (
	"../../../../dbConnect"
	"../../../gameObjects/squad"
	"log"
)

func Squad(squad *squad.Squad) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM squad_thorium_slots WHERE id_squad = $1", squad.ID)
	if err != nil {
		log.Fatal("delete squad, thorium" + err.Error())
	}

	_, err = tx.Exec("DELETE FROM squad_inventory WHERE id_squad = $1", squad.ID)
	if err != nil {
		log.Fatal("delete squad, inventory" + err.Error())
	}

	_, err = tx.Exec("DELETE FROM squad_units_equipping WHERE id_squad = $1", squad.ID)
	if err != nil {
		log.Fatal("delete squad, squad_units_equipping" + err.Error())
	}

	_, err = tx.Exec("DELETE FROM squad_units WHERE id_squad = $1", squad.ID)
	if err != nil {
		log.Fatal("delete squad, squad_units" + err.Error())
	}

	_, err = tx.Exec("DELETE FROM squads WHERE id = $1", squad.ID)
	if err != nil {
		log.Fatal("delete squad" + err.Error())
	}

	tx.Commit()
}
