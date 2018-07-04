package updateSquad

import (
	"../../gameObjects/squad"
	"../../../dbConnect"
	"log"
)

func Squad(squad *squad.Squad) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	if err != nil {
		log.Fatal("update squad tx error: " + err.Error())
	}

	_, err = tx.Exec("UPDATE squads SET active=$1, in_game=$2 WHERE id=$3",
		squad.Active, squad.InGame, squad.ID)

	if err != nil {
		log.Fatal("update squad" + err.Error())
	}

	InventorySquad(squad, tx)
	MotherShip(squad, tx)
	Units(squad, tx)

	tx.Commit()
}