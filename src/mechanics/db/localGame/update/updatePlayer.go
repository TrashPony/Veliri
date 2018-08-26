package update

import (
	"log"
	"../../../player"
	"../../../../dbConnect"
)

func Player(client *player.Player) {
	_, err := dbConnect.GetDBConnect().Exec("Update action_game_user SET ready=$1 WHERE id_game=$2 AND id_user=$3",
		client.GetReady(), client.GetGameID(), client.GetID())

	if err != nil {
		println("update game user stat")
		log.Fatal(err)
	}
}