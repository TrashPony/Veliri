package get

import (
	"log"
	"../../../dbConnect"
	"../../gameObjects/coordinate"
)

func Respawns(mapID int) map[int]*coordinate.Coordinate {

	rows, err := dbConnect.GetDBConnect().Query("Select id, x, y " +
		"FROM map_constructor " +
		"WHERE id_type=1 AND id_map = $1", mapID)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var respawns = make(map[int]*coordinate.Coordinate)

	for rows.Next() {
		var resp coordinate.Coordinate

		err := rows.Scan(&resp.ID, &resp.X, &resp.Y)
		if err != nil {
			log.Fatal(err)
		}

		respawns[resp.ID] = &resp
	}
	return respawns
}
