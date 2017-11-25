package objects

import (
	"log"
)

func GetUserStat(idGame int) []*UserStat {

	rows, err := db.Query("Select agu.id_game, users.name, agu.start_structure, agu.price, agu.ready FROM action_game_user as agu, users WHERE agu.id_user=users.id AND id_game=$1", idGame)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*UserStat, 0)
	for rows.Next() {
		var user UserStat
		err := rows.Scan(&user.IdGame, &user.Name, &user.IdResp, &user.Price, &user.Ready)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}

	return users
}
