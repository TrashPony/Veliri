package update

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"log"
)

func Game(game *localGame.Game) {
	_, err := dbConnect.GetDBConnect().Exec("Update action_games SET phase=$1, step=$2 WHERE id=$3", game.Phase, game.Step, game.Id)

	if err != nil {
		println("update game")
		log.Fatal(err)
	}
}
