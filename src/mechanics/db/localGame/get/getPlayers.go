package get

import (
	"../../../../dbConnect"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../player"
	"../../../players"
	"encoding/json"
	"log"
)

func Players(game *localGame.Game) []*player.Player {

	rows, err := dbConnect.GetDBConnect().Query("Select users.name, "+
		"agu.ready, "+
		"users.id "+
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

		users = append(users, client)
		memoryPlayer(client)
	}

	return users
}

func memoryPlayer(client *player.Player) {
	rows, err := dbConnect.GetDBConnect().Query(
		"SELECT unit FROM user_memory_unit WHERE id_user = $1 AND id_game = $2", client.GetID(), client.GetGameID())
	if err != nil {
		println("get user memory unit")
		log.Fatal(err)
	}
	defer rows.Close()

	var jsonUnit []byte
	var memoryUnit unit.Unit

	for rows.Next() {
		rows.Scan(&jsonUnit)
	}

	json.Unmarshal(jsonUnit, &memoryUnit)
	client.AddNewMemoryHostileUnit(memoryUnit)
}
