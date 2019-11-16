package ai

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/plant_life_game"
	"time"
)

/* ФУНКЦИЯ ПО УПРОВЛЕНИЮ ОВОЩАМИ! */

// реализация алгоритма зеро плей игры "Жизнь" для динамиской популяции растений на карте
// https://ru.wikipedia.org/wiki/%D0%98%D0%B3%D1%80%D0%B0_%C2%AB%D0%96%D0%B8%D0%B7%D0%BD%D1%8C%C2%BB

func TerrainLifeInit() {
	allMaps := maps.Maps.GetAllMap()
	for _, mp := range allMaps {
		go growTerrainWorker(mp)
		go plant_life_game.PopulationWorker(mp)
	}
}

func growTerrainWorker(mp *_map.Map) {
	for {

		growCycle := 100
		growTime := 2500

		for {

			growCycle--

			for _, x := range mp.GetCopyMapDynamicObjects() {
				for _, obj := range x {
					// радар игроков сам обновит обьект
					if obj.Scale < obj.MaxScale && !collisions.CheckObjectCollision(obj, mp, true) {
						obj.Scale++
						obj.CalculateScale()
						obj.GrowTime = growTime
					}
				}
			}

			if growCycle == 0 {
				break
			}

			time.Sleep(time.Millisecond * time.Duration(growTime))
		}

		plant_life_game.Evolution(mp)
	}
}
