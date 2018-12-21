package update

import (
	"../../../../dbConnect"
	"../../../gameObjects/squad"
	"log"
)

func Squad(squad *squad.Squad, full bool) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	if err != nil {
		log.Fatal("update squad tx error: " + err.Error())
	}

	_, err = tx.Exec("UPDATE squads SET active=$1, in_game=$2, q=$4, r=$5, id_map=$6 WHERE id=$3",
		squad.Active, squad.InGame, squad.ID, squad.Q, squad.R, squad.MapID)

	if err != nil {
		log.Fatal("update squad" + err.Error())
	}

	if full {
		InventorySquad(squad, tx)
		Units(squad, tx)
	}
	MotherShip(squad, tx)

	tx.Commit()
}
