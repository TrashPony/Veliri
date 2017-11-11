package lobby

import (
	"log"
)

func GetRespawns(nameMap string)([]Respawn)  {

	rows, err := db.Query("Select * FROM respawns WHERE id_map = (Select id from map WHERE name=$1)", nameMap)
	if err != nil {
		log.Fatal(err)
	}

	var respawns = make([]Respawn, 0)
	var resp Respawn

	for rows.Next() {
		err := rows.Scan(&resp.Id, &resp.Name,&resp.IdMap, &resp.X, &resp.Y)
		if err != nil {
			log.Fatal(err)
		}
		respawns = append(respawns, resp)
	}
	return respawns
}