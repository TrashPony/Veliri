package lobby

import (
	"log"
	"../../dbConnect"
)

func GetMapList() []Map {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, x_size, y_size, id_type, level, specification FROM maps")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var maps = make([]Map, 0)
	var mp Map

	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.XSize, &mp.YSize, &mp.DefaultTypeID, &mp.DefaultLevel, &mp.Specification)
		if err != nil {
			log.Fatal(err)
		}
		row := dbConnect.GetDBConnect().QueryRow("SELECT COUNT(*) as Respawns " +
			"FROM map_constructor " +
			"WHERE id_type=1 AND id_map = $1;", mp.Id)
		errors := row.Scan(&mp.Respawns)
		if errors != nil {
			log.Fatal(errors)
		}
		maps = append(maps, mp)
	}
	return maps
}

func GetMap(id int) Map {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, x_size, y_size, id_type, level, specification FROM maps WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mp Map

	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.XSize, &mp.YSize, &mp.DefaultTypeID, &mp.DefaultLevel, &mp.Specification)
		if err != nil {
			log.Fatal(err)
		}
		row := dbConnect.GetDBConnect().QueryRow("SELECT COUNT(*) as Respawns " +
			"FROM map_constructor " +
			"WHERE id_type=1 AND id_map = $1;", mp.Id)
		errors := row.Scan(&mp.Respawns)
		if errors != nil {
			log.Fatal(errors)
		}
	}

	return mp
}
