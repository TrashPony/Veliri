package get

import (
	"../../../../dbConnect"
	"../../../localGame"
	"../../../player"
	"../../../players"
	"log"
)

func Players(game *localGame.Game) []*player.Player {

	rows, err := dbConnect.GetDBConnect().Query("Select users.name, " +
		"agu.ready, " +
		"users.id, " +
		"agu.move, " +
		"agu.sub_move, " +
		"queue_move_pos " +
		""+
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
		var move bool
		var subMove bool
		var queueMovePos int

		err := rows.Scan(&login, &ready, &id, &move, &subMove, &queueMovePos)
		if err != nil {
			log.Fatal(err)
		}

		client, ok := players.Users.Get(id)

		if !ok {
			client = players.Users.Add(id, login)
		}

		client.SetReady(ready)
		client.SetGameID(game.Id)
		client.Move = move
		client.SubMove = subMove
		client.QueueMovePos = queueMovePos

		users = append(users, client)
	}

	return users
}
