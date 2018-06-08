package db

import (
	"log"
	"../player"
)

func UpdatePlayer(client *player.Player) {
	_, err := db.Exec("Update action_game_user SET ready=$1 WHERE id_game=$2", client.GetReady(), client.GetGameID())

	if err != nil {
		println("update game user stat")
		log.Fatal(err)
	}

	UpdatePlayerEquip(client)
}

func UpdatePlayerEquip(client *player.Player)  {

	for _, playerEquip := range client.GetEquip() {

		_, err := db.Exec("Update action_game_equipping SET used=$1 WHERE id=$2 AND id_game=$3 AND id_user=$4",
			playerEquip.Used, playerEquip.Id, client.GetGameID(), client.GetID())

		if err != nil {
			println("update game equip")
			log.Fatal(err)
		}
	}
}
