package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func throwItems(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)

	if user != nil {
		err, newBox, box := globalGame.ThrowItems(user, msg.ThrowItems)

		if err != nil {
			go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
		} else {
			go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
			if newBox {
				go SendMessage(Message{Event: "NewBox", Box: box, X: user.GetSquad().GlobalX, Y: user.GetSquad().GlobalY, IDMap: user.GetSquad().MapID})
			} else {
				// если мы не создали новый ящик то обновляем старый у всех кто ближе мин растояния
				users, rLock := globalGame.Clients.GetAll()
				defer rLock.Unlock()
				for _, user := range users {
					boxX, boxY := globalGame.GetXYCenterHex(box.Q, box.R)
					dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, boxX, boxY)

					if dist < 175 { // что бы содержимое ящика не видили те кто далеко
						go SendMessage(Message{Event: "UpdateBox", IDUserSend: user.GetID(), BoxID: box.ID,
							Inventory: box.GetStorage(), Size: box.CapacitySize, IDMap: user.GetSquad().MapID})
					}
				}
			}
		}
	}
}
