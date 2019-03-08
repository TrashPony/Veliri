package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func Rotate(idMap, q, r, rotate, speed, XOffset, YOffset, XShadowOffset, YShadowOffset, ShadowIntensity int) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET rotate = $1, animate_speed = $5, "+
		" x_offset = $6, y_offset = $7, x_shadow_offset = $8, y_shadow_offset = $9, shadow_intensity = $10 "+
		"WHERE id_map = $2 AND q=$3 AND r = $4",
		rotate, idMap, q, r, speed, XOffset, YOffset, XShadowOffset, YShadowOffset, ShadowIntensity)
	if err != nil {
		log.Fatal("update rotate coordinate" + err.Error())
	}
}
