package db

import (
	"../gameMap"
	"../game"
	"log"
)

func GetMap(game *game.Game) gameMap.Map {

	rows, err := db.Query("Select * FROM maps WHERE id = $1", game.MapID)
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

	GetCoordinatesMap(&mp, game)

	return mp
}
