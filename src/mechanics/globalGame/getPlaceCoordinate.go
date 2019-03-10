package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
)

func GetPlaceCoordinate(user *player.Player, users map[*websocket.Conn]*player.Player, mp *_map.Map) {

	if user.GetSquad().GlobalX == 0 && user.GetSquad().GlobalY == 0 {
		x, y := GetXYCenterHex(user.GetSquad().Q, user.GetSquad().R)
		user.GetSquad().GlobalX = x
		user.GetSquad().GlobalY = y

		user.GetSquad().ToX = float64(x)
		user.GetSquad().ToY = float64(y)

		user.GetSquad().CurrentSpeed = 0
	}

	findPlace := false
	for _, gameUser := range users {
		if gameUser.GetID() != user.GetID() && !user.GetSquad().InSky {
			dist := GetBetweenDist(gameUser.GetSquad().GlobalX, gameUser.GetSquad().GlobalY,
				user.GetSquad().GlobalX, user.GetSquad().GlobalY)

			if dist < 150 {
				findPlace = true
			}
		}
	}

	if findPlace {
		resp, _ := mp.GetCoordinate(user.GetSquad().Q, user.GetSquad().R)
		respCoordinates := coordinate.GetCoordinatesRadius(resp, 2)

		for _, respFakeCoordinate := range respCoordinates {
			respCoordinate, ok := mp.GetCoordinate(respFakeCoordinate.Q, respFakeCoordinate.R)
			if ok && respCoordinate.Move {
				x, y := GetXYCenterHex(respCoordinate.Q, respCoordinate.R)
				find := false

				for _, gameUser := range users {
					dist := GetBetweenDist(gameUser.GetSquad().GlobalX, gameUser.GetSquad().GlobalY, x, y)
					if dist < 150 && !user.GetSquad().InSky {
						find = true
					}
				}

				if !find {
					user.GetSquad().GlobalX = x
					user.GetSquad().GlobalY = y
					break
				}
			}
		}
	}
}
