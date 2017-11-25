package lobby

import (
	"log"
)

func GetRespawns(nameMap string) []Respawn {

	rows, err := db.Query("Select id, x, y, id_map, name FROM map_constructor WHERE type='respawn' AND id_map = (Select id from map WHERE name=$1)", nameMap)
	if err != nil {
		log.Fatal(err)
	}

	var respawns = make([]Respawn, 0)
	var resp Respawn

	for rows.Next() {
		err := rows.Scan(&resp.Id, &resp.X, &resp.Y, &resp.IdMap, &resp.Name)
		if err != nil {
			log.Fatal(err)
		}
		respawns = append(respawns, resp)
	}
	return respawns
}

type Respawn struct {
	Id    int
	Name  string
	IdMap int
	X     int
	Y     int
}
