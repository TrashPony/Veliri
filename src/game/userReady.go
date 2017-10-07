package game

import (
	"database/sql"
	"log"
)

func UserReady(idUser int, idGame string) (string)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}
	// устанавливает фраг готовности пользователя в тру
	rows, err := db.Query("UPDATE action_game_user  SET ready = true WHERE id_user=$1 AND id_game=$2", idUser, idGame)
	if err != nil {
		log.Fatal(err)
	}
	// берем готовность всех пользователей
	rows, err = db.Query("Select ready FROM action_game_user WHERE id_game=$1", idGame)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ready = make([]bool, 0)
	var user bool

	for rows.Next() {
		err := rows.Scan(&user)
		if err != nil {
			log.Fatal(err)
		}
		ready = append(ready, user)
	}

	var allReady bool
	for i := 0; i < len(ready); i++ {
		if i == 0 { allReady = ready[0] }
		allReady = allReady && ready[i]
	}

	var phase string
	// если все игроки готовы то начинается смена фазы
	if allReady {
		phase = PhaseСhange()
		return phase
	} else {
		return phase
	}
}
