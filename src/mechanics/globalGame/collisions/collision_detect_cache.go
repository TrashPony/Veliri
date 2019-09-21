package collisions

import "sync"

// тут хранятся просчитанные шаблоны координат, что бы не проверять координаты при каждом поиске пути
// mapID, bodyID, scale, x, y
var GeoDataMapsCache = createCacheStore()
var mx sync.Mutex

func createCacheStore() map[int]map[int]map[int]map[int]bool {
	return make(map[int]map[int]map[int]map[int]bool, 1000)
}

func addCacheCoordinate(mapId, bodyID, x, y int, passed bool) {
	mx.Lock()
	defer mx.Unlock()

	if GeoDataMapsCache[mapId] == nil {
		GeoDataMapsCache[mapId] = make(map[int]map[int]map[int]bool)
	}

	if GeoDataMapsCache[mapId][bodyID] == nil {
		GeoDataMapsCache[mapId][bodyID] = make(map[int]map[int]bool)
	}

	if GeoDataMapsCache[mapId][bodyID][x] == nil {
		GeoDataMapsCache[mapId][bodyID][x] = make(map[int]bool)
	}

	GeoDataMapsCache[mapId][bodyID][x][y] = passed
}

func getCacheCoordinate(mapId, bodyID, x, y int) (bool, bool) {
	mx.Lock()
	defer mx.Unlock()

	v, ok := GeoDataMapsCache[mapId][bodyID][x][y]

	if ok {
		return true, v
	} else {
		return false, false
	}
}
