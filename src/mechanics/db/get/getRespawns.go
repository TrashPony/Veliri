package get

import (
	"log"
	"../../../dbConnect"
	"../../gameObjects/coordinate"
)

func Respawns(mapID int) []*coordinate.Coordinate {

	rows, err := dbConnect.GetDBConnect().Query("Select id, x, y " +
		"FROM map_constructor " +
		"WHERE id_type=1 AND id_map = $1", mapID)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var respawns = make([]*coordinate.Coordinate, 0)

	for rows.Next() {
		var resp coordinate.Coordinate

		err := rows.Scan(&resp.ID, &resp.X, &resp.Y)
		if err != nil {
			log.Fatal(err)
		}

		respawns = append(respawns, &resp)
	}
	return respawns
}
