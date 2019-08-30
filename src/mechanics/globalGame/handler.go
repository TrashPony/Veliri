package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/gorilla/websocket"
)

func HandlerDetect(moveUnit *unit.Unit) *coordinate.Coordinate {
	mp, _ := maps.Maps.GetByID(moveUnit.MapID)

	for _, coor := range mp.HandlersCoordinates {
		xHandle, yHandle := GetXYCenterHex(coor.Q, coor.R)
		dist := int(GetBetweenDist(moveUnit.X, moveUnit.Y, xHandle, yHandle))
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
			if user.GetSquad() != nil && coor.ToMapID == user.GetSquad().MatherShip.MapID {
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
