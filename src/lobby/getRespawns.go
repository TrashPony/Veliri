package lobby

import (
	"log"
	"strconv"
)

func GetRespawns(mapID int) []*Respawn {

	rows, err := db.Query("Select id, x, y, id_map FROM map_constructor WHERE type='respawn' AND id_map = (Select id from maps WHERE id=$1)", mapID)
	if err != nil {
		log.Fatal(err)
	}

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
