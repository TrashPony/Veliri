package update

import (
	"../../../../dbConnect"
	"../../../player"
	"encoding/json"
	"log"
)

func Player(client *player.Player) {
	_, err := dbConnect.GetDBConnect().Exec("Update action_game_user SET ready=$1"+
		" WHERE id_game=$2 AND id_user=$3",
		client.GetReady(), client.GetGameID(), client.GetID())

	if err != nil {
		println("update game user stat")
		log.Fatal(err)
	}

	memoryPlayer(client)
}

func memoryPlayer(client *player.Player) {
	for _, memoryUnit := range client.GetMemoryHostileUnits() {

		var find bool

		rows, err := dbConnect.GetDBConnect().Query("SELECT EXISTS("+
			"SELECT id FROM user_memory_unit WHERE id_user = $1 AND id_game = $2 AND id_unit = $3"+
			")", client.GetID(), client.GetGameID(), memoryUnit.ID)
		if err != nil {
			println("get memory unit")
			log.Fatal(err)
		}

		for rows.Next() {
			err := rows.Scan(&find)
			if err != nil {
				println("scan get memory unit")
				log.Fatal(err)
			}
		}
		rows.Close()

		jsonUnit, err := json.Marshal(memoryUnit)
		if err != nil {
			println("unit to json")
			log.Fatal(err)
		}

		if find {
			_, err := dbConnect.GetDBConnect().Exec("Update user_memory_unit "+
				" SET unit = $1 "+
				" WHERE id_game=$2 AND id_user=$3 AND id_unit = $4",
				jsonUnit, client.GetGameID(), client.GetID(), memoryUnit.ID)
			if err != nil {
				println("update memory unit")
				log.Fatal(err)
			}
		} else {
			_, err = dbConnect.GetDBConnect().Exec("INSERT INTO user_memory_unit (unit, id_user, id_game, id_unit) VALUES ($1, $2, $3, $4)",
				jsonUnit, client.GetID(), client.GetGameID(), memoryUnit.ID)
			if err != nil {
				println("add memory unit")
				log.Fatal(err)
			}
		}
	}
}
