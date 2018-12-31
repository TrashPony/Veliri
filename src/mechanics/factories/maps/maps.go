package maps

import (
	"../../db/get"
	"../../gameObjects/coordinate"
	"../../gameObjects/map"
	"../../gameObjects/resource"
)

type MapStore struct {
	maps map[int]_map.Map
}

var Maps = NewMapStore()

func NewMapStore() *MapStore {
	m := &MapStore{maps: get.Maps()}

	for id, mp := range m.maps {
		respawns := 0
		for _, q := range mp.OneLayerMap { // считает количество респаунов на карте
			for _, mapCoordinate := range q {
				if mapCoordinate.Type == "respawn" {
					respawns++
				}
			}
		}

		if mp.Global { // если карта глобальная генерим на ней ресурсы
			resourceGenerator(&mp)
		}

		mp.Respawns = respawns
		m.maps[id] = mp
	}

	return m
}

func (m *MapStore) GetByID(id int) (*_map.Map, bool) {
	var newMap _map.Map
	newMap, ok := m.maps[id]
	return &newMap, ok
}

func (m *MapStore) GetAllMap() map[int]_map.Map {
	return m.maps
}

func (m *MapStore) GetRespawns(id int) map[int]*coordinate.Coordinate {
	var newMap _map.Map
	newMap, _ = m.maps[id]
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

func (m *MapStore) GetReservoirByQR(q, r, mapID int) *resource.Map {
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

func (m *MapStore) RemoveReservoirByQR(q, r, mapID int) {
	mp, findMap := m.maps[mapID]
	if findMap {
		qLine, findQ := mp.Reservoir[q]
		if findQ {
			delete(qLine, r)
			mp.OneLayerMap[q][r].Move = true
		}
	}
}
