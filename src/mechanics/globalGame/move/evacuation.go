package move

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"sync"
)

func LaunchEvacuation(unit *unit.Unit, mp *_map.Map) ([]*unit.PathUnit, int, *base.Transport, error) {

	mapBases := bases.Bases.GetBasesByMap(mp.Id)
	minDist := 0.0
	evacuationBase := &base.Base{}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	for _, mapBase := range mapBases {

		dist := game_math.GetBetweenDist(unit.X, unit.Y, mapBase.X, mapBase.Y)
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
				startX, startY = evacuationBase.X, evacuationBase.Y
			} else {
				startX, startY = transport.X, transport.Y
			}

			_, path := To(float64(startX), float64(startY), 15, 1, 1,
				float64(unit.X), float64(unit.Y), transport.Rotate, 25)

			return path, evacuationBase.ID, transport, nil
		} else {
			return nil, 0, nil, errors.New("no available transport")
		}
	} else {
		return nil, 0, nil, errors.New("no available base")
	}
}

func ReturnEvacuation(unit *unit.Unit, mp *_map.Map, baseID int, transport *base.Transport) []*unit.PathUnit {
	mapBase, _ := bases.Bases.Get(baseID)

	_, path := To(float64(unit.X), float64(unit.Y), 15, 15, 15,
		float64(mapBase.X), float64(mapBase.Y), transport.Rotate, 10)
	return path
}
