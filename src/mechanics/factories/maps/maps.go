package maps

import (
	dbMap "github.com/TrashPony/Veliri/src/mechanics/db/maps"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/anomaly"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
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
			}
		}

		for _, q := range mp.StaticObjects {
			for _, staticObj := range q {
				// если координата хранит инвентарь то создаем его в карте
				// для статичных обьектов которые существуют всегда и имеют свой ИД инвентарь приемлем
				if staticObj.Inventory {

					box, mx := boxes.Boxes.GetByXY(staticObj.X, staticObj.Y, mp.Id)
					mx.Unlock()

					if box == nil {
						objectBox := &boxInMap.Box{
							MapID:            mp.Id,
							CapacitySize:     100.00,
							Protect:          false,
							X:                staticObj.X,
							Y:                staticObj.Y,
							HP:               -1,
							OwnedByMapObject: true,
						}
						objectBox.GetStorage().Slots = make(map[int]*inventory.Slot)
						objectBox.GetStorage().SetSlotsSize(999)

						staticObj.BoxID = boxes.Boxes.InsertNewBox(objectBox).ID
					} else {
						staticObj.BoxID = box.ID
					}
				}
			}
		}
		m.maps[id] = mp
	}

	// собираем в масив все входы в сектор
	for id, mp := range m.maps {
		mp.EntryPoints = m.GetEntryPointsByMapID(id)
	}

	// TODO парсим названия и описания обьектов на карте
	//for _, mp := range m.maps {
	//	for _, x := range mp.StaticObjects {
	//		for _, mapCoordinate := range x {
	//
	//			if mapCoordinate.Name != "" {
	//				// userName, BaseName, ToBaseName, ToSectorName, userFraction
	//				toMapName := ""
	//				if mapCoordinate.Handler == "sector" {
	//					toMapName = m.maps[mapCoordinate.ToMapID].Name
	//				} else {
	//
	//					// ¯\_(ツ)_/¯
	//
	//					for _, mp2 := range m.maps {
	//						for _, q2 := range mp2.OneLayerMap {
	//							for _, mapCoordinate2 := range q2 {
	//								if mapCoordinate2.Handler == "sector" && mapCoordinate2.ToMapID == mp.Id {
	//									for _, points := range mapCoordinate2.Positions {
	//										if points.X == mapCoordinate.X && points.Y == mapCoordinate.Y {
	//											toMapName = mp2.Name
	//										}
	//									}
	//								}
	//							}
	//						}
	//					}
	//				}
	//				mapCoordinate.Name = dialog.ProcessingText(mapCoordinate.Name, "", "", "", toMapName, "")
	//				mapCoordinate.Description = dialog.ProcessingText(mapCoordinate.Description, "", "", "", toMapName, "")
	//			}
	//		}
	//	}
	//}

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

func (m *mapStore) GetReservoirByXY(x, y, mapID int) *resource.Map {
	mp, findMap := m.maps[mapID]
	if findMap {
		q, findQ := mp.Reservoir[x]
		if findQ {
			reservoir := q[y]
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

func (m *mapStore) GetMapAnomaly(mapID, x, y int) *anomaly.Anomaly {
	for _, anomalyMap := range Maps.anomaly[mapID] {
		if anomalyMap != nil {
			dist := game_math.GetBetweenDist(x, y, anomalyMap.GetX(), anomalyMap.GetY())
			if dist < 150 {
				return anomalyMap
			}
		}
	}
	return nil
}

func (m *mapStore) RemoveMapAnomaly(mapID int, removeAnomaly *anomaly.Anomaly) {
	for i, anomalyMap := range Maps.anomaly[mapID] {
		if anomalyMap != nil && anomalyMap.GetX() == removeAnomaly.GetX() && anomalyMap.GetY() == removeAnomaly.GetY() {
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
