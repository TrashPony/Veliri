package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/gorilla/websocket"
)

func GetPlaceCoordinate(user *player.Player, users map[*websocket.Conn]*player.Player, mp *_map.Map) {

	if user.GetSquad().MatherShip.X == 0 && user.GetSquad().MatherShip.Y == 0 {
		x, y := GetXYCenterHex(user.GetSquad().MatherShip.Q, user.GetSquad().MatherShip.R)
		user.GetSquad().MatherShip.X = x
		user.GetSquad().MatherShip.Y = y

		user.GetSquad().MatherShip.ToX = float64(x)
		user.GetSquad().MatherShip.ToY = float64(y)

		user.GetSquad().MatherShip.CurrentSpeed = 0
	}

	findPlace := false
	for _, gameUser := range users {
		if gameUser.GetSquad() != nil && gameUser.GetID() != user.GetID() && !user.GetSquad().InSky {
			dist := GetBetweenDist(gameUser.GetSquad().MatherShip.X, gameUser.GetSquad().MatherShip.Y,
				user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y)

			if dist < 150 {
				findPlace = true
			}
		}
	}

	if findPlace {
		resp, _ := mp.GetCoordinate(user.GetSquad().MatherShip.Q, user.GetSquad().MatherShip.R)
		respCoordinates := coordinate.GetCoordinatesRadius(resp, 2)

		for _, respFakeCoordinate := range respCoordinates {
			respCoordinate, ok := mp.GetCoordinate(respFakeCoordinate.Q, respFakeCoordinate.R)
			if ok && respCoordinate.Move {
				x, y := GetXYCenterHex(respCoordinate.Q, respCoordinate.R)
				find := false

				for _, gameUser := range users {
					dist := GetBetweenDist(gameUser.GetSquad().MatherShip.X, gameUser.GetSquad().MatherShip.Y, x, y)
					if dist < 150 && !user.GetSquad().InSky {
						find = true
					}
				}

				if !find {
					user.GetSquad().MatherShip.X = x
					user.GetSquad().MatherShip.Y = y
					break
				}
			}
		}
	}
}
