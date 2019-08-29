package globalGame

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"sync"
)

func LaunchEvacuation(user *player.Player, mp *_map.Map) ([]unit.PathUnit, int, *base.Transport, error) {

	mapBases := bases.Bases.GetBasesByMap(mp.Id)
	minDist := 0.0
	evacuationBase := &base.Base{}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	for _, mapBase := range mapBases {

		x, y := GetXYCenterHex(mapBase.Q, mapBase.R)
		dist := GetBetweenDist(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, x, y)
		transport := mapBase.GetFreeTransport()

		if ((dist < minDist && int(dist) < mapBase.GravityRadius) ||
			(minDist == 0 && int(dist) < mapBase.GravityRadius)) && transport != nil {
			minDist = dist
			evacuationBase = mapBase
		}
	}

	if evacuationBase != nil {
		transport := evacuationBase.GetFreeTransport()
		if transport != nil {
			var startX, startY int

			transport.Job = true

			if transport.X == 0 && transport.Y == 0 {
				startX, startY = GetXYCenterHex(evacuationBase.Q, evacuationBase.R)
			} else {
				startX = transport.X
				startY = transport.Y
			}

			_, path := MoveTo(float64(startX), float64(startY), 15, 15, 15,
				float64(user.GetSquad().MatherShip.X), float64(user.GetSquad().MatherShip.Y), 0, mp,
				true, nil, false, false, nil)

			return path, evacuationBase.ID, transport, nil
		} else {
			return nil, 0, nil, errors.New("no available transport")
		}
	} else {
		return nil, 0, nil, errors.New("no available base")
	}
}

func ReturnEvacuation(user *player.Player, mp *_map.Map, baseID int) []unit.PathUnit {
	mapBase, _ := bases.Bases.Get(baseID)
	endX, endY := GetXYCenterHex(mapBase.Q, mapBase.R)

	_, path := MoveTo(float64(user.GetSquad().MatherShip.X), float64(user.GetSquad().MatherShip.Y), 250, 15, 15,
		float64(endX), float64(endY), 0, mp, true, nil, false, false, nil)
	return path
}
