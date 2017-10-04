package initGame

import (
	"database/sql"
	"log"
	"strconv"
)

func GetUserStat(idGame string, idUser int) (UserStat)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * FROM action_game_user WHERE id_game=" + idGame + " AND id_user=" + strconv.Itoa(idUser))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user UserStat

	for rows.Next() {
		err := rows.Scan(&user.id_game, &user.id_user, &user.price)
		if err != nil {
			log.Fatal(err)
		}
	}

	return user
}

type UserStat struct {
	id_game int
	id_user int
	price  int
}