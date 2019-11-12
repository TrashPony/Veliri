package mapEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

func UpdateMapCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map, oldX, oldY int) {
	// TODO
	//positions, err := json.Marshal(coordinate.Positions)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//_, err = dbConnect.GetDBConnect().Exec(""+
	//	"UPDATE map_constructor "+
	//	"SET "+
	//	"texture_over_flore = $4, "+
	//	"texture_priority = $5, "+
	//	"rotate = $6, "+
	//	"animate_speed = $7, "+
	//	"x_shadow_offset = $8, "+
	//	"y_shadow_offset = $9, "+
	//	"shadow_intensity = $10, "+
	//	"scale = $11, "+
	//	"shadow = $12, "+
	//	"transport = $13, "+
	//	"handler = $14, "+
	//	"to_base_id = $15, "+
	//	"to_map_id = $16,"+
	//	"id_type = $17, "+
	//	"object_priority = $18,"+
	//	"to_positions = $19,"+
	//	"x = $20,"+
	//	"y = $21 "+
	//	""+
	//	"WHERE id_map = $1 AND x=$2 AND y = $3",
	//	mp.Id,
	//	oldX,
	//	oldY,
	//	coordinate.TextureOverFlore,
	//	coordinate.TexturePriority,
	//	coordinate.ObjRotate,
	//	coordinate.AnimationSpeed,
	//	coordinate.XShadowOffset,
	//	coordinate.YShadowOffset,
	//	coordinate.ShadowIntensity,
	//	coordinate.Scale,
	//	coordinate.Shadow,
	//	coordinate.Transport,
	//	coordinate.Handler,
	//	coordinate.ToBaseID,
	//	coordinate.ToMapID,
	//	coordinate.ID,
	//	coordinate.ObjectPriority,
	//	string(positions),
	//	coordinate.X,
	//	coordinate.Y,
	//)
	//if err != nil {
	//	log.Fatal("update mc coordinate" + err.Error())
	//}
}
