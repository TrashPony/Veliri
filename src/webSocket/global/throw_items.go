package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
)

func throwItems(user *player.Player, msg Message) {

	err, newBox, box := globalGame.ThrowItems(user, msg.ThrowItems)

	if err != nil {
		go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
	} else {
		go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
		if newBox {
			go SendMessage(Message{Event: "NewBox", Box: box, X: user.GetSquad().MatherShip.X, Y: user.GetSquad().MatherShip.Y, IDMap: user.GetSquad().MapID})
		} else {
			// если мы не создали новый ящик то обновляем старый у всех кто ближе мин растояния
			users, rLock := globalGame.Clients.GetAll()
			defer rLock.Unlock()
			for _, user := range users {
				boxX, boxY := globalGame.GetXYCenterHex(box.Q, box.R)
				dist := globalGame.GetBetweenDist(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, boxX, boxY)

				if dist < 175 { // что бы содержимое ящика не видили те кто далеко
					go SendMessage(Message{Event: "UpdateBox", IDUserSend: user.GetID(), BoxID: box.ID,
						Inventory: box.GetStorage(), Size: box.CapacitySize, IDMap: user.GetSquad().MapID})
				}
			}
		}
	}
}
