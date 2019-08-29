package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/gorilla/websocket"
)

func HandlerDetect(moveUser *player.Player) *coordinate.Coordinate {
	mp, _ := maps.Maps.GetByID(moveUser.GetSquad().MapID)

	for _, coor := range mp.HandlersCoordinates {
		xHandle, yHandle := GetXYCenterHex(coor.Q, coor.R)
		dist := int(GetBetweenDist(moveUser.GetSquad().MatherShip.X, moveUser.GetSquad().MatherShip.Y, xHandle, yHandle))
		if dist < 60 && coor.Handler != "" {
			return coor
		}
	}
	return nil
}

func CheckHandlerCoordinate(coor *coordinate.Coordinate, users map[*websocket.Conn]*player.Player) *coordinate.Coordinate {

	// nil все клетки заняты

	for _, exit := range coor.Positions {

		x, y := GetXYCenterHex(exit.Q, exit.R)
		busy := false

		for _, user := range users {
			if user.GetSquad() != nil && coor.ToMapID == user.GetSquad().MapID {
				dist := GetBetweenDist(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, x, y)
				if dist < 135 {
					busy = true
				}
			}
		}

		if !busy {
			return exit
		}
	}

	return nil
}
