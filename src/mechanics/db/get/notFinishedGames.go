package get

import (
	"../../../dbConnect"
	"../../lobby/notFinishedGames"
	"log"
)

func GetNotFinishedGames(userID int) []notFinishedGames.NotFinishedGames {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, id_map, step, phase, winner, ready FROM action_games, action_game_user WHERE action_game_user.id_game=action_games.id AND action_game_user.id_user=$1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var games = make([]notFinishedGames.NotFinishedGames, 0)

	for rows.Next() {
		var game notFinishedGames.NotFinishedGames
		var mapID int

		err := rows.Scan(&game.Id, &game.Name, &mapID, &game.Step, &game.Phase, &game.Winner, &game.Ready)
		if err != nil {
			log.Fatal(err)
		}

		//mp := GetMap(mapID)
		//game.Map = mp

		games = append(games, game)
	}

	return games
}
