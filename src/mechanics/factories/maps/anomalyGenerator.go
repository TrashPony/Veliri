package maps

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/anomaly"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
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
				/// todo невероятный говнокод но ради фана

				pages := make([]dialog.Page, 0)

				ask1 := make([]dialog.Ask, 0)
				ask1 = append(ask1, dialog.Ask{Text: "Прочитать запись", ToPage: 2})
				ask1 = append(ask1, dialog.Ask{Text: "О-о-о нет я сваливаю", ToPage: 0})
				pages = append(pages, dialog.Page{Text: "Вы находите старый ржавый не на что не похожий информационный пакет, вы попытаись " +
					"подколючится к нему и считать информацию но сходу не удалось расшифровать запись, спустя не" +
					" продолжительное время для человека и целую вечность для машины вы смогли расшифровать информацию " +
					"и удивились тому насколько глубока мысль тех кто оставил этот пакет здесь когда то очень давно. <br><br> " +
					"Информация в пакете гласила:", Asc: ask1})

				ask2 := make([]dialog.Ask, 0)
				ask2 = append(ask2, dialog.Ask{Text: "Понятно", ToPage: 0}) // если 0 то close
				pages = append(pages, dialog.Page{Text: "\"Ты пидор\"", Asc: ask2})

				anomalyMap.SetDialog(&dialog.Dialog{Pages: pages})
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
