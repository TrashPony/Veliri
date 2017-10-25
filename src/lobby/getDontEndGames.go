package lobby

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func GetDontEndGames(userName string)([]DontEndGames)  {
	user := GetUsers("WHERE name='" + userName + "'")

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select id, name, id_map, step, phase, winner, ready FROM action_games, action_game_user WHERE action_game_user.id_game=action_games.id AND action_game_user.id_user=$1", user.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var games = make([]DontEndGames, 0)
	var game DontEndGames

	for rows.Next() {
		err := rows.Scan(&game.Id, &game.Name, &game.IdMap, &game.Step, &game.Phase, &game.Winner, &game.Ready)
		if err != nil {
			log.Fatal(err)
		}
		games = append(games, game)
	}

	return games
}
