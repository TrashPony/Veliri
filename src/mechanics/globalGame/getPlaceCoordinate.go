package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/gorilla/websocket"
)

func GetPlaceCoordinate(placeUnit unit.Unit, units map[*websocket.Conn]*unit.ShortUnitInfo, mp *_map.Map) {

	if placeUnit.X == 0 && placeUnit.Y == 0 {
		x, y := GetXYCenterHex(placeUnit.Q, placeUnit.R)
		placeUnit.X = x
		placeUnit.Y = y

		placeUnit.ToX = float64(x)
		placeUnit.ToY = float64(y)

		placeUnit.CurrentSpeed = 0
	}

	findPlace := false
	for _, gameUnit := range units {
		if gameUnit.ID != placeUnit.ID && !placeUnit.InSky {

			dist := GetBetweenDist(gameUnit.X, gameUnit.Y, placeUnit.X, placeUnit.Y)

			if dist < 150 {
				findPlace = true
			}
		}
	}

	if findPlace {
		resp, _ := mp.GetCoordinate(placeUnit.Q, placeUnit.R)
		respCoordinates := coordinate.GetCoordinatesRadius(resp, 2)

		for _, respFakeCoordinate := range respCoordinates {
			respCoordinate, ok := mp.GetCoordinate(respFakeCoordinate.Q, respFakeCoordinate.R)
			if ok && respCoordinate.Move {
				x, y := GetXYCenterHex(respCoordinate.Q, respCoordinate.R)
				find := false

				for _, gameUnit := range units {
					dist := GetBetweenDist(gameUnit.X, gameUnit.Y, x, y)
					if dist < 150 && !placeUnit.InSky {
						find = true
					}
				}

				if !find {
					placeUnit.X = x
					placeUnit.Y = y
					break
				}
			}
		}
	}
}
