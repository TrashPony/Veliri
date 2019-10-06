package mapEditor

import (
	"encoding/json"
	"fmt"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"log"
)

func UpdateMapCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map) {
	positions, err := json.Marshal(coordinate.Positions)
	if err != nil {
		fmt.Println(err)
	}

	_, err = dbConnect.GetDBConnect().Exec(""+
		"UPDATE map_constructor "+
		"SET "+
		"texture_over_flore = $4, "+
		"texture_priority = $5, "+
		"rotate = $6, "+
		"animate_speed = $7, "+
		"x_offset = $8, "+
		"y_offset = $9, "+
		"x_shadow_offset = $10, "+
		"y_shadow_offset = $11, "+
		"shadow_intensity = $12, "+
		"scale = $13, "+
		"shadow = $14, "+
		"transport = $15, "+
		"handler = $16, "+
		"to_base_id = $17, "+
		"to_map_id = $18,"+
		"id_type = $19, "+
		"object_priority = $20,"+
		"to_positions = $21 "+
		""+
		"WHERE id_map = $1 AND x=$2 AND y = $3",
		mp.Id,
		coordinate.X,
		coordinate.Y,
		coordinate.TextureOverFlore,
		coordinate.TexturePriority,
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
		string(positions),
	)
	if err != nil {
		log.Fatal("update mc coordinate" + err.Error())
	}
}
