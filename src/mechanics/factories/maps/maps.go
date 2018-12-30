package maps

import (
	"../../db/get"
	"../../gameObjects/coordinate"
	"../../gameObjects/map"
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
