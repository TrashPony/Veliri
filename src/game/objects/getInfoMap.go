package objects

import (
	"log"
	"strconv"
)

func GetInfoMap(idMap int) Map {

	rows, err := db.Query("Select * FROM maps WHERE id =" + strconv.Itoa(idMap))
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
	Id    int
	Name  string
	Xsize int
	Ysize int
	Type  string
}
