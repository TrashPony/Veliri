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
		err := rows.Scan(&user.Id_game, &user.Id_user, &user.Id_resp, &user.Price, &user.Ready)
		if err != nil {
			log.Fatal(err)
		}
	}

	return user
}

type UserStat struct {
	Id_game int
	Id_user int
	Id_resp int
	Price  int
	Ready string
}