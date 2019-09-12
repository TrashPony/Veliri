package generators

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
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

		q := rand.Intn(mp.QSize)
		r := rand.Intn(mp.RSize)

		coordinate, _ := mp.GetCoordinate(q, r)
		_, findRes := mp.Reservoir[q][r]

		if !findRes && coordinate.Move && coordinate.AnimateSpriteSheets == "" && coordinate.TextureObject == "" && checkPlace(mp, q, r) {
			i++
			newRes, _ := gameTypes.Resource.GetMapReservoirByID(typeRes.TypeID)

			newRes.Q = q
			newRes.R = r
			newRes.Rotate = rand.Intn(360)
			newRes.MapID = mp.Id

			if !newRes.Move() {
				mp.OneLayerMap[coordinate.Q][coordinate.R].Move = false // т.к. на координате ресурс то координата не проходима
			}

			mp.AddResourceInMap(newRes)
		}
	}
}

func checkPlace(mp *_map.Map, q, r int) bool {

	// ресурсы должны быть дальше на 150 px от
	// баз
	// респанов баз
	// хендлеров

	minDist := 150.0

	game_math.GetXYCenterHex(q, r)

	x, y := game_math.GetXYCenterHex(q, r)

	for _, base := range bases.Bases.GetBasesByMap(mp.Id) {
		baseX, baseY := game_math.GetXYCenterHex(base.Q, base.R)
		if game_math.GetBetweenDist(x, y, baseX, baseY) < 350 {
			return false
		}
	}

	for _, handler := range mp.HandlersCoordinates {
		handlerX, handlerY := game_math.GetXYCenterHex(handler.Q, handler.R)
		if game_math.GetBetweenDist(x, y, handlerX, handlerY) < minDist {
			return false
		}

	}

	entryPoints := maps.Maps.GetEntryPointsByMapID(mp.Id)
	for _, exit := range entryPoints {
		handlerX, handlerY := game_math.GetXYCenterHex(exit.Q, exit.R)
		if game_math.GetBetweenDist(x, y, handlerX, handlerY) < minDist {
			return false
		}
	}

	for _, geoPoint := range mp.GeoData {
		if game_math.GetBetweenDist(x, y, geoPoint.X, geoPoint.Y) < float64(30+geoPoint.Radius) {
			return false
		}
	}

	for _, xLine := range mp.OneLayerMap {
		for _, coordinate := range xLine {
			if strings.Contains(coordinate.TextureObject, "road") {
				coordinateX, coordinateY := game_math.GetXYCenterHex(coordinate.Q, coordinate.R)
				if game_math.GetBetweenDist(x, y, coordinateX, coordinateY) < minDist {
					return false
				}
			}
		}
	}

	return true
}
