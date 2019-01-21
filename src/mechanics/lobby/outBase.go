package lobby

import (
	"../../webSocket/globalGame"
	"../db/base"
	"../db/squad/update"
	"../factories/bases"
	"../player"
	"errors"
	"time"
)

func OutBase(user *player.Player) error {

	// todo проверить топливо
	// todo проверить что бы респаун был свободен

	if user.GetSquad().MatherShip.Body != nil && user.GetSquad().MatherShip.HP > 0 {

		gameBase, find := bases.Bases.Get(user.InBaseID)
		if !find {
			return errors.New("no base")
		}

		for globalGame.CheckTransportCoordinate(gameBase.RespQ, gameBase.RespR, 10) {
			// запускаем механизм проверки и эвакуации игрока с респауна))))
			time.Sleep(time.Millisecond * 100)
		}

		user.GetSquad().Q = gameBase.RespQ
		user.GetSquad().R = gameBase.RespR
		user.GetSquad().MapID = gameBase.MapID

		base.UserOutBase(user.GetID())
		user.InBaseID = 0

		update.Squad(user.GetSquad(), true)
		return nil
	} else {
		if user.GetSquad().MatherShip.Body == nil {
			return errors.New("no body")
		}
		if user.GetSquad().MatherShip.HP == 0 {
			return errors.New("body damage")
		}
		return nil
	}
}
