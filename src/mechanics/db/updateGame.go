package db

import (
	"../game"
	"log"
)

func UpdateGame(game *game.Game)  {
	_, err := db.Exec("Update action_games SET phase=$1, step=$2 WHERE id_game=$3", game.Phase, game.Step, game.Id)

	if err != nil {
		println("update game")
		log.Fatal(err)
	}
}