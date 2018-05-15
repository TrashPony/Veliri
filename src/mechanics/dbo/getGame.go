package dbo

import (
	"../game"
	"log"
)

func GetGame(idGame int) *game.Game {

	rows, err := db.Query("Select id, name, id_map, step, phase, winner FROM action_games WHERE id=$1", idGame)
	if err != nil {
		println("Error GetInfo Game")
		log.Fatal(err)
	}
	defer rows.Close()

	var loadGame game.Game

	for rows.Next() {
		err := rows.Scan(&loadGame.Id, &loadGame.Name, &loadGame.MapID, &loadGame.Step, &loadGame.Phase, &loadGame.Winner)
		if err != nil {
			println("Error GetInfo Game")
			log.Fatal(err)
		}
	}

	return &loadGame
}
