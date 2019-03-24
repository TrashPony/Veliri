package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"log"
)

func Game(playerID int) *localGame.Game {
	// игрок может быть одновременно только в 1 битве, поэтому находим ид по таблице action_game_user
	rowsID, err := dbConnect.GetDBConnect().Query("Select id_game FROM action_game_user WHERE id_user=$1", playerID)
	if err != nil {
		log.Fatal(err, "Error GetInfo Game")
	}
	defer rowsID.Close()

	idGame := 0
	for rowsID.Next() {
		err := rowsID.Scan(&idGame)
		if err != nil {
			log.Fatal(err, "get game by user id")
		}
	}

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, id_map, step, phase, winner FROM action_games WHERE id=$1", idGame)
	if err != nil {
		log.Fatal(err, "Error GetInfo Game")
	}
	defer rows.Close()

	var loadGame localGame.Game

	for rows.Next() {
		err := rows.Scan(&loadGame.Id, &loadGame.Name, &loadGame.MapID, &loadGame.Step, &loadGame.Phase, &loadGame.Winner)
		if err != nil {
			log.Fatal(err, "Error GetInfo Game")
		}
	}

	return &loadGame
}
