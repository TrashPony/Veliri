package box

import (
	"../../../dbConnect"
	"../../gameObjects/box"
	"log"
)

func Insert(newBox *box.Box) *box.Box {
	id := 0
	err := dbConnect.GetDBConnect().QueryRow("INSERT INTO "+
		"box_in_map "+
		"(password, destroy_time, id_map, id_box_type, q, r, rotate) "+
		"VALUES "+
		"($1, $2, $3, $4, $5, $6, $7) "+
		"RETURNING id",
		newBox.GetPassword(), newBox.DestroyTime, newBox.MapID, newBox.TypeID, newBox.Q, newBox.R, newBox.Rotate).Scan(&id)
	if err != nil {
		log.Fatal("add new box " + err.Error())
	}

	newBox.ID = id
	getTypeBox(newBox)
	Inventory(newBox)

	return newBox
}
