package mechanics

/*func PhaseСhange(game *Game) (string, error) {

	rows, err := db.Query("Select phase, step from action_games WHERE id=$1", game.GetStat().Id)
	if err != nil {
		println("123")
	}
	defer rows.Close()

	_, err = db.Query("UPDATE action_game_unit  SET action = $1 WHERE id_game=$2", true, game.GetStat().Id)
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
		// ставит новую фазу
		_, err := db.Query("UPDATE action_games SET phase=$2, step=$3 WHERE id=$1", game.GetStat().Id, "move", step+1)
		if err != nil {
			return "", err
		}
		// сбрасывает всем пользователям готовность в игре
		_, errr := db.Query("UPDATE action_game_user SET ready=$2 WHERE id_game=$1", game.GetStat().Id, false)
		if errr != nil {
			return "", errr
		} else {
			phase = "move"
		}
	} else {
		if phase == "move" {
			_, err := db.Query("UPDATE action_games SET phase=$2 WHERE id=$1", game.GetStat().Id, "targeting")
			if err != nil {
				return "", err
			}
			_, errr := db.Query("UPDATE action_game_user SET ready=$2 WHERE id_game=$1", game.GetStat().Id, false)
			if errr != nil {
				return "", errr
			} else {
				phase = "targeting"
			}
		} else {
			if phase == "targeting" {
				_, err := db.Query("UPDATE action_games SET phase=$2 WHERE id=$1", game.GetStat().Id, "attack")
				if err != nil {
					return "", err
				}
				_, errr := db.Query("UPDATE action_game_user SET ready=$2 WHERE id_game=$1", game.GetStat().Id, false)
				if errr != nil {
					return "", errr
				} else {
					phase = "attack"
				}
			}
		}
	}
	return phase, nil
}*/
