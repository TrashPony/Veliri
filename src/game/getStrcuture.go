package game

import (
	"log"
)

func GetAllStrcuture(idGame int) (structures map[int]map[int]*Structure) {

	rows, err := db.Query("Select st.type, users.name, st.range_view, struct.x, struct.y FROM action_game_structure as struct, structure_type as st, users WHERE users.id=struct.id_user AND st.id=struct.id_type AND id_game=$1", idGame)
	if err != nil {
		log.Fatal(err)
	}


	structures = make(map[int]map[int]*Structure)

	for rows.Next() {
		var structure Structure
		err := rows.Scan(&structure.Type, &structure.NameUser, &structure.WatchZone, &structure.X, &structure.Y)
		if err != nil {
			log.Fatal(err)
		}

		if structures[structure.X] != nil {
			structures[structure.X][structure.Y] = &structure
		} else {
			structures[structure.X] = make(map[int]*Structure)
			structures[structure.X][structure.Y] = &structure
		}
	}

	return
}