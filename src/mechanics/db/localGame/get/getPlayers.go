package get

import (
	"../../../localGame"
	"../../../player"
	"../../../players"
	"../../../../dbConnect"
	"log"
)

func Players(game *localGame.Game) []*player.Player {

	rows, err := dbConnect.GetDBConnect().Query("Select users.name, agu.ready, users.id "+
		"FROM action_game_user as agu, users "+
		"WHERE agu.id_user=users.id AND agu.id_game=$1", game.Id)
	if err != nil {
		println("get game user stat")
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*player.Player, 0)

	for rows.Next() {
		var client *player.Player

		var login string
		var ready bool
		var id int

		err := rows.Scan(&login, &ready, &id)
		if err != nil {
			log.Fatal(err)
		}

		client, ok := players.Users.Get(id)

		if !ok {
			client = players.Users.Add(id, login)
		}

		client.SetReady(ready)
		client.SetGameID(game.Id)

		// todo watchZone.UpdateWatchZone(game, &client)

		users = append(users, client)
	}

	return users
}