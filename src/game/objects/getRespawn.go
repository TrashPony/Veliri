package objects

import (
	"database/sql"
	"log"
)

func GetRespawns(idUser int, idGame string)(Respawn)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select r.id, r.name, r.id_map, r.x, r.y FROM respawns as r, map, action_game_user as agu WHERE r.id=(Select respawns_id FROM action_game_user as agu WHERE agu.id_game=$1 and agu.id_user=$2 LIMIT 1)", idGame, idUser)
	if err != nil {
		log.Fatal(err)
	}

	var resp Respawn

	for rows.Next() {
		err := rows.Scan(&resp.Id, &resp.Name,&resp.IdMap, &resp.X, &resp.Y)
		if err != nil {
			log.Fatal(err)
		}
	}
	return resp
}

type Respawn struct {
	Id	  int
	Name  string
	IdMap int
	X     int
	Y     int
}

