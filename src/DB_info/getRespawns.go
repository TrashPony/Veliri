package DB_info

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func GetRespawns(nameMap string)([]Respawn)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * FROM respawns WHERE id_map = (Select id from map WHERE name=$1)", nameMap)
	if err != nil {
		log.Fatal(err)
	}

	var respawns = make([]Respawn, 0)
	var resp Respawn

	for rows.Next() {
		err := rows.Scan(&resp.Id, &resp.IdMap, &resp.X, &resp.Y)
		if err != nil {
			log.Fatal(err)
		}
		respawns = append(respawns, resp)
	}
	return respawns
}