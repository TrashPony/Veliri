package lobby

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"strconv"
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
	openGames[Games{nameGame, nameMap, nameCreator, ""}] = true
}
func ConnectGame(nameGame string, userName string) (string, bool)  {
	for game := range openGames {
		if game.nameGame == nameGame{
			game.nameNewPlayer = userName
			StartNewGame(game.nameMap, game)
			DelNewGame(game.nameCreator)
			return game.nameCreator, true
		}
	}
	return "", false
}
func DelNewGame(nameCreator string)  {
	for game := range openGames {
		if game.nameCreator == nameCreator{
			delete(openGames,game)
		}
	}
}

func GetMapList()([]Map)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * FROM map")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var maps = make([]Map, 0)
	var mp Map

	for rows.Next() {
		err := rows.Scan(&mp.id, &mp.name, &mp.xSize, &mp.ySize, &mp.Type)
		if err != nil {
			log.Fatal(err)
		}
		maps = append(maps, mp)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return maps
}

func DontEndGames(UserName string)(string)  {
	var users = GetUsers()
	var playerId int = 0
	for _, user := range users {
		if user.name == UserName {
			playerId = user.id
		}
	}

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select name FROM activegame WHERE idplayer1=" + strconv.Itoa(playerId) + " OR idplayer2=" + strconv.Itoa(playerId))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var games string = ""
	var game ActiveGames

	for rows.Next() {
		err := rows.Scan(&game.name)
		if err != nil {
			log.Fatal(err)
		}
		games = games + game.name  + ":"
	}

	return games
}

func GetUsers()([]User)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users = make([]User, 0)
	var user User

	for rows.Next() {
		err := rows.Scan(&user.id, &user.name, &user.password, &user.mail)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}

func MapList()(string)  {

	var maps = GetMapList()
	var responseNameMap = ""
	for _, bk := range maps {
		responseNameMap = responseNameMap + bk.name + ":"
	}
	return responseNameMap
}

type User struct {
	id int
	name string
	password string
	mail string
}

type Games struct{
	nameGame      string
	nameMap       string
	nameCreator   string
	nameNewPlayer string
}

type ActiveGames struct{
	name      string
}

type Map struct {
	id	int
	name    string
	xSize   int
	ySize   int
	Type    string
}