package move

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func GetGravity(x, y, mapID int) bool { // если тру то highGravity
	mapBases := bases.Bases.GetBasesByMap(mapID)

	for _, mapBase := range mapBases {
		dist := game_math.GetBetweenDist(x, y, mapBase.X, mapBase.Y)
		if int(dist) < mapBase.GravityRadius {
			return false
		}
	}

	return true
}
