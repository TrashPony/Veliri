package maps

import (
	dbMap "github.com/TrashPony/Veliri/src/mechanics/db/maps"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/anomaly"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
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
	m := &mapStore{
		maps:    dbMap.Maps(),
		anomaly: make(map[int][]*anomaly.Anomaly),
	}

	for id, mp := range m.maps {
		mp.HandlersCoordinates = make([]*coordinate.Coordinate, 0)

		for _, q := range mp.OneLayerMap {
			for _, mapCoordinate := range q {
				// переносим координаты со слушателями в отдельный масив для удобного доступа
				if mapCoordinate.Handler != "" || mapCoordinate.Transport {
					mapCoordinate.HandlerOpen = true
					mp.HandlersCoordinates = append(mp.HandlersCoordinates, mapCoordinate)
				}

				// если координата хранит инвентарь то создаем его в карте
				if mapCoordinate.ObjectInventory {

					box, mx := boxes.Boxes.GetByQR(mapCoordinate.Q, mapCoordinate.R, mp.Id)
					mx.Unlock()

					if box == nil {
						objectBox := &boxInMap.Box{
							MapID:            mp.Id,
							CapacitySize:     100.00,
							Protect:          false,
							Q:                mapCoordinate.Q,
							R:                mapCoordinate.R,
							HP:               -1,
							OwnedByMapObject: true,
						}
						objectBox.GetStorage().Slots = make(map[int]*inventory.Slot)
						objectBox.GetStorage().SetSlotsSize(999)

						mapCoordinate.BoxID = boxes.Boxes.InsertNewBox(objectBox).ID
					} else {
						mapCoordinate.BoxID = box.ID
					}
				}
			}
		}
		m.maps[id] = mp
	}

	for _, mp := range m.maps {
		for _, q := range mp.OneLayerMap {
			for _, mapCoordinate := range q {
				if mapCoordinate.ObjectName != "" {
					// userName, BaseName, ToBaseName, ToSectorName, userFraction
					toMapName := ""
					if mapCoordinate.Handler == "sector" {
						toMapName = m.maps[mapCoordinate.ToMapID].Name
					} else {

						// ¯\_(ツ)_/¯

						for _, mp2 := range m.maps {
							for _, q2 := range mp2.OneLayerMap {
								for _, mapCoordinate2 := range q2 {
									if mapCoordinate2.Handler == "sector" && mapCoordinate2.ToMapID == mp.Id {
										for _, points := range mapCoordinate2.Positions {
											if points.Q == mapCoordinate.Q && points.R == mapCoordinate.R {
												toMapName = mp2.Name
											}
										}
									}
								}
							}
						}
					}
					mapCoordinate.ObjectName = dialog.ProcessingText(mapCoordinate.ObjectName, "", "", "", toMapName, "")
					mapCoordinate.ObjectDescription = dialog.ProcessingText(mapCoordinate.ObjectDescription, "", "", "", toMapName, "")
				}
			}
		}
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

func (m *mapStore) GetEntryPointsByMapID(id int) []*coordinate.Coordinate {
	entryPoints := make([]*coordinate.Coordinate, 0)

	for _, mp := range m.maps {
		for _, handler := range mp.HandlersCoordinates {
			if handler.Handler == "sector" && handler.ToMapID == id {
				entryPoints = append(entryPoints, handler)
			}
		}
	}

	return entryPoints
}
