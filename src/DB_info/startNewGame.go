package DB_info

import (
	"database/sql"
	"log"
)

func StartNewGame(nameGame string) (string, bool)  {
	for game := range openGames {
		if openGames[game].Name == nameGame && len(openGames[game].Users) > 1{
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

func InitNewGame(mapName string, game LobbyGames)(string) {
	var maps = GetMapList()

	var idMap int = 0


	for _, mp := range maps {
		if mp.Name == mapName{
			idMap = mp.Id
		}
	}
	idGame := SendToDB(game.Name, idMap)

	usersId := make([]int,0)
	for userName := range game.Users {
		user := GetUsers("WHERE name='" + userName + "'")
		usersId = append(usersId, user.Id)
	}
	if len(usersId) > 1 {
		UsersToDB(idGame, usersId)
		return idGame
	} else {
		return ""
	}
}

func SendToDB(Name string, idMap int)(string)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game")
	if err != nil {
		log.Fatal(err)
	}

	_ ,err = db.Exec("INSERT INTO action_games (name, id_map, step, phase, winner) VALUES ($1, $2, $3, $4, $5)",    // добавляем новую игру в БД
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

func UsersToDB(id string, usersId []int)  {

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(usersId); i++ {
		_, err = db.Exec("INSERT INTO action_game_user (id_game, id_user, price, ready) VALUES ($1, $2, $3, $4)", // добавляем новую игру в БД
			id, usersId[i], 100, "false") // id карты, 0 - ход, Фаза Инициализации (растановка войск), id первого, второго игрока, цена для покупку моба 1, 2 игрока, игра не завершена
		if err != nil {
			log.Fatal(err)
		}
	}
}
