package generators

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/anomaly"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"math/rand"
)

func AnomalyGenerator(mp *_map.Map) {

	i := 0

	for i < 35 {

		typeAnomaly := rand.Intn(4)

		q := rand.Intn(mp.QSize)
		r := rand.Intn(mp.RSize)
		coordinatePlace, _ := mp.GetCoordinate(q, r)

		if coordinatePlace.Move {

			anomalyMap := &anomaly.Anomaly{Type: typeAnomaly, MapID: mp.Id}

			anomalyMap.SetQ(q)
			anomalyMap.SetR(r)
			anomalyMap.SetPower(rand.Intn(6))

			// коробка с ресурсами 2+лвл
			if typeAnomaly == 0 {
				anomalyMap.SetBox(boxes.Boxes.GetAnomalyRandomBox(typeAnomaly, gameTypes.Boxes.GetRandomBox()))
			}

			// коробка с чертежом
			if typeAnomaly == 1 {
				anomalyMap.SetBox(boxes.Boxes.GetAnomalyRandomBox(typeAnomaly, gameTypes.Boxes.GetRandomBox()))
			}

			// руда
			if typeAnomaly == 2 {
				anomalyMap.SetRes(gameTypes.Resource.GetRandomMapResource())
			}

			// текст
			if typeAnomaly == 3 {
				// todo TEXT
			}

			maps.Maps.AddNewAnomaly(anomalyMap, mp.Id)
			i++
		}
	}
}
