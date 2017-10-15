package initGame

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetGame(idGame string) (Game) {
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
		err := rows.Scan(&game.id, &game.name, &game.idMap, &game.step, &game.phase, &game.winner)
		if err != nil {
			log.Fatal(err)
		}
	}

	return game
}

type Game struct {
	id int
	name string
	idMap int
	step int
	phase string
	winner string
}