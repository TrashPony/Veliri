package getLocalGame

import (
	"../../gameObjects/map"
	"../../localGame"
	"../../../dbConnect"
	"log"
)

func Map(game *localGame.Game) _map.Map {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, x_size, y_size, id_type, level, specification FROM maps WHERE id = $1", game.MapID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mp _map.Map
	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.XSize, &mp.YSize, &mp.DefaultTypeID, &mp.DefaultLevel, &mp.Specification)
		if err != nil {
			log.Fatal(err)
		}
	}

	CoordinatesMap(&mp, game)

	return mp
}
