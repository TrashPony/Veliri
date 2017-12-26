package game

import (
	"log"
)

type UserStat struct {
	IdGame int
	Name   string
	Price  int
	Ready  bool
	RespX int
	RespY int
}


func GetUserStat(idGame int) []*UserStat {

	rows, err := db.Query("Select agu.id_game, users.name, agu.price, agu.ready, ags.x, ags.y FROM action_game_user as agu, action_game_structure as ags, users WHERE ags.id=agu.start_structure AND agu.id_user=users.id AND agu.id_game=$1", idGame)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]*UserStat, 0)
	for rows.Next() {
		var user UserStat
		err := rows.Scan(&user.IdGame, &user.Name, &user.Price, &user.Ready, &user.RespX, &user.RespY)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}

	return users
}
