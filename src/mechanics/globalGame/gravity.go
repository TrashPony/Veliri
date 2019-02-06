package globalGame

import "github.com/TrashPony/Veliri/src/mechanics/factories/bases"

func GetGravity(x, y, mapID int) bool { // если тру то highGravity
	mapBases := bases.Bases.GetBasesByMap(mapID)

	for _, mapBase := range mapBases {
		xBase, yBase := GetXYCenterHex(mapBase.Q, mapBase.R)

		dist := GetBetweenDist(x, y, xBase, yBase)
		if int(dist) < mapBase.GravityRadius {
			return false
		}
	}

	return true
}
