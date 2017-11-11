package lobby

import (
	"log"
)

func GetMapList()([]Map)  {

	rows, err := db.Query("Select * FROM map")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var maps = make([]Map, 0)
	var mp Map

	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.XSize, &mp.YSize, &mp.Type)
		if err != nil {
			log.Fatal(err)
		}
		row := db.QueryRow("SELECT COUNT(*) as Respawns FROM respawns WHERE id_map=$1", mp.Id)
		errors := row.Scan(&mp.Respawns)
		if errors != nil {
			log.Fatal(errors)
		}
		maps = append(maps, mp)
	}
	return maps
}