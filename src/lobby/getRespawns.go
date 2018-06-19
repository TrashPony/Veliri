package lobby

import (
	"log"
	"strconv"
)

func GetRespawns(mapID int) []*Respawn {

	rows, err := db.Query("Select id, x, y, id_map " +
		"FROM map_constructor " +
		"WHERE id_type=1 AND id_map = $1", mapID)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var respawns = make([]*Respawn, 0)

	for rows.Next() {
		var resp Respawn

		err := rows.Scan(&resp.Id, &resp.X, &resp.Y, &resp.IdMap)
		if err != nil {
			log.Fatal(err)
		}

		resp.Name = strconv.Itoa(len(respawns))
		respawns = append(respawns, &resp)
	}
	return respawns
}

type Respawn struct {
	Id       int
	Name     string
	IdMap    int
	X        int
	Y        int
	UserName string
}
