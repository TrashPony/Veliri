package mechanics

import (
	"database/sql"
)

func PhaseСhange(idGame int)(string, error)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		return "", err
	}

	rows, err := db.Query("Select phase from action_games WHERE id=$1", idGame)
	if err != nil {
		println("123")
	}
	defer rows.Close()

	var phase string

	for rows.Next() {
		err := rows.Scan(&phase)
		if err != nil {
			return "", err
		}
	}

	if phase == "attack" || phase == "Init" {
		_, err := db.Query("UPDATE action_games SET phase=$2 WHERE id=$1",idGame, "move")
		if err != nil { // TODO : зарефакторить
			return "", err
		}
		_, errr := db.Query("UPDATE action_game_user SET ready=$2 WHERE id_game=$1",idGame, false)
		if errr != nil {
			return "", errr
		} else {
			phase = "move"
		}
	} else {
		if phase == "move" {
			_, err := db.Query("UPDATE action_games SET phase=$2 WHERE id=$1",idGame, "targeting")
			if err != nil { // TODO : зарефакторить
				return "", err
			}
			_, errr := db.Query("UPDATE action_game_user SET ready=$2 WHERE id_game=$1",idGame, false)
			if errr != nil {
				return "", errr
			} else {
				phase = "targeting"
			}
		} else {
			if phase == "targeting" {
				_, err := db.Query("UPDATE action_games SET phase=$2 WHERE id=$1",idGame, "attack")
				if err != nil { // TODO : зарефакторить
					return "", err
				}
				_, errr := db.Query("UPDATE action_game_user SET ready=$2 WHERE id_game=$1",idGame, false)
				if errr != nil {
					return "", errr
				} else {
					phase = "attack"
				}
			}
		}
	}
	return phase, nil
}