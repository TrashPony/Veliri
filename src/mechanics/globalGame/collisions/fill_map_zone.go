package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

func FillMapZone(x, y int, zonePoints *[]*_map.ObstaclePoint, mp *_map.Map) {
	zoneRect := getRect(float64(x), float64(y), 200, 200)

	for i := 0; i < len(mp.GeoData); i++ {
		if zoneRect.detectCollisionRectToCircle(&point{x: float64(mp.GeoData[i].X), y: float64(mp.GeoData[i].Y)}, mp.GeoData[i].Radius) {
			*zonePoints = append(*zonePoints, mp.GeoData[i])
		}
	}
}
