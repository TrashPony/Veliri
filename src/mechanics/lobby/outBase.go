package lobby

import (
	"../player"
	"errors"
	"../factories/bases"
	"../db/squad/update"
	"../db/base"
	"../../webSocket/globalGame"
)

func OutBase(user *player.Player) error {

	// todo проверить топливо
	// todo проверить что бы респаун был свободен

	if user.GetSquad().MatherShip.Body != nil && user.GetSquad().MatherShip.HP > 0 {


		gameBase, find := bases.Bases.Get(user.InBaseID)
		if !find {
			return errors.New("no base")
		}

		globalGame.RespCheck(gameBase) // запускаем механизм проверки и эвакуации игрока с респауна))))

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
