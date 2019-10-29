package mapEditor

import (
	"encoding/json"
	"fmt"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"log"
)

func InsertMapCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map) {
	positions, err := json.Marshal(coordinate.Positions)
	if err != nil {
		fmt.Println(err)
	}
	_, err = dbConnect.GetDBConnect().Exec(""+
		"INSERT INTO map_constructor ("+
		"id_map, "+
		"x, "+
		"y, "+
		"texture_over_flore, "+
		"texture_priority, "+
		"rotate, "+
		"animate_speed, "+
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
		"object_priority,"+
		"to_positions "+
		""+
		") "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)",
		mp.Id,
		coordinate.X,
		coordinate.Y,
		coordinate.TextureOverFlore,
		coordinate.TexturePriority,
		coordinate.ObjRotate,
		coordinate.AnimationSpeed,
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
		log.Fatal("add new mc coordinate " + err.Error())
	}
}
