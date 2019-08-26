package box

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"log"
)

func Insert(newBox *boxInMap.Box) *boxInMap.Box {

	if newBox.MapID != 0 {
		id := 0
		err := dbConnect.GetDBConnect().QueryRow("INSERT INTO "+
			"box_in_map "+
			"(password, destroy_time, id_map, id_box_type, q, r, rotate, current_hp) "+
			"VALUES "+
			"($1, $2, $3, $4, $5, $6, $7, $7) "+
			"RETURNING id",
			newBox.GetPassword(), newBox.DestroyTime, newBox.MapID, newBox.TypeID, newBox.Q, newBox.R, newBox.Rotate, newBox.HP).Scan(&id)
		if err != nil {
			log.Fatal("add new box " + err.Error())
		}
		newBox.ID = id
		Inventory(newBox)
	}
	getTypeBox(newBox)

	return newBox
}
