package lobby

import (
	"database/sql"
	"log"
)

func StartNewGame(mapName string, game Games)  {
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
	SendToDB(idMap, idPlayer1, idPlayer2)
}

func SendToDB(idMap int, idPlayer1 int, idPlayer2 int)(int64, error)  {

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game")
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("INSERT INTO activegame (idmap, step, phase, idplayer1, idplayer2, price1, price2, gameend) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",    // добавляем новую игру в БД
		idMap, 0, "Init", idPlayer1, idPlayer2, 100, 100, false) // id карты, 0 - ход, Фаза Инициализации (растановка войск), id первого, второго игрока, цена для покупку моба 1, 2 игрока, игра не завершена
	if err != nil {
		log.Fatal(err)
	}

	return res.RowsAffected()
}
