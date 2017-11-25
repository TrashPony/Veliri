package objects

import (
	"log"
)

func GetRespawns(idUser int, idGame int) Respawn {

	rows, err := db.Query("Select r.id, r.type, r.id_map, r.x, r.y FROM map_constructor as r, map, action_game_user as agu WHERE r.id=(Select start_structure FROM action_game_user as agu WHERE agu.id_game=$1 and agu.id_user=$2 LIMIT 1)", idGame, idUser)
	if err != nil {
		log.Fatal(err)
	}

	var resp Respawn

	for rows.Next() {
		err := rows.Scan(&resp.Id, &resp.Name, &resp.IdMap, &resp.X, &resp.Y)
		if err != nil {
			log.Fatal(err)
		}
	}
	return resp
}

type Respawn struct {
	Id    int
	Name  string
	IdMap int
	X     int
	Y     int
}
