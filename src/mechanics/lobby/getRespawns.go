package lobby

import (
	"log"
	"../../dbConnect"
	"../localGame/map/coordinate"
)

func GetRespawns(mapID int) []*coordinate.Coordinate {

	rows, err := dbConnect.GetDBConnect().Query("Select x, y, id_map " +
		"FROM map_constructor " +
		"WHERE id_type=1 AND id_map = $1", mapID)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var respawns = make([]*coordinate.Coordinate, 0)

	for rows.Next() {
		var resp coordinate.Coordinate

		err := rows.Scan(&resp.X, &resp.Y)
		if err != nil {
			log.Fatal(err)
		}

		respawns = append(respawns, &resp)
	}
	return respawns
}
