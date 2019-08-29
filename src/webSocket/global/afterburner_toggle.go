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

	Move(user, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
	go SendMessage(Message{
		Event:       "AfterburnerToggle",
		Afterburner: user.GetSquad().MatherShip.Afterburner,
		IDUserSend:  user.GetID(),
		IDMap:       user.GetSquad().MapID})
}
