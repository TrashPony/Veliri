package game

import (
	"log"
)

type UserStat struct {
	IdGame int
	Name   string
	Ready  bool
}

func GetUserStat(idGame int) []*UserStat {

	rows, err := db.Query("Select agu.id_game, users.name, agu.ready FROM action_game_user as agu, users WHERE agu.id_user=users.id AND agu.id_game=$1", idGame)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*UserStat, 0)
	for rows.Next() {
		var user UserStat
		err := rows.Scan(&user.IdGame, &user.Name, &user.Ready)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}

	return users
}
