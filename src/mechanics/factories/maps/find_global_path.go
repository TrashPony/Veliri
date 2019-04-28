package maps

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"math"
)

type SearchMap struct {
	ID     int
	Map    *_map.ShortInfoMap
	F      float64
	Parent *SearchMap
}

// поиск по графам, нормально не отдебажан

func (m *mapStore) FindGlobalPath(startSectorID, endSectorID int) ([]*SearchMap, []*coordinate.Coordinate) { // возращает ячейки пеереходов из сектора в сектор
	startSector, _ := m.GetByID(startSectorID)
	endSector, _ := m.GetByID(endSectorID)

	if startSector == nil || endSector == nil {
		return nil, nil
	}

	start := &SearchMap{ID: startSector.Id, Map: startSector.GetShortInfoMap()}
	end := &SearchMap{ID: endSector.Id, Map: endSector.GetShortInfoMap()}

	openPoints, closePoints := make(map[int]*SearchMap), make(map[int]*SearchMap) // создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints[start.ID] = start                                                  // кладем в карту посещенных точек стартовую точку

	var path []*SearchMap
	var noSortedPath []*SearchMap

	// перменная добавляет в цену номер волны
	wall := 0

	for {

		if len(openPoints) == 0 {
			return nil, nil
		}
		wall++
		current := getOpenPoint(openPoints) // Берем точку с мин стоимостью пути
		if current.ID == end.ID {           // если текущая точка и есть конец начинаем генерить путь
			for !(current.ID == start.ID) {
				current = current.Parent
				if !(current.ID == start.ID) {
					// если текущая точка попрежнему не стартовая то добавляем ее в путь
					noSortedPath = append(noSortedPath, current)
				}
			}
			break
		}
		parseNeighbours(current, openPoints, closePoints, m, end, wall)
	}

	// сразу добавим в путь стартовую точку т.к. нам с нее нужен будет переход
	path = append(path, start)
	for i := len(noSortedPath); i > 0; i-- {
		path = append(path, noSortedPath[i-1])
	}
	// и послуюднюю что бы знать куда прыгать в конце пути
	path = append(path, end)

	var transitionPoints []*coordinate.Coordinate
	for i := 0; i < len(path); i++ {
		if i+1 < len(path) {
			transitionPoint := path[i].Map.GetEntryTySector(path[i+1].ID)
			transitionPoint.MapID = path[i].Map.Id
			transitionPoints = append(transitionPoints, transitionPoint)
		}
	}
	return path, transitionPoints
}

func parseNeighbours(current *SearchMap, openPoints, closePoints map[int]*SearchMap, m *mapStore, end *SearchMap, wall int) {

	delete(openPoints, current.ID)    // удаляем ячейку из не посещенных
	closePoints[current.ID] = current // добавляем в массив посещенные

	// надо взять все переходы и это будут соседи
	entrySectors := current.Map.GetAllEntrySectors()

	for _, entry := range entrySectors {
		mp, _ := m.GetByID(entry.ToMapID)
		// проверяем что карта существует, и что мы ее уже не обработали
		if mp == nil || closePoints[mp.Id] != nil || openPoints[mp.Id] != nil {
			continue
		}

		openPoints[mp.Id] = &SearchMap{
			ID:     mp.Id,
			Map:    mp.GetShortInfoMap(),
			Parent: current,
		}
		// добавяем номер волны что бы все последующие волны имели большую цену
		openPoints[mp.Id].F = getF(openPoints[mp.Id], end) + float64(wall)
	}
}

func getF(curr, end *SearchMap) float64 {
	tmp := math.Abs(float64(curr.Map.XGlobal - end.Map.XGlobal)) // вычисляем разницу между точкой и концом пути по Х
	tmp += math.Abs(float64(curr.Map.YGlobal - end.Map.YGlobal)) // вычисляем разницу между точкой и концом пути по Y и сумируем с раницой по X
	return tmp
}

func getOpenPoint(openMaps map[int]*SearchMap) *SearchMap {
	maxF := 0.0

	var minMap *SearchMap
	var safeMinMap *SearchMap

	for _, p := range openMaps {
		if p.F > maxF {
			maxF = p.F
		}
	}

	for _, p := range openMaps {
		if p.F < maxF {
			minMap = p
		}
		if p.F == maxF {
			safeMinMap = p
		}
	}

	// иногда может не быть карты с меньшей стоимостью пути
	if minMap != nil {
		return minMap
	} else {
		return safeMinMap
	}
}
