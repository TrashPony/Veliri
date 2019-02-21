package maps

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/anomaly"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"math/rand"
)

func anomalyGenerator(mp *_map.Map, m *mapStore) {
	i := 0

	for i < 35 {

		typeAnomaly := rand.Intn(4)

		q := rand.Intn(mp.QSize)
		r := rand.Intn(mp.RSize)
		coordinatePlace, _ := mp.GetCoordinate(q, r)

		if coordinatePlace.Move {
			if m.anomaly[mp.Id] == nil {
				m.anomaly[mp.Id] = make([]*anomaly.Anomaly, 0)
			}

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

			m.anomaly[mp.Id] = append(m.anomaly[mp.Id], anomalyMap)
			i++
		}
	}
}

func (m *mapStore) GetAllMapAnomaly(mapID int) []*anomaly.Anomaly {
	return Maps.anomaly[mapID]
}

func (m *mapStore) GetMapAnomaly(mapID, q, r int) *anomaly.Anomaly {
	for _, anomalyMap := range Maps.anomaly[mapID] {
		if anomalyMap != nil && anomalyMap.GetQ() == q && anomalyMap.GetR() == r {
			return anomalyMap
		}
	}
	return nil
}

func (m *mapStore) RemoveMapAnomaly(mapID, q, r int) {
	for i, anomalyMap := range Maps.anomaly[mapID] {
		if anomalyMap != nil && anomalyMap.GetQ() == q && anomalyMap.GetR() == r {
			Maps.anomaly[mapID][i] = nil
		}
	}
}
