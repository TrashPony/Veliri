package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func throwItems(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)

	if user != nil {
		err, newBox, box := globalGame.ThrowItems(user, msg.ThrowItems)

		if err != nil {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
		} else {
			globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
			if newBox {
				globalPipe <- Message{Event: "NewBox", Box: box, X: user.GetSquad().GlobalX, Y: user.GetSquad().GlobalY, idMap: user.GetSquad().MapID}
			} else {
				// если мы не создали новый ящик то обновляем старый у всех кто ближе мин растояния
				for _, user := range Clients.GetAll() {
					boxX, boxY := globalGame.GetXYCenterHex(box.Q, box.R)
					dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, boxX, boxY)

					if dist < 175 { // что бы содержимое ящика не видили те кто далеко
						globalPipe <- Message{Event: "UpdateBox", idUserSend: user.GetID(), BoxID: box.ID,
							Inventory: box.GetStorage(), Size: box.CapacitySize, idMap: user.GetSquad().MapID}
					}
				}

			}
		}
	}
}
