package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"log"
)

func Map(game *localGame.Game) _map.Map {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, q_size, r_size, id_type, level, specification, global, in_game "+
		"FROM maps WHERE id = $1", game.MapID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mp _map.Map
	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.QSize, &mp.RSize, &mp.DefaultTypeID, &mp.DefaultLevel,
			&mp.Specification, &mp.Global, &mp.InGame)
		if err != nil {
			log.Fatal(err)
		}
	}

	CoordinatesMap(&mp, game)

	return mp
}
