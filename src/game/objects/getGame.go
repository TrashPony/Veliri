package objects

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetGame(idGame int) (Game) {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}
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