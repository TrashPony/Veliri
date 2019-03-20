package get

import (
	"encoding/json"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
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
		println("get local game user stat")
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*player.Player, 0)

	for rows.Next() {
		var login string
		var ready bool
		var id int

		err := rows.Scan(&login, &ready, &id)
		if err != nil {
			log.Fatal(err, "scan local game user stat")
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

	for rows.Next() {
		var jsonUnit []byte
		var memoryUnit unit.Unit

		rows.Scan(&jsonUnit)
		json.Unmarshal(jsonUnit, &memoryUnit)

		rows, err := dbConnect.GetDBConnect().Query(
			"SELECT move, action_point FROM squad_units WHERE id=$1", memoryUnit.ID)
		if err != nil {
			println("get move params memory unit")
			log.Fatal(err)
		}

		for rows.Next() {
			rows.Scan(&memoryUnit.Move, &memoryUnit.ActionPoints)
		}

		client.AddNewMemoryHostileUnit(memoryUnit)
	}
}
