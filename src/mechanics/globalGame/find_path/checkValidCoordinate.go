package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"math"
	"sync"
)

var mx sync.Mutex

func checkValidForMoveCoordinate(gameMap *_map.Map, x, y, pX, pY, pRotate int, gameUnit *unit.Unit, scaleMap int,
	allUnits map[int]*unit.ShortUnitInfo, end *coordinate.Coordinate) (*coordinate.Coordinate, bool) {

	mx.Lock()
	geoCoordinate, ok := gameMap.GeoDataMaps[scaleMap][x][y]
	mx.Unlock()

	//needRad := math.Atan2(float64(end.Y-y), float64(end.X-x))
	//needAngle := int(needRad * 180 / 3.14)
	//diffRotate := pRotate - needAngle
	//if diffRotate < 0 {
	//	diffRotate = 360 - diffRotate
	//}
	//
	//if diffRotate != 0 { // если разница есть то поворачиваем корпус
	//
	//}

	needRad := math.Atan2(float64(pY-y), float64(pX-x))
	needAngle := int(needRad * 180 / 3.14)
	diffRotate := 0

	if pRotate > needAngle {
		diffRotate = pRotate - needAngle
	} else {
		diffRotate = needAngle - pRotate
	}

	if diffRotate > 180 {
		return nil, false
	}

	newCoor := &coordinate.Coordinate{X: x, Y: y, Rotate: needAngle}

	free, _ := globalGame.CheckCollisionsPlayers(gameUnit, x*scaleMap, y*scaleMap, pRotate, allUnits)
	if !free {
		return nil, false
	}

	if ok {
		if geoCoordinate.State != BLOCKED {
			return newCoor, true
		}
	} else {
		possible, _, _, _ := globalGame.CheckCollisionsOnStaticMap(x*scaleMap, y*scaleMap, pRotate, gameMap, gameUnit.Body, true)

		addGeoCoordinate(newCoor, gameMap, scaleMap, possible)

		if possible {
			return newCoor, true
		}
	}

	return nil, false
}

func addGeoCoordinate(new *coordinate.Coordinate, gameMap *_map.Map, scale int, possible bool) {
	// координаты в мапе нельзя отдавать т.к. они там изменяются и происходит какойто пиздец
	// TODO обновлять информацию спустя какоето время, т.к. орна все же может менятся, например руды для копания
	mx.Lock()
	defer mx.Unlock()
	if gameMap.GeoDataMaps == nil {
		gameMap.GeoDataMaps = make(map[int]map[int]map[int]coordinate.Coordinate)
	}

	if _, ok := gameMap.GeoDataMaps[scale]; !ok {
		gameMap.GeoDataMaps[scale] = make(map[int]map[int]coordinate.Coordinate)
	}

	if _, ok := gameMap.GeoDataMaps[scale][new.X]; !ok {
		gameMap.GeoDataMaps[scale][new.X] = make(map[int]coordinate.Coordinate)
	}

	if possible {
		new.State = FREE
	} else {
		new.State = BLOCKED
	}
	gameMap.GeoDataMaps[scale][new.X][new.Y] = coordinate.Coordinate{X: new.X, Y: new.Y, State: new.State}
}
