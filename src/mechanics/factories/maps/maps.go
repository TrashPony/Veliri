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
	// todo тут можно сразу заполнять инфу например о респаунах
	return &MapStore{maps: get.Maps()}
}

func (m *MapStore) GetByID(id int) (*_map.Map, bool) {
	var newMap _map.Map
	newMap, ok := m.maps[id]
	return &newMap, ok
}

func (m *MapStore) GetInfoMap(id int) _map.Map {
	var newMap _map.Map
	newMap, _ = m.maps[id]

	respawns := 0

	for _, q := range newMap.OneLayerMap { // считает количество респаунов на карте
		for _, mapCoordinate := range q {
			if mapCoordinate.Type == "respawn" {
				respawns++
			}
		}
	}
	newMap.Respawns = respawns

	return newMap
}

func (m *MapStore) GetAllMap() []_map.Map {
	mps := make([]_map.Map, 0)

	for _, mp := range m.maps {
		respawns := 0

		for _, q := range mp.OneLayerMap { // считает количество респаунов на карте
			for _, mapCoordinate := range q {
				if mapCoordinate.Type == "respawn" {
					respawns++
				}
			}
		}

		mp.Respawns = respawns
		mps = append(mps, mp)
	}

	return mps
}

func (m *MapStore) GetRespawns(id int) map[int]*coordinate.Coordinate {
	var newMap _map.Map
	newMap, _ = m.maps[id]
	var respawns = make(map[int]*coordinate.Coordinate)

	for _, q := range newMap.OneLayerMap { // считает количество респаунов на карте
		for _, mapCoordinate := range q {
			if mapCoordinate.Type == "respawn" {
				respawns[len(respawns)+1] = mapCoordinate
			}
		}
	}

	return respawns
}
