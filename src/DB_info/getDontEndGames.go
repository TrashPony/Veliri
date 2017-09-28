package DB_info

import (
	"database/sql"
	"log"
	"strconv"
	_ "github.com/lib/pq"
)

func DontEndGames(UserName string)(string)  {
	var users = GetUsers()
	var playerId int = 0
	for _, user := range users {
		if user.name == UserName {
			playerId = user.id
		}
	}

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select name FROM activegame WHERE idplayer1=" + strconv.Itoa(playerId) + " OR idplayer2=" + strconv.Itoa(playerId))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var games string = ""
	var game ActiveGames

	for rows.Next() {
		err := rows.Scan(&game.name)
		if err != nil {
			log.Fatal(err)
		}
		games = games + game.name  + ":"
	}

	return games
}
