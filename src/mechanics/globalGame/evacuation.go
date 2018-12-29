package globalGame

import (
	"errors"
	"../player"
	"../gameObjects/map"
	"../factories/bases"
	"../gameObjects/base"
	"sync"
)

func LaunchEvacuation(user *player.Player, mp *_map.Map) ([]PathUnit, int, *base.Transport, error) {

	mapBases := bases.Bases.GetBasesByMap(mp.Id)
	minDist := 0.0
	evacuationBase := &base.Base{}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	for _, mapBase := range mapBases {

		x, y := GetXYCenterHex(mapBase.Q, mapBase.R)
		dist := GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
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

			_, path := MoveTo(float64(startX), float64(startY), 250, 15, 15,
				float64(user.GetSquad().GlobalX), float64(user.GetSquad().GlobalY), 0, mp,
				true, nil, false, false)

			return path, evacuationBase.ID, transport, nil
		} else {
			return nil, 0, nil, errors.New("no available transport")
		}
	} else {
		return nil, 0, nil, errors.New("no available base")
	}
}

func ReturnEvacuation(user *player.Player, mp *_map.Map, baseID int) []PathUnit {
	mapBase, _ := bases.Bases.Get(baseID)
	endX, endY := GetXYCenterHex(mapBase.Q, mapBase.R)

	_, path := MoveTo(float64(user.GetSquad().GlobalX), float64(user.GetSquad().GlobalY), 250, 15, 15,
		float64(endX), float64(endY), 0, mp, true, nil, false, false)
	return path
}
