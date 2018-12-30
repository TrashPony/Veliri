package maps

import (
	"../../gameObjects/map"
	"../../gameObjects/resource"
	"../gameTypes"
	"math/rand"
	"time"
)

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
		rand.Seed(time.Now().UnixNano())

		q := rand.Intn(mp.QSize)
		r := rand.Intn(mp.RSize)

		coordinate, _ := mp.GetCoordinate(q, r)
		_, findRes := mp.Reservoir[q][r]

		if !findRes && coordinate.Move && coordinate.AnimateSpriteSheets == "" && coordinate.TextureObject == "" {
			i++
			newRes, _ := gameTypes.Resource.GetMapReservoirByID(typeRes.TypeID)

			newRes.Q = q
			newRes.R = r
			newRes.Rotate = rand.Intn(360)
			newRes.MapID = mp.Id

			mp.OneLayerMap[coordinate.Q][coordinate.R].Move = false // т.к. на координате ресурс то координата не проходима

			addResourceInMap(mp, newRes)
		}
	}
}

func addResourceInMap(mp *_map.Map, reservoir *resource.Map) {
	if mp.Reservoir[reservoir.Q] != nil {
		mp.Reservoir[reservoir.Q][reservoir.R] = reservoir
	} else {
		mp.Reservoir[reservoir.Q] = make(map[int]*resource.Map)
		mp.Reservoir[reservoir.Q][reservoir.R] = reservoir
	}
}
