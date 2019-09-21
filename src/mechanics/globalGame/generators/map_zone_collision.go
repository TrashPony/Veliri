package generators

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

// создаем сектора с данными о проходимости для быстродействия методов движения
func UpdateMapZoneCollision() {
	for _, mp := range maps.Maps.GetAllMap() {
		FillMapZone(mp)
	}
}

func FillMapZone(mp *_map.Map) {
	mp.GeoZone = make([][][]*_map.ObstaclePoint, 100)

	for x := 50; x < mp.QSize*game_math.HexagonWidth; x += 100 {

		mp.GeoZone[x/100] = make([][]*_map.ObstaclePoint, 100)

		for y := 50; y < mp.RSize*game_math.HexagonHeight; y += 100 {

			mp.GeoZone[x/100][y/100] = make([]*_map.ObstaclePoint, 0)

			collisions.FillMapZone(x, y, &mp.GeoZone[x/100][y/100], mp)
		}
	}
}
