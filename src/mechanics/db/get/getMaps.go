package get

import (
	"../../../dbConnect"
	gameMap "../../gameObjects/map"
	"log"
)

func MapList() []gameMap.Map {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, q_size, r_size, id_type, level, specification FROM maps")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var maps = make([]gameMap.Map, 0)
	var mp gameMap.Map

	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.QSize, &mp.RSize, &mp.DefaultTypeID, &mp.DefaultLevel, &mp.Specification)
		if err != nil {
			log.Fatal(err)
		}
		row := dbConnect.GetDBConnect().QueryRow("SELECT COUNT(*) as Respawns "+
			"FROM map_constructor "+
			"WHERE id_type=1 AND id_map = $1;", mp.Id)
		errors := row.Scan(&mp.Respawns)
		if errors != nil {
			log.Fatal(errors)
		}
		maps = append(maps, mp)
	}
	return maps
}

func Map(id int) gameMap.Map {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, q_size, r_size, id_type, level, specification FROM maps WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mp gameMap.Map

	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.QSize, &mp.RSize, &mp.DefaultTypeID, &mp.DefaultLevel, &mp.Specification)
		if err != nil {
			log.Fatal(err)
		}
		row := dbConnect.GetDBConnect().QueryRow("SELECT COUNT(*) as Respawns "+
			"FROM map_constructor "+
			"WHERE id_type=1 AND id_map = $1;", mp.Id)
		errors := row.Scan(&mp.Respawns)
		if errors != nil {
			log.Fatal(errors)
		}
	}

	return mp
}
