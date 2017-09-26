package lobby

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var openGames = make(map[Games]bool)

func OpenGameList()([]string) {
	var responseNameGame = ""
	var responseNameMap = ""
	var responseNameUser = ""

	for game := range openGames {
		responseNameGame = responseNameGame + game.nameGame + ":" ;
		responseNameMap = responseNameMap + game.nameMap + ":";
		responseNameUser = responseNameUser + game.nameCreator + ":";
	}

	var games []string
	games = append(games, responseNameGame)
	games = append(games, responseNameMap)
	games = append(games, responseNameUser)

	return games
}

func CreateNewGame (nameGame string, nameMap string, nameCreator string ) {
	openGames[Games{nameGame, nameMap, nameCreator}] = true
}
func DelNewGame(nameCreator string)  {
	println("зашел")
	for game := range openGames {
		println(game.nameCreator)
		println(nameCreator)
		if game.nameCreator == nameCreator{
			println("Удалил")
			delete(openGames,game)
		}
	}
}

func MapList()(string)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * FROM map")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	maps := make([]*Map, 0)
	for rows.Next() {
		mp := new(Map)
		err := rows.Scan(&mp.id, &mp.name, &mp.xSize, &mp.ySize, &mp.Type)
		if err != nil {
			log.Fatal(err)
		}
		maps = append(maps, mp)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	var responseNameMap = ""
	for _, bk := range maps {
		responseNameMap = responseNameMap + bk.name + ":";
	}
	return responseNameMap
}

type Games struct{
	nameGame string
	nameMap string
	nameCreator string
}

type Map struct {
	id	int
	name    string
	xSize   int
	ySize   int
	Type    string
}