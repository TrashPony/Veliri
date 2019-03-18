package lobby

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/base"
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/TrashPony/Veliri/src/webSocket/globalGame"
	"time"
)

func OutBase(user *player.Player) error {

	// todo проверить топливо

	if user.GetSquad() != nil && user.GetSquad().MatherShip.Body != nil && user.GetSquad().MatherShip.HP > 0 {

		gameBase, find := bases.Bases.Get(user.InBaseID)
		if !find {
			return errors.New("no base")
		}

		for globalGame.CheckTransportCoordinate(gameBase.RespQ, gameBase.RespR, 10, 95, gameBase.MapID) {
			// запускаем механизм проверки и эвакуации игрока с респауна))))
			time.Sleep(time.Millisecond * 100)
		}

		user.GetSquad().Q = gameBase.RespQ
		user.GetSquad().R = gameBase.RespR
		user.GetSquad().MapID = gameBase.MapID
		user.LastBaseID = user.InBaseID
		base.UserOutBase(user.GetID())
		user.InBaseID = 0

		dbPlayer.UpdateUser(user)
		update.Squad(user.GetSquad(), true)
		return nil
	} else {
		if user.GetSquad() == nil || user.GetSquad().MatherShip.Body == nil {
			return errors.New("no body")
		}
		if user.GetSquad().MatherShip.HP == 0 {
			return errors.New("body damage")
		}
		return nil
	}
}
