package objects

import (
	"database/sql"
	"log"
	"strconv"
)

func GetMap(idMap int)(Map)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * FROM map WHERE id =" + strconv.Itoa(idMap))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mp Map
	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.Xsize, &mp.Ysize, &mp.Type)
		if err != nil {
			log.Fatal(err)
		}
	}

	return mp
}

type Map struct {
	Id int
	Name string
	Xsize int
	Ysize int
	Type string
}