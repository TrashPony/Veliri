package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func afterburnerToggle(user *player.Player, msg Message) {
	// перегрузка доступна только у мс

	if user.GetSquad().MatherShip.Afterburner {
		user.GetSquad().MatherShip.Afterburner = false
	} else {
		user.GetSquad().MatherShip.Afterburner = true
	}

	msg.ToX = user.GetSquad().MatherShip.ToX
	msg.ToY = user.GetSquad().MatherShip.ToY

	// пересчитываем путь т.к. эффективность двиготеля изменилась
	// перегрузку может использовать только мп поэтому создаем масив с одним юнитом
	msg.UnitsID = []int{user.GetSquad().MatherShip.ID}
	if user.GetSquad().MatherShip.MoveChecker {
		Move(user, msg, false)
	}

	go SendMessage(Message{
		Event:       "AfterburnerToggle",
		Afterburner: user.GetSquad().MatherShip.Afterburner,
		IDUserSend:  user.GetID(),
		IDMap:       user.GetSquad().MatherShip.MapID})
}
