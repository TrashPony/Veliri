package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"log"
)

func UpdateMapCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map) {
	//todo positions

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
		"to_base_id = $18, "+
		"to_map_id = $19,"+
		"id_type = $20, "+
		"object_priority = $21 "+
		""+
		"WHERE id_map = $1 AND x=$2 AND y = $3",
		mp.Id,
		coordinate.X,
		coordinate.Y,
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
		log.Fatal("update mc coordinate" + err.Error())
	}
}
