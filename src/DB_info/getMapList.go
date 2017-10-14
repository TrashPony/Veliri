package DB_info

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

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
		err := rows.Scan(&mp.Id, &mp.Name, &mp.XSize, &mp.YSize, &mp.Type)
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