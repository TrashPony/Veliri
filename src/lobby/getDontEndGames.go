package lobby

import (
	"log"
)

func GetDontEndGames(userName string)([]DontEndGames)  {
	user := GetUsers("WHERE name='" + userName + "'")

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
