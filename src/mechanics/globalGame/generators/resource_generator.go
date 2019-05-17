package generators

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"math/rand"
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

		go gameBase.ConsumptionBaseResource()

		for _, recycled := range gameTypes.Resource.GetAllRecycled() {
			gameBase.CurrentResources[recycled.TypeID] = &inventory.Slot{
				Item:     recycled,
				Quantity: 0,
				Type:     "recycle",
				ItemID:   recycled.TypeID,
			}
		}
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

	// ресурсы должны быть дальше на 550 px от
	// баз
	// респанов баз
	// хендлеров

	minDist := 550.0

	globalGame.GetXYCenterHex(q, r)

	x, y := globalGame.GetXYCenterHex(q, r)

	for _, base := range bases.Bases.GetBasesByMap(mp.Id) {
		baseX, baseY := globalGame.GetXYCenterHex(base.Q, base.R)
		if globalGame.GetBetweenDist(x, y, baseX, baseY) < minDist {
			return false
		}
	}

	for _, handler := range mp.HandlersCoordinates {
		handlerX, handlerY := globalGame.GetXYCenterHex(handler.Q, handler.R)
		if globalGame.GetBetweenDist(x, y, handlerX, handlerY) < minDist {
			return false
		}

		if handler.Handler == "sector" {
			handlerX, handlerY = globalGame.GetXYCenterHex(handler.ToQ, handler.ToR)
			if globalGame.GetBetweenDist(x, y, handlerX, handlerY) < minDist {
				return false
			}
		}
	}

	return true
}
