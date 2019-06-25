package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"time"
)

// вокрер обновляет игрокам информацию состояния базы
func BaseStatusSender() {
	for {

		// todo возможна проблема конкуретного дотупа
		// сравнивать предыдущие и текущие состояние базы, если есть различия то отсылать статус базы
		for _, user := range usersLobbyWs {
			userBase, _ := bases.Bases.Get(user.InBaseID)
			if userBase != nil && user.LastBaseEfficiency != userBase.Efficiency {
				BaseStatus(user)
			}
		}

		time.Sleep(time.Second) // проверяем каждую секунду
	}
}

func BaseStatus(user *player.Player) {
	userBase, _ := bases.Bases.Get(user.InBaseID)
	userBase.GetSumEfficiency()
	for _, resource := range userBase.CurrentResources {
		userBase.GetRecyclePercent(resource.ItemID)
	}

	user.LastBaseEfficiency = userBase.Efficiency

	lobbyPipe <- Message{
		Event:  "BaseStatus",
		UserID: user.GetID(),
		Base:   userBase,
	}
}
