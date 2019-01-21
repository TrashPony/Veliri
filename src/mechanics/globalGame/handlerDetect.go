package globalGame

import (
	"../factories/maps"
	"../player"
)

func HandlerDetect(moveUser *player.Player, mapID int) {
	mp, _ := maps.Maps.GetByID(mapID)

	for _, coordinate := range mp.HandlersCoordinates {
		xHandle, yHandle := GetXYCenterHex(coordinate.Q, coordinate.R)
		dist := int(GetBetweenDist(moveUser.GetSquad().GlobalX, moveUser.GetSquad().GlobalY, xHandle, yHandle))
		if dist < 60 {
			if coordinate.Handler == "base" {

			}
			if coordinate.Handler == "sector" {

			}
		}
	}
}
