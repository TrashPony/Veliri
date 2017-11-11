package objects

import (
	"log"
)

func GetUserStat(idGame int) ([]UserStat)  {

	rows, err := db.Query("Select agu.id_game, users.name, agu.respawns_id, agu.price, agu.ready FROM action_game_user as agu, users WHERE agu.id_user=users.id AND id_game=$1", idGame)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]UserStat,0)
	var user UserStat
	for rows.Next() {
		err := rows.Scan(&user.IdGame, &user.Name, &user.IdResp, &user.Price, &user.Ready)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}

