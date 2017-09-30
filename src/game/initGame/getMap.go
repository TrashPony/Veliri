package initGame

import (
	"database/sql"
	"log"
	"strconv"
)

func GetMap(idMap int)(int, int)  {
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
		err := rows.Scan(&mp.id, &mp.name, &mp.xsize, &mp.ysize, &mp.Type)
		if err != nil {
			log.Fatal(err)
		}
	}

	return mp.xsize, mp.ysize
}

type Map struct {
	id int
	name string
	xsize int
	ysize int
        Type string

}