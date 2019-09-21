package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func checkValidForMoveCoordinate(gameMap *_map.Map, x, y, pX, pY, pRotate int, gameUnit *unit.Unit, scaleMap int,
	allUnits map[int]*unit.ShortUnitInfo) (*coordinate.Coordinate, bool) {

	newCoor := &coordinate.Coordinate{X: x, Y: y, Rotate: game_math.GetBetweenAngle(float64(x), float64(y), float64(pX), float64(pY))}

	if allUnits != nil {
		free, _ := collisions.CheckCollisionsPlayers(gameUnit, x*scaleMap, y*scaleMap, pRotate, allUnits)
		if !free {
			return nil, false
		}
	}

	possible, _ := collisions.CheckCollisionsOnStaticMap(x*scaleMap, y*scaleMap, newCoor.Rotate, gameMap, gameUnit.Body, false, true)

	if possible {
		return newCoor, true
	}

	return nil, false
}

//func addGeoCoordinate(new *coordinate.Coordinate, gameMap *_map.Map, scale int, possible bool) {
//
//	// координаты в мапе нельзя отдавать т.к. они там изменяются и происходит какойто пиздец
//	// TODO обновлять информацию спустя какоето время, т.к. она все же может менятся, например руды для копания
//
//	mx.Lock()
//	defer mx.Unlock()
//	if gameMap.GeoDataMaps == nil {
//		gameMap.GeoDataMaps = make([][][][][]bool, 1000)
//
//	}
//
//	if _, ok := gameMap.GeoDataMaps[scale]; !ok {
//		gameMap.GeoDataMaps[scale] = make(map[int]map[int]coordinate.Coordinate)
//	}
//
//	if _, ok := gameMap.GeoDataMaps[scale][new.X]; !ok {
//		gameMap.GeoDataMaps[scale][new.X] = make(map[int]coordinate.Coordinate)
//	}
//
//	if possible {
//		new.State = FREE
//	} else {
//		new.State = BLOCKED
//	}
//
//	gameMap.GeoDataMaps[scale][new.X][new.Y] = coordinate.Coordinate{X: new.X, Y: new.Y, State: new.State}
//}
