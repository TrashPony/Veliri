package plant_life_game

import "github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"

func GetCountPlantIMap(textureName string, mp *_map.Map) int {
	count := 0

	for _, x := range mp.GetCopyMapDynamicObjects() {
		for _, obj := range x {
			if obj.Texture == textureName {
				count++
			}
		}
	}

	return count
}
