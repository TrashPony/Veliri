package globalGame

import (
	"../../mechanics/factories/boxes"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
	"../../mechanics/gameObjects/box"
)

func openBox(ws *websocket.Conn, msg Message) {
	user := usersGlobalWs[ws]

	mapBox := boxes.Boxes.Get(msg.BoxID)

	if mapBox != nil {
		x, y := globalGame.GetXYCenterHex(mapBox.Q, mapBox.R)

		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)

		if dist < 150 {
			if user.GetSquad().MoveChecker {
				user.GetSquad().GetMove() <- true // останавливаем движение
			}
			ws.WriteJSON(Message{Event: msg.Event, BoxID: mapBox.ID, Inventory: mapBox.GetStorage(), Size: mapBox.CapacitySize})
		}
	}
}

func useBox(ws *websocket.Conn, msg Message) {
	var err error
	var mapBox *box.Box

	user, _ := usersGlobalWs[ws]

	if msg.Event == "getItemFromBox" {
		err, mapBox = globalGame.GetItemFromBox(user, msg.BoxID, msg.Slot)
	} else {
		err, mapBox = globalGame.PlaceItemToBox(user, msg.BoxID, msg.Slot)
	}

	if err != nil {
		globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID()}
	} else {

		globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID()}

		for _, user := range usersGlobalWs {

			boxX, boxY := globalGame.GetXYCenterHex(mapBox.Q, mapBox.R)
			dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, boxX, boxY)

			if dist < 175 { // что бы содержимое ящика не видили те кто далеко
				globalPipe <- Message{Event: "UpdateBox", idUserSend: user.GetID(), BoxID: mapBox.ID,
					Inventory: mapBox.GetStorage(), Size: mapBox.CapacitySize}
			}
		}
	}
}
