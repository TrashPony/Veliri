package mechanics

/*func UserReady(client *player.Player, game *game.Game) (string, error, bool) {

	// устанавливает фраг готовности пользователя в тру
	rows, err := db.Query("UPDATE action_game_user  SET ready = true WHERE id_user=$1 AND id_game=$2", client.GetID(), game.Id)
	if err != nil {
		return "", err, false
	}
	// берем готовность всех пользователей
	rows, err = db.Query("Select ready FROM action_game_user WHERE id_game=$1", game.Id)
	if err != nil {
		return "", err, false
	}
	defer rows.Close()

	var ready = make([]bool, 0)
	var user bool

	for rows.Next() {
		err := rows.Scan(&user)
		if err != nil {
			return "", err, false
		}
		ready = append(ready, user)
	}

	// обновляем статус игрока в памяти
	client.SetReady(true)

	var allReady bool
	for i := 0; i < len(ready); i++ {
		if i == 0 {
			allReady = ready[0]
		}
		allReady = allReady && ready[i]
	}

	var phase string
	// если все игроки готовы то начинается смена фазы
	if allReady {
		phase, err = PhaseСhange(game)
		if err != nil {
			return "", err, false
		} else {
			return phase, nil, true
		}
	} else {
		return phase, nil, false
	}
}*/