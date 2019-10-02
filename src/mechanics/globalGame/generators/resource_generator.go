package generators

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math/rand"
	"strings"
)

func GenerateObjectsMap() {
	for _, mp := range maps.Maps.GetAllMap() {
		if mp.Global {
			resourceGenerator(mp)
			AnomalyGenerator(mp)
			baseInit(mp)
		}
	}
}

func baseInit(mp *_map.Map) {
	// инитим все базы на количество ресурсов в них и расход

	for _, gameBase := range bases.Bases.GetBasesByMap(mp.Id) {

		gameBase.CurrentResources = make(map[int]*inventory.Slot)

		for _, recycled := range gameTypes.Resource.GetAllRecycled() {
			gameBase.CurrentResources[recycled.TypeID] = &inventory.Slot{
				Item:     recycled,
				Quantity: 0,
				Type:     "recycle",
				ItemID:   recycled.TypeID,
			}
		}

		go gameBase.ConsumptionBaseResource()
	}
}

func resourceGenerator(mp *_map.Map) {

	allTypeResource := gameTypes.Resource.GetAllTypeMapResource()

	mp.Reservoir = make(map[int]map[int]*resource.Map)

	for _, typeRes := range allTypeResource {
		generate(mp, typeRes, 5)
	}
}

func generate(mp *_map.Map, typeRes resource.Map, count int) {
	i := 0

	for i < count {

		x := rand.Intn(mp.XSize)
		y := rand.Intn(mp.YSize)

		if checkPlace(mp, x, y) {
			i++
			newRes, _ := gameTypes.Resource.GetMapReservoirByID(typeRes.TypeID)

			newRes.X = x
			newRes.Y = y
			newRes.Rotate = rand.Intn(360)
			newRes.MapID = mp.Id

			if !newRes.Move() {
				coordinateMap, ok := mp.OneLayerMap[newRes.X][newRes.Y]
				if ok {
					coordinateMap.Move = false // т.к. на координате ресурс то координата не проходима
				} else {
					coordinate.AddIntCoordinate(mp.OneLayerMap, &coordinate.Coordinate{Move: false})
				}
			}

			mp.AddResourceInMap(newRes)
		}
	}
}

func checkPlace(mp *_map.Map, x, y int) bool {

	// ресурсы должны быть дальше на 150 px от
	// баз
	// респанов баз
	// хендлеров

	minDist := 150.0

	for _, base := range bases.Bases.GetBasesByMap(mp.Id) {
		if game_math.GetBetweenDist(x, y, base.X, base.Y) < 350 {
			return false
		}
	}

	for _, handler := range mp.HandlersCoordinates {
		if game_math.GetBetweenDist(x, y, handler.X, handler.Y) < minDist {
			return false
		}

	}

	entryPoints := maps.Maps.GetEntryPointsByMapID(mp.Id)
	for _, exit := range entryPoints {
		if game_math.GetBetweenDist(x, y, exit.X, exit.Y) < minDist {
			return false
		}
	}

	for _, geoPoint := range mp.GeoData {
		if game_math.GetBetweenDist(x, y, geoPoint.X, geoPoint.Y) < float64(30+geoPoint.Radius) {
			return false
		}
	}

	for _, xLine := range mp.OneLayerMap {
		for _, coordinateMap := range xLine {
			if strings.Contains(coordinateMap.TextureObject, "road") {
				if game_math.GetBetweenDist(x, y, coordinateMap.X, coordinateMap.Y) < minDist {
					return false
				}
			}
		}
	}

	return true
}
