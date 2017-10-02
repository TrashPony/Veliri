package createUnit

import (
	"database/sql"
	"log"
)

func CreateUnit(idGame string, idPlayer string, unitType string, x string, y string)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд

	if err != nil {
		log.Fatal(err)
	}

	idType, hp := GetUnitType(unitType)	

	rows, err := db.Query("INSERT INTO actiongamesunit (idgame, idunittype, idplayer, hp, action, idtarget, x, y) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		idGame, idType, idPlayer, hp, true, 0, x, y)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


}

func GetUnitType(unitType string) (string, string)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select id, hp, price From unittype WHERE type=" + "'" +unitType + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id string
	var hp string
	var price int
	for rows.Next() {
		err := rows.Scan(&id, &hp, &price)
		if err != nil {
			log.Fatal(err)
		}
	}

	Price(price)

	return id, hp
}

func Price(price int)  {

}
