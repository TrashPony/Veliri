package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"log"
)

func InsertMapCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map) {
	//todo positions
	_, err := dbConnect.GetDBConnect().Exec(""+
		"INSERT INTO map_constructor ("+
		"id_map, "+
		"q, "+
		"r, "+
		"texture_over_flore, "+
		"texture_priority, "+
		"level, "+
		"rotate, "+
		"animate_speed, "+
		"x_offset, "+
		"y_offset, "+
		"x_shadow_offset, "+
		"y_shadow_offset, "+
		"shadow_intensity, "+
		"scale, "+
		"shadow, "+
		"transport, "+
		"handler, "+
		"to_base_id, "+
		"to_map_id,"+
		"id_type, "+
		"object_priority "+
		""+
		") "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)",
		mp.Id,
		coordinate.Q,
		coordinate.R,
		coordinate.TextureOverFlore,
		coordinate.TexturePriority,
		coordinate.Level,
		coordinate.ObjRotate,
		coordinate.AnimationSpeed,
		coordinate.XOffset,
		coordinate.YOffset,
		coordinate.XShadowOffset,
		coordinate.YShadowOffset,
		coordinate.ShadowIntensity,
		coordinate.Scale,
		coordinate.Shadow,
		coordinate.Transport,
		coordinate.Handler,
		coordinate.ToBaseID,
		coordinate.ToMapID,
		coordinate.ID,
		coordinate.ObjectPriority,
	)
	if err != nil {
		log.Fatal("add new mc coordinate " + err.Error())
	}
}
