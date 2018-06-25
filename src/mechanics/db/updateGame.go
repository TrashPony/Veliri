package db

import (
	"../game"
	"log"
	"../../dbConnect"
)

func UpdateGame(game *game.Game)  {
	_, err := dbConnect.GetDBConnect().Exec("Update action_games SET phase=$1, step=$2 WHERE id=$3", game.Phase, game.Step, game.Id)

	if err != nil {
		println("update game")
		log.Fatal(err)
	}
}