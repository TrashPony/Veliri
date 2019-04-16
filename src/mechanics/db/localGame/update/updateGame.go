package update

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"log"
)

func Game(game *localGame.Game) {
	_, err := dbConnect.GetDBConnect().Exec("Update action_games SET phase=$1, step=$2, end_game=$4 WHERE id=$3", game.Phase, game.Step, game.Id, game.End)

	if err != nil {
		log.Fatal("update game", err)
	}

	updatePacts(game)
}

func updatePacts(game *localGame.Game) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM action_game_pacts WHERE id_game = $1", game.Id)
	if err != nil {
		log.Fatal("delete all pacts game" + err.Error())
	}

	for _, pact := range game.Pacts {
		_, err := dbConnect.GetDBConnect().Exec("INSERT INTO action_game_pacts (id_game, id_user, id_to_user) VALUES ($1, $2, $3)",
			game.Id, pact.UserID1, pact.UserID2)
		if err != nil {
			log.Fatal("add pact to game" + err.Error())
		}
	}
}
