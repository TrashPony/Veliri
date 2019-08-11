package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func SellDetail(user *player.Player, msg Message) {
	// todo продумать систему ценобразования на базах

	// можно продовать ток 10 или 100
	if msg.Count == 10 || msg.Count == 100 {

		userBase, _ := bases.Bases.Get(user.InBaseID)
		storage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)
		// база может купить только сырье, тесть recycle
		err := storage.RemoveItem(msg.ID, "recycle", msg.Count)
		if err != nil {
			lobbyPipe <- Message{Event: "Error", Error: err.Error(), UserID: user.GetID()}
			return
		}

		// отдаем ресурс базе
		if userBase.CurrentResources[msg.ID] != nil {
			userBase.CurrentResources[msg.ID].Quantity += msg.Count
		}

		// todo проверить текущую цену с тем что пришло в сообщение, если разные отправлять ошибку msg.Price
		user.SetCredits(user.GetCredits() + msg.Count*1)

		//обновлять фронтенд
		GetDetails(user)
	} else {
		lobbyPipe <- Message{Event: "Error", Error: "wrong count", UserID: user.GetID()}
		return
	}
}
