package game

func PhaseСhange(idGame int) (string, error) {

	rows, err := db.Query("Select phase, step from action_games WHERE id=$1", idGame)
	if err != nil {
		println("123")
	}
	defer rows.Close()

	_, err = db.Query("UPDATE action_game_unit  SET action = $1 WHERE id_game=$2", true, idGame)
	if err != nil {

	}
	var phase string
	var step int
	for rows.Next() {
		err := rows.Scan(&phase, &step)
		if err != nil {
			return "", err
		}
	}

	if phase == "attack" || phase == "Init" {
		_, err := db.Query("UPDATE action_games SET phase=$2, step=$3 WHERE id=$1", idGame, "move", step+1)
		if err != nil { // TODO : зарефакторить
			return "", err
		}
		_, errr := db.Query("UPDATE action_game_user SET ready=$2 WHERE id_game=$1", idGame, false)
		if errr != nil {
			return "", errr
		} else {
			phase = "move"
		}
	} else {
		if phase == "move" {
			_, err := db.Query("UPDATE action_games SET phase=$2 WHERE id=$1", idGame, "targeting")
			if err != nil { // TODO : зарефакторить
				return "", err
			}
			_, errr := db.Query("UPDATE action_game_user SET ready=$2 WHERE id_game=$1", idGame, false)
			if errr != nil {
				return "", errr
			} else {
				phase = "targeting"
			}
		} else {
			if phase == "targeting" {
				_, err := db.Query("UPDATE action_games SET phase=$2 WHERE id=$1", idGame, "attack")
				if err != nil { // TODO : зарефакторить
					return "", err
				}
				_, errr := db.Query("UPDATE action_game_user SET ready=$2 WHERE id_game=$1", idGame, false)
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
