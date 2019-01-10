package globalGame

import (
	"../../mechanics/factories/boxes"
	"../../mechanics/gameObjects/boxInMap"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func placeNewBox(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)
	if user != nil {
		err, newBox := globalGame.PlaceNewBox(user, msg.Slot, msg.BoxPassword)
		if err != nil {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID()}
		} else {
			globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID()}
			globalPipe <- Message{Event: "NewBox", Box: newBox, X: user.GetSquad().GlobalX, Y: user.GetSquad().GlobalY}
		}
	}
}

func openBox(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)

	if user != nil {
		mapBox, mx := boxes.Boxes.Get(msg.BoxID)
		mx.Unlock()

		if mapBox != nil {

			x, y := globalGame.GetXYCenterHex(mapBox.Q, mapBox.R)
			dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)

			if dist < 150 {
				if user.GetSquad().MoveChecker {
					user.GetSquad().GetMove() <- true // останавливаем движение
				}

				if mapBox.Protect {
					if mapBox.GetPassword() == msg.BoxPassword || mapBox.GetPassword() == 0 {
						globalPipe <- Message{Event: msg.Event, BoxID: mapBox.ID, Inventory: mapBox.GetStorage(),
							Size: mapBox.CapacitySize, idUserSend: user.GetID()}
					} else {
						if msg.BoxPassword == 0 {
							globalPipe <- Message{Event: msg.Event, BoxID: mapBox.ID, Error: "need password", idUserSend: user.GetID()}
						} else {
							globalPipe <- Message{Event: "Error", BoxID: mapBox.ID, Error: "wrong password", idUserSend: user.GetID()}
						}
					}
				} else {
					globalPipe <- Message{Event: msg.Event, BoxID: mapBox.ID, Inventory: mapBox.GetStorage(),
						Size: mapBox.CapacitySize, idUserSend: user.GetID()}
				}
			}
		}
	}
}

func useBox(ws *websocket.Conn, msg Message) {
	var err error
	var mapBox *boxInMap.Box

	user := Clients.GetByWs(ws)

	if user != nil {
		if msg.Event == "getItemFromBox" {
			err, mapBox = globalGame.GetItemFromBox(user, msg.BoxID, msg.Slot)
		} else {
			err, mapBox = globalGame.PlaceItemToBox(user, msg.BoxID, msg.Slot)
		}

		if err != nil {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID()}
		} else {
			globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID()}

			updateBoxInfo(mapBox)
		}
	}
}

func boxToBox(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)

	err, getBox, toBox := globalGame.BoxToBox(user, msg.BoxID, msg.Slot, msg.ToBoxID)
	if err != nil {
		globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID()}
	} else {
		globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID()}

		updateBoxInfo(getBox)
		updateBoxInfo(toBox)
	}
}

func updateBoxInfo(box *boxInMap.Box)  {
	usersGlobalWs, mx := Clients.GetAll()
	mx.Unlock()

	for _, user := range usersGlobalWs {
		boxX, boxY := globalGame.GetXYCenterHex(box.Q, box.R)
		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, boxX, boxY)

		if dist < 175 { // что бы содержимое ящика не видили те кто далеко
			globalPipe <- Message{Event: "UpdateBox", idUserSend: user.GetID(), BoxID: box.ID,
				Inventory: box.GetStorage(), Size: box.CapacitySize}
		}
	}
}
