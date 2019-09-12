package move

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func GetGravity(x, y, mapID int) bool { // если тру то highGravity
	mapBases := bases.Bases.GetBasesByMap(mapID)

	for _, mapBase := range mapBases {
		xBase, yBase := game_math.GetXYCenterHex(mapBase.Q, mapBase.R)

		dist := game_math.GetBetweenDist(x, y, xBase, yBase)
		if int(dist) < mapBase.GravityRadius {
			return false
		}
	}

	return true
}
