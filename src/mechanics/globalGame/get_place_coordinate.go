package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func GetPlaceCoordinate(placeUnit *unit.Unit, units map[int]*unit.ShortUnitInfo) {

	// TODO
	//mp, _ := maps.Maps.GetByID(placeUnit.MapID)

	//if placeUnit.X == 0 && placeUnit.Y == 0 {
	//	x, y := game_math.GetXYCenterHex(placeUnit.X, placeUnit.Y)
	//	placeUnit.X = x
	//	placeUnit.Y = y
	//
	//	placeUnit.ToX = float64(x)
	//	placeUnit.ToY = float64(y)
	//
	//	placeUnit.CurrentSpeed = 0
	//}
	//
	//findPlace := false
	//for _, gameUnit := range units {
	//	if gameUnit.ID != placeUnit.ID && !placeUnit.InSky {
	//
	//		dist := game_math.GetBetweenDist(gameUnit.X, gameUnit.Y, placeUnit.X, placeUnit.Y)
	//
	//		if dist < 80 {
	//			findPlace = true
	//		}
	//	}
	//}
	//
	//if findPlace {
	//	resp, _ := mp.GetCoordinate(placeUnit.X, placeUnit.Y)
	//	respCoordinates := coordinate.GetCoordinatesRadius(resp, 2)
	//
	//	for _, respFakeCoordinate := range respCoordinates {
	//
	//		respCoordinate, ok := mp.GetCoordinate(respFakeCoordinate.Q, respFakeCoordinate.R)
	//
	//		if ok && respCoordinate.Move {
	//			x, y := game_math.GetXYCenterHex(respCoordinate.Q, respCoordinate.R)
	//			find := false
	//
	//			for _, gameUnit := range units {
	//				dist := game_math.GetBetweenDist(gameUnit.X, gameUnit.Y, x, y)
	//				if dist < 80 && !placeUnit.InSky {
	//					find = true
	//				}
	//			}
	//
	//			if !find {
	//				placeUnit.X = x
	//				placeUnit.Y = y
	//				break
	//			}
	//		}
	//	}
	//}
}
