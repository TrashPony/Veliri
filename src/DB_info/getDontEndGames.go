package DB_info

import (
	"database/sql"
	"log"
	"strconv"
	_ "github.com/lib/pq"
)

func DontEndGames(userName string)(string, string)  {
	userId := strconv.Itoa(GetID("WHERE name='" + userName + "'"))

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select name, id FROM action_games WHERE id=( Select id_game From action_game_user Where id_user=" + userId + ")")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var games string = ""
	var game ActiveGames
	var ids string = ""

	for rows.Next() {
		err := rows.Scan(&game.name, &game.id)
		if err != nil {
			log.Fatal(err)
		}
		games = games + game.name  + ":"
		ids = ids + strconv.Itoa(game.id) + ":"
	}

	return games, ids
}
