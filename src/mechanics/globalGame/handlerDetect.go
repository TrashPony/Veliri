package globalGame

import (
	"../factories/maps"
	"../gameObjects/coordinate"
	"../player"
)

func HandlerDetect(moveUser *player.Player) *coordinate.Coordinate {
	mp, _ := maps.Maps.GetByID(moveUser.GetSquad().MapID)

	for _, coor := range mp.HandlersCoordinates {
		xHandle, yHandle := GetXYCenterHex(coor.Q, coor.R)
		dist := int(GetBetweenDist(moveUser.GetSquad().GlobalX, moveUser.GetSquad().GlobalY, xHandle, yHandle))
		if dist < 20 && coor.Handler != "" {
			return coor
		}
	}
	return nil
}