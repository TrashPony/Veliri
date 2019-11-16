package plant_life_game

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func Evolution(mp *_map.Map) {

	checkCoordinate := func(x, y int) bool {

		obj := mp.GetDynamicObjects(x, y)
		if obj != nil && obj.Texture == "plant_4" {
			return true
		} else {
			return false
		}
	}

	checkCountNeighbors := func(x, y int, life bool) *lifeObject {
		count := 0
		neighbors := make([]*coordinate.Coordinate, 0)

		// если вышли за пределы карты то появляемся в противоположеной стороны
		// todo но чет это не работает
		if mp.XSize < x {
			x = x - mp.XSize
		}
		if mp.YSize < y {
			y = y - mp.YSize
		}

		//строго лево
		if checkCoordinate(x-game_math.CellSize, y) {
			count++
		} else {
			neighbors = append(neighbors, &coordinate.Coordinate{X: x - game_math.CellSize, Y: y})
		}

		//строго право
		if checkCoordinate(x+game_math.CellSize, y) {
			count++
		} else {
			neighbors = append(neighbors, &coordinate.Coordinate{X: x + game_math.CellSize, Y: y})
		}

		//верх центр
		if checkCoordinate(x, y-game_math.CellSize) {
			count++
		} else {
			neighbors = append(neighbors, &coordinate.Coordinate{X: x, Y: y - game_math.CellSize})
		}

		//низ центр
		if checkCoordinate(x, y+game_math.CellSize) {
			count++
		} else {
			neighbors = append(neighbors, &coordinate.Coordinate{X: x, Y: y + game_math.CellSize})
		}

		//верх лево
		if checkCoordinate(x-game_math.CellSize, y-game_math.CellSize) {
			count++
		} else {
			neighbors = append(neighbors, &coordinate.Coordinate{X: x - game_math.CellSize, Y: y - game_math.CellSize})
		}

		//верх право
		if checkCoordinate(x+game_math.CellSize, y-game_math.CellSize) {
			count++
		} else {
			neighbors = append(neighbors, &coordinate.Coordinate{X: x + game_math.CellSize, Y: y - game_math.CellSize})
		}

		//низ лево
		if checkCoordinate(x-game_math.CellSize, y+game_math.CellSize) {
			count++
		} else {
			neighbors = append(neighbors, &coordinate.Coordinate{X: x - game_math.CellSize, Y: y + game_math.CellSize})
		}

		//низ право
		if checkCoordinate(x+game_math.CellSize, y+game_math.CellSize) {
			count++
		} else {
			neighbors = append(neighbors, &coordinate.Coordinate{X: x + game_math.CellSize, Y: y + game_math.CellSize})
		}

		return &lifeObject{X: x, Y: y, life: life, Neighbors: neighbors, countLifeNeighbors: count}
	}

	lifeObjects := make([]*lifeObject, 0)

	// смотрим все живые клетки на наличие соседей, и будущего состояния жизни
	for _, x := range mp.GetCopyMapDynamicObjects() {
		for _, obj := range x {
			if obj.Texture == "plant_4" {

				lifeObj := checkCountNeighbors(obj.X, obj.Y, true)

				if lifeObj.countLifeNeighbors == 3 || lifeObj.countLifeNeighbors == 2 {
					lifeObj.futureLife = true
				}

				lifeObjects = append(lifeObjects, lifeObj)
			}
		}
	}

	// смотрим все мертвые клетки будущее состояния жизни
	for _, lifeUnit := range lifeObjects {
		for _, deadUnit := range lifeUnit.Neighbors {

			deadObj := checkCountNeighbors(deadUnit.X, deadUnit.Y, false)

			if deadObj.countLifeNeighbors == 3 {
				deadObj.futureLife = true
				lifeObjects = append(lifeObjects, deadObj)
			}
		}
	}

	// мы применяем действия к клеткам
	for _, terrainUnit := range lifeObjects {
		if !terrainUnit.life && terrainUnit.futureLife {
			createPlant(terrainUnit.X, terrainUnit.Y, "plant_4", mp)
		}

		if terrainUnit.life && !terrainUnit.futureLife {
			mp.RemoveDynamicObjectByXY(terrainUnit.X, terrainUnit.Y)
		}
	}
}
