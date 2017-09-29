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
	var users = GetUsers()

	var idMap int = 0
	var idPlayer1 int = 0
	var idPlayer2 int = 0


	for _, user := range users {
		if user.name == game.nameCreator{
			idPlayer1 = user.id
		}
		if user.name == game.nameNewPlayer{
			idPlayer2 = user.id
		}
	}

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

	_ ,err = db.Exec("INSERT INTO activegame (name, idmap, step, phase, idplayer1, idplayer2, price1, price2, gameend) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",    // добавляем новую игру в БД
		Name, idMap, 0, "Init", idPlayer1, idPlayer2, 100, 100, false) // id карты, 0 - ход, Фаза Инициализации (растановка войск), id первого, второго игрока, цена для покупку моба 1, 2 игрока, игра не завершена
	if err != nil {
		log.Fatal(err)
	}

	var id string

	rows, err := db.Query("Select id FROM activegame ORDER BY id DESC LIMIT 1")

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}

	return id
}
