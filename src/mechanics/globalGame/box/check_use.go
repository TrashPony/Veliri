package box

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"sync"
)

func checkUseBox(user *player.Player, boxID int) (error, *boxInMap.Box, *sync.Mutex) {
	mapBox, mx := boxes.Boxes.Get(boxID)
	if mapBox != nil {
		boxX, boxY := game_math.GetXYCenterHex(mapBox.Q, mapBox.R)

		dist := game_math.GetBetweenDist(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, boxX, boxY)
		if dist < 150 {
			return nil, mapBox, mx
		} else {
			mx.Unlock()
			return errors.New("no min dist"), nil, nil
		}
	} else {
		mx.Unlock()
		return errors.New("no find box"), nil, nil
	}
}
