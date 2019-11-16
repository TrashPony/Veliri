package ai

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/dynamic_objects"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math/rand"
	"time"
)

/* ФУНКЦИЯ ПО УПРОВЛЕНИЮ ОВОЩАМИ! */

// реализация алгоритма зеро плей игры "Жизнь" для динамиской популяции растений на карте
// https://ru.wikipedia.org/wiki/%D0%98%D0%B3%D1%80%D0%B0_%C2%AB%D0%96%D0%B8%D0%B7%D0%BD%D1%8C%C2%BB

func TerrainLifeInit() {
	allMaps := maps.Maps.GetAllMap()
	for _, mp := range allMaps {
		go growTerrainWorker(mp)
		go populationWorker(mp)
	}
}

type lifeObject struct {
	life               bool
	futureLife         bool
	countLifeNeighbors int
	Neighbors          []*coordinate.Coordinate
	X                  int
	Y                  int
}

var TypeSources = map[string][]coordinate.Coordinate{
	"glider":  {{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: -1}, {X: 1, Y: -2}},
	"flasher": {{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}},
}

func initPopulationSource(typeSource string, x, y int, texturePlant string, mp *_map.Map) {
	// TODO поворот на 90, 180 и 270 градусов
	source, _ := TypeSources[typeSource]
	for _, offset := range source {
		createPlant(
			x+(offset.X*game_math.CellSize)+game_math.CellSize/2, // game_math.CellSize/2 что бы брался пиксель в центре клетки
			y+(offset.Y*game_math.CellSize)+game_math.CellSize/2,
			texturePlant,
			mp,
		)
	}
}

func createPlant(x, y int, texturePlant string, mp *_map.Map) {
	plant := dynamic_objects.DynamicObjects.GetDynamicObjectByTexture(texturePlant, rand.Intn(360))
	plant.X = x
	plant.Y = y
	plant.Scale = 1
	plant.MaxScale = rand.Intn(22) + 3
	plant.HP = plant.MaxHP
	plant.CalculateScale()
	mp.AddDynamicObject(plant)
}

func populationWorker(mp *_map.Map) {
	// инитим первую колонию
	// TODO генерировать рандомно, проверять место куда можно поставить обьекты
	initPopulationSource("glider", 25*game_math.CellSize, 40*game_math.CellSize, "plant_4", mp)
	initPopulationSource("flasher", 20*game_math.CellSize, 40*game_math.CellSize, "plant_4", mp)
	initPopulationSource("flasher", 22*game_math.CellSize, 40*game_math.CellSize, "plant_4", mp)
	initPopulationSource("glider", 15*game_math.CellSize, 30*game_math.CellSize, "plant_4", mp)

	for {
		// todo должен следить что бы на карте всегда была популяция
		time.Sleep(time.Millisecond * 200)
	}
}

func Evolution(mp *_map.Map) {

	checkCoordinate := func(x, y int) bool {

		// если вышли за пределы карты то появляемся в противоположеной стороны
		// todo но чет это не работает
		if mp.XSize < x {
			x = x - mp.XSize
		}
		if mp.YSize < y {
			y = y - mp.YSize
		}

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
			// если вышли за пределы карты то появляемся в противоположеной стороны
			// todo но чет это не работает
			x, y := terrainUnit.X, terrainUnit.Y
			if mp.XSize < terrainUnit.X {
				x = terrainUnit.X - mp.XSize
			}
			if mp.YSize < terrainUnit.Y {
				y = terrainUnit.Y - mp.YSize
			}
			createPlant(x, y, "plant_4", mp)
		}

		if terrainUnit.life && !terrainUnit.futureLife {
			mp.RemoveDynamicObjectByXY(terrainUnit.X, terrainUnit.Y)
		}
	}
}

func growTerrainWorker(mp *_map.Map) {
	for {

		growCycle := 30
		growTime := 1000

		for {

			time.Sleep(time.Millisecond * time.Duration(growTime))
			growCycle--

			for _, x := range mp.GetCopyMapDynamicObjects() {
				for _, obj := range x {
					// радар игроков сам обновит обьект
					if obj.Scale < obj.MaxScale {
						obj.Scale++
						obj.CalculateScale()
						obj.GrowTime = growTime
					}
				}
			}

			if growCycle == 0 {
				break
			}
		}

		Evolution(mp)
	}
}
