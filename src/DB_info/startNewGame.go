package DB_info

import (
	"database/sql"
	"log"
)

func StartNewGame(nameGame string) (string, bool)  {
	for game := range openGames {
		if openGames[game].nameGame == nameGame && openGames[game].nameNewPlayer != ""{
			id := InitNewGame(openGames[game].nameMap, openGames[game])
			return id, true
		}
	}
	return "", false
}

func InitNewGame(mapName string, game Games)(string) {
	var maps = GetMapList()

	var idMap int = 0
	var idPlayer1 int = GetID("WHERE name='" + game.nameCreator + "'")
	var idPlayer2 int = GetID("WHERE name='" + game.nameNewPlayer + "'")

	for _, mp := range maps {
		if mp.name == mapName{
			idMap = mp.id
		}
	}
	id := SendToDB(game.nameGame ,idMap, idPlayer1, idPlayer2)
	return id
}

func SendToDB(Name string, idMap int, idPlayer1 int, idPlayer2 int)(string)  {
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

	UsersToDB(id, idPlayer1, idPlayer2)

	return id
}

func UsersToDB(id string, play1 int, play2 int)  {

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game")
	if err != nil {
		log.Fatal(err)
	}

	_ ,err = db.Exec("INSERT INTO action_game_user (id_game, id_user, price) VALUES ($1, $2, $3)",    // добавляем новую игру в БД
		id, play1, 100) // id карты, 0 - ход, Фаза Инициализации (растановка войск), id первого, второго игрока, цена для покупку моба 1, 2 игрока, игра не завершена
	if err != nil {
		log.Fatal(err)
	}

	_ ,err = db.Exec("INSERT INTO action_game_user (id_game, id_user, price) VALUES ($1, $2, $3)",    // добавляем новую игру в БД
		id, play2, 100) // id карты, 0 - ход, Фаза Инициализации (растановка войск), id первого, второго игрока, цена для покупку моба 1, 2 игрока, игра не завершена
	if err != nil {
		log.Fatal(err)
	}
}
