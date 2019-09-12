package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/box"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func throwItems(user *player.Player, msg Message) {
	// выкинуть вещи из  трюма может только мп
	err, newBox, mapBox := box.ThrowItems(user, msg.ThrowItems)

	if err != nil {
		go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
	} else {
		go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
		if newBox {
			go SendMessage(Message{Event: "NewBox", Box: mapBox, X: user.GetSquad().MatherShip.X, Y: user.GetSquad().MatherShip.Y, IDMap: user.GetSquad().MatherShip.MapID})
		} else {
			// если мы не создали новый ящик то обновляем старый у всех кто ближе мин растояния
			users, rLock := globalGame.Clients.GetAll()
			defer rLock.Unlock()
			for _, user := range users {
				boxX, boxY := game_math.GetXYCenterHex(mapBox.Q, mapBox.R)
				dist := game_math.GetBetweenDist(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, boxX, boxY)

				if dist < 175 { // что бы содержимое ящика не видили те кто далеко
					go SendMessage(Message{Event: "UpdateBox", IDUserSend: user.GetID(), BoxID: mapBox.ID,
						Inventory: mapBox.GetStorage(), Size: mapBox.CapacitySize, IDMap: user.GetSquad().MatherShip.MapID})
				}
			}
		}
	}
}
