package updateSquad

import (
	"../../gameObjects/squad"
	"../../../dbConnect"
	"log"
)

func Squad(squad *squad.Squad) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE squads SET active=$1, in_game=$2 WHERE id=$3",
		squad.Active, squad.InGame, squad.ID)

	if err != nil {
		log.Fatal("update squad" + err.Error())
	}

	InventorySquad(squad)
	MotherShip(squad)
	Units(squad)
}