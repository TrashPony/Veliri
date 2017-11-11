package objects

import (
	"log"
)

func GetGame(idGame int) (Game) {

	rows, err := db.Query("Select * FROM action_games WHERE id=$1", idGame)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var game Game

	for rows.Next() {
		err := rows.Scan(&game.Id, &game.Name, &game.IdMap, &game.Step, &game.Phase, &game.Winner)
		if err != nil {
			log.Fatal(err)
		}
	}

	return game
}

type Game struct {
	Id int
	Name string
	IdMap int
	Step int
	Phase string
	Winner string
}