package get

import (
	"../../../localGame"
	"../../../../dbConnect"
	"log"
)

func Game(idGame int) *localGame.Game {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, id_map, step, phase, winner FROM action_games WHERE id=$1", idGame)
	if err != nil {
		println("Error GetInfo Game")
		log.Fatal(err)
	}
	defer rows.Close()

	var loadGame localGame.Game

	for rows.Next() {
		err := rows.Scan(&loadGame.Id, &loadGame.Name, &loadGame.MapID, &loadGame.Step, &loadGame.Phase, &loadGame.Winner)
		if err != nil {
			println("Error GetInfo Game")
			log.Fatal(err)
		}
	}

	return &loadGame
}
