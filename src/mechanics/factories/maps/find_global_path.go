package maps

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

type SearchMap struct {
	ID     int
	Map    *_map.Map
	F      int
	Parent *SearchMap
}

// поиск по графам, нормально не отдебажан

func (m *mapStore) FindGlobalPath(startSectorID, endSectorID int) ([]*SearchMap, []*coordinate.Coordinate) { // возращает ячейки пеереходов из сектора в сектор
	startSector, _ := m.GetByID(startSectorID)
	endSector, _ := m.GetByID(endSectorID)

	if startSector == nil || endSector == nil {
		return nil, nil
	}

	start := &SearchMap{ID: startSector.Id, Map: startSector}
	end := &SearchMap{ID: endSector.Id, Map: endSector}

	openPoints, closePoints := make(map[int]*SearchMap), make(map[int]*SearchMap) // создаем 2 карты для посещенных (open) и непосещеных (close) точек
	openPoints[start.ID] = start                                                  // кладем в карту посещенных точек стартовую точку

	var path []*SearchMap
	var noSortedPath []*SearchMap

	for {

		if len(openPoints) <= 0 {
			return nil, nil
		}

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

		delete(openPoints, current.ID)    // удаляем ячейку из не посещенных
		closePoints[current.ID] = current // добавляем в массив посещенные

		// надо взять все переходы и это будут соседи
		entrySectors := current.Map.GetAllEntrySectors()
		for _, entry := range entrySectors {
			mp, _ := m.GetByID(entry.ToMapID)
			// проверяем что карта существует, и что мы ее уже не обработали
			if mp == nil && closePoints[mp.Id] == nil && openPoints[mp.Id] == nil {
				continue
			}
			//TODO fatal error: out of memory когда есть сектор к оторый нет перехода
			openPoints[mp.Id] = &SearchMap{
				ID:     mp.Id,
				Map:    mp,
				F:      current.F + 1, // т.к. все равноценно то цена пути всегда 1.
				Parent: current,
			}
		}
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

func getOpenPoint(openMaps map[int]*SearchMap) *SearchMap {
	maxF := 0
	var minMap *SearchMap
	for _, p := range openMaps {
		if p.F > maxF {
			maxF = p.F
		}
	}

	for _, p := range openMaps {
		if p.F <= maxF {
			minMap = p
		}
	}
	return minMap
}
