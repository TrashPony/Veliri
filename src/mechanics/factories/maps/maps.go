package maps

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/anomaly"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/getlantern/deepcopy"
	"math/rand"
)

type mapStore struct {
	maps    map[int]*_map.Map
	anomaly map[int][]*anomaly.Anomaly
}

var Maps = newMapStore()

func newMapStore() *mapStore {
	m := &mapStore{maps: get.Maps(), anomaly: make(map[int][]*anomaly.Anomaly)}

	for id, mp := range m.maps {
		respawns := 0
		mp.HandlersCoordinates = make([]*coordinate.Coordinate, 0)

		for _, q := range mp.OneLayerMap { // считает количество респаунов на карте
			for _, mapCoordinate := range q {
				if mapCoordinate.Type == "respawn" {
					respawns++
				}

				if mapCoordinate.Handler != "" || mapCoordinate.Transport {
					mapCoordinate.HandlerOpen = true
					mp.HandlersCoordinates = append(mp.HandlersCoordinates, mapCoordinate)
				}
			}
		}

		//if mp.Global { // если карта глобальная генерим на ней ресурсы
		//	anomalyGenerator(mp, m) // сначала генерить аномалии что бы можно было использовать больше ячеек
		//	resourceGenerator(mp)
		//}

		mp.Respawns = respawns
		m.maps[id] = mp
	}

	return m
}

func (m *mapStore) GetByID(id int) (*_map.Map, bool) {
	newMap, ok := m.maps[id]
	return newMap, ok
}

func (m *mapStore) GetCopyByID(id int) (*_map.Map, bool) {

	var copyMap _map.Map
	oldMap, ok := m.maps[id]

	err := deepcopy.Copy(&copyMap, &oldMap) // функция глубокого копировния (very slow, but work)
	if err != nil {
		println(err.Error())
	}

	return &copyMap, ok
}

func (m *mapStore) GetAllMap() map[int]*_map.Map {
	return m.maps
}

func (m *mapStore) GetAllShortInfoMap() map[int]*_map.ShortInfoMap {
	shortMap := make(map[int]*_map.ShortInfoMap)
	for _, mp := range m.maps {
		shortMap[mp.Id] = mp.GetShortInfoMap()
	}
	return shortMap
}

func (m *mapStore) GetRespawns(id int) map[int]*coordinate.Coordinate {
	newMap, _ := m.maps[id]
	var respawns = make(map[int]*coordinate.Coordinate)

	for _, q := range newMap.OneLayerMap { // считает количество респаунов на карте
		for _, mapCoordinate := range q {
			if mapCoordinate.Type == "respawn" {
				mapCoordinate.ID = len(respawns) + 1
				respawns[len(respawns)+1] = mapCoordinate
			}
		}
	}

	return respawns
}

func (m *mapStore) GetReservoirByQR(q, r, mapID int) *resource.Map {
	mp, findMap := m.maps[mapID]
	if findMap {
		q, findQ := mp.Reservoir[q]
		if findQ {
			reservoir := q[r]
			return reservoir
		}
	}
	return nil
}

func (m *mapStore) RemoveReservoirByQR(q, r, mapID int) {
	mp, findMap := m.maps[mapID]
	if findMap {
		qLine, findQ := mp.Reservoir[q]
		if findQ {
			delete(qLine, r)
			mp.OneLayerMap[q][r].Move = true
		}
	}
}

func (m *mapStore) GetRandomMap() *_map.Map {
	count := 0
	count2 := rand.Intn(len(m.maps))
	for _, mp := range m.maps {
		if count == count2 {
			return mp
		}
		count++
	}
	return nil
}

func (m *mapStore) GetAllMapAnomaly(mapID int) []*anomaly.Anomaly {
	return Maps.anomaly[mapID]
}

func (m *mapStore) AddNewAnomaly(newAnomaly *anomaly.Anomaly, mapID int) {
	if m.anomaly[mapID] == nil {
		m.anomaly[mapID] = make([]*anomaly.Anomaly, 0)
	}

	m.anomaly[mapID] = append(m.anomaly[mapID], newAnomaly)
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
