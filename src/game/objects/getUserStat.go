package objects

import (
	"database/sql"
	"log"
)

func GetUserStat(idGame string) ([]UserStat)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

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

