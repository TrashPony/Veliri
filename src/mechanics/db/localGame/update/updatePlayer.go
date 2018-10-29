package update

import (
	"../../../../dbConnect"
	"../../../player"
	"log"
)

func Player(client *player.Player) {
	_, err := dbConnect.GetDBConnect().Exec("Update action_game_user SET ready=$1, move=$4, sub_move=$5 " +
		" WHERE id_game=$2 AND id_user=$3",
		client.GetReady(), client.GetGameID(), client.GetID(), client.Move, client.SubMove)

	if err != nil {
		println("update game user stat")
		log.Fatal(err)
	}
}
