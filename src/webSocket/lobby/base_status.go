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
		// todo не спамить а сравнивать предыдущие и текущие состояние, если есть различия то отсылать
		for _, user := range usersLobbyWs {
			BaseStatus(user)
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

	lobbyPipe <- Message{
		Event:  "BaseStatus",
		UserID: user.GetID(),
		Base:   userBase,
	}
}
