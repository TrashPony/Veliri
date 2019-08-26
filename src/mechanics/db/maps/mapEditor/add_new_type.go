package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func AddNewTypeCoordinate(typeCoordinate, textureFlore, textureObject, animateSpriteSheets string, animateLoop, move, view, attack bool) int {

	var id int
	err := dbConnect.GetDBConnect().QueryRow("INSERT INTO coordinate_type "+
		"(type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, unit_overlap) "+
		" "+
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		typeCoordinate, textureFlore, textureObject, animateSpriteSheets, animateLoop, move, view, attack, true).Scan(&id)
	if err != nil {
		log.Fatal("add new global type " + err.Error())
	}

	return id
}
