package dbo

import (
	"../gameMap"
	"log"
)

func GetMap(idMap int) gameMap.Map {

	rows, err := db.Query("Select * FROM maps WHERE id = $1", idMap)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mp gameMap.Map
	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.XSize, &mp.YSize, &mp.Type, &mp.Specification)
		if err != nil {
			log.Fatal(err)
		}
	}

	GetCoordinatesMap(&mp)

	return mp
}
