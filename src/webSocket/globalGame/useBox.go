package globalGame

import (
	"../../mechanics/factories/boxes"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
	"../../mechanics/gameObjects/box"
)

func openBox(ws *websocket.Conn, msg Message, stopMove chan bool, moveChecker *bool) {
	user := usersGlobalWs[ws]

	mapBox := boxes.Boxes.Get(msg.BoxID)

	if mapBox != nil {
		x, y := globalGame.GetXYCenterHex(mapBox.Q, mapBox.R)

		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)

		if dist < 150 {
			if *moveChecker {
				stopMove <- true // останавливаем движение
			}
			ws.WriteJSON(Message{Event: msg.Event, BoxID: mapBox.ID, Inventory: mapBox.GetStorage(), Size: mapBox.CapacitySize})
		}
	}
}

func useBox(ws *websocket.Conn, msg Message) {
	var err error
	var mapBox *box.Box

	if msg.Event == "getItemFromBox" {
		err, mapBox = globalGame.GetItemFromBox(usersGlobalWs[ws], msg.BoxID, msg.Slot)
	} else {
		err, mapBox = globalGame.PlaceItemToBox(usersGlobalWs[ws], msg.BoxID, msg.Slot)
	}

	if err != nil {
		ws.WriteJSON(Message{Event: "Error", Error: err.Error()})
	} else {

		ws.WriteJSON(Message{Event: "UpdateInventory"})

		for ws, user := range usersGlobalWs {

			boxX, boxY := globalGame.GetXYCenterHex(mapBox.Q, mapBox.R)
			dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, boxX, boxY)
			println(dist)
			if dist < 175 { // что бы содержимое ящика не видили те кто далеко
				ws.WriteJSON(Message{Event: "UpdateBox", BoxID: mapBox.ID, Inventory: mapBox.GetStorage(), Size: mapBox.CapacitySize})
			}
		}
	}
}
