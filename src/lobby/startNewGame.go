package lobby

import (
	"log"
)

func StartNewGame(nameGame string) (string, bool) {
	for game := range openGames {
		if openGames[game].Name == nameGame && len(openGames[game].Users) > 1 {
			id := InitNewGame(openGames[game].Map, openGames[game])
			if id != "" {
				return id, true
			} else {
				return "", false
			}
		}
	}
	return "", false
}

func InitNewGame(mapName string, game LobbyGames) string {
	var maps = GetMapList()

	var idMap = 0

	for _, mp := range maps {
		if mp.Name == mapName {
			idMap = mp.Id
		}
	}
	idGame := SendToDB(game.Name, idMap)

	usersAndRespId := make(map[User]Respawn)
	for userName := range game.Users {
		for respawns := range game.Respawns {
			if game.Respawns[respawns] == userName {
				user := GetUsers("WHERE name='" + userName + "'")
				usersAndRespId[user] = respawns
			}
		}
	}

	if len(usersAndRespId) > 1 {
		UsersToDB(idGame, usersAndRespId)
		return idGame
	} else {
		return ""
	}
}

func SendToDB(Name string, idMap int) string {
	var err error

	_, err = db.Exec("INSERT INTO action_games (name, id_map, step, phase, winner) VALUES ($1, $2, $3, $4, $5)", // добавляем новую игру в БД
		Name, idMap, 0, "Init", "") // id карты, 0 - ход, Фаза Инициализации (растановка войск), id первого, второго игрока, цена для покупку моба 1, 2 игрока, игра не завершена
	if err != nil {
		log.Fatal(err)
	}

	var id string
	rows, err := db.Query("Select id FROM action_games ORDER BY id DESC LIMIT 1")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}

	return id
}

func UsersToDB(id string, usersAndResp map[User]Respawn) {
	var err error

	for user, resp := range usersAndResp {
		_, err = db.Exec("INSERT INTO action_game_structure (id_game, id_type, id_user, x, y) VALUES ($1, $2, $3, $4, $5)",
			id, 1, user.Id, resp.X, resp.Y)    // добавляем респаун игрока
		if err != nil {
			log.Fatal(err)
		}

		var idResp string
		rows, err := db.Query("Select id FROM action_game_structure WHERE id_game=$1 AND id_type=$2 AND id_user=$3 ORDER BY id DESC LIMIT 1", id, 1, user.Id)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&idResp)
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err = db.Exec("INSERT INTO action_game_user (id_game, id_user, start_structure, price, ready) VALUES ($1, $2, $3, $4, $5)", // добавляем новую игру в БД
			id, user.Id, idResp, 100, "false") // id карты, 0 - ход, id респа,  Фаза Инициализации (растановка войск), id первого, второго игрока, цена для покупку моба 1, 2 игрока, игра не завершена
		if err != nil {
			log.Fatal(err)
		}
	}
}
