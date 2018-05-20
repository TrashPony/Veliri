package db

import (
	"log"
	"../player"
)

func UpdatePlayer(client *player.Player) { // впринципе тут больше нечего нельзя обновить :)
	_, err := db.Exec("Update action_game_user SET ready=$1 WHERE id_game=$2", client.GetReady(), client.GetGameID())

	if err != nil {
		println("update game user stat")
		log.Fatal(err)
	}
}
