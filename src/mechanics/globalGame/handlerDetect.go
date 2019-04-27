package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func HandlerDetect(moveUser *player.Player) *coordinate.Coordinate {
	mp, _ := maps.Maps.GetByID(moveUser.GetSquad().MapID)

	for _, coor := range mp.HandlersCoordinates {
		xHandle, yHandle := GetXYCenterHex(coor.Q, coor.R)
		dist := int(GetBetweenDist(moveUser.GetSquad().GlobalX, moveUser.GetSquad().GlobalY, xHandle, yHandle))
		if dist < 60 && coor.Handler != "" {
			return coor
		}
	}
	return nil
}
