package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"log"
)

func UpdateMapCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map) {
	_, err := dbConnect.GetDBConnect().Exec(""+
		"UPDATE map_constructor "+
		"SET "+
		"texture_over_flore = $4, "+
		"texture_priority = $5, "+
		"level = $6, "+
		"rotate = $7, "+
		"animate_speed = $8, "+
		"x_offset = $9, "+
		"y_offset = $10, "+
		"x_shadow_offset = $11, "+
		"y_shadow_offset = $12, "+
		"shadow_intensity = $13, "+
		"scale = $14, "+
		"shadow = $15, "+
		"transport = $16, "+
		"handler = $17, "+
		"to_q = $18, "+
		"to_r = $19, "+
		"to_base_id = $20, "+
		"to_map_id = $21,"+
		"id_type = $22 "+
		""+
		"WHERE id_map = $1 AND q=$2 AND r = $3",
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
		coordinate.ToQ,
		coordinate.ToR,
		coordinate.ToBaseID,
		coordinate.ToMapID,
		coordinate.ID,
	)
	if err != nil {
		log.Fatal("update mc coordinate" + err.Error())
	}
}
