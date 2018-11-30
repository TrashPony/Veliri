package mapEditor

import (
	"../../../dbConnect"
	"log"
)

func AddNewTypeCoordinate(typeCoordinate, textureFlore, textureObject, animateSpriteSheets string, animateLoop, move,
	view, attack bool, impactRadius int, scale int, shadow bool) int {
	var id int
	err := dbConnect.GetDBConnect().QueryRow("INSERT INTO coordinate_type "+
		"(type, texture_flore, texture_object, animate_sprite_sheets, animate_loop, move, view, attack, impact_radius, scale, shadow) "+
		" "+
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		typeCoordinate, textureFlore, textureObject, animateSpriteSheets, animateLoop, move, view, attack, impactRadius,
		scale, shadow).Scan(&id)
	if err != nil {
		log.Fatal("add new global type " + err.Error())
	}

	return id
}
