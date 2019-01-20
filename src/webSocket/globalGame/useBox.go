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
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
		} else {
			globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
			globalPipe <- Message{Event: "NewBox", Box: newBox, X: user.GetSquad().GlobalX, Y: user.GetSquad().GlobalY, idMap: user.GetSquad().MapID}
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
				stopMove(ws, true)

				if mapBox.Protect {
					if mapBox.GetPassword() == msg.BoxPassword || mapBox.GetPassword() == 0 {
						globalPipe <- Message{Event: msg.Event, BoxID: mapBox.ID, Inventory: mapBox.GetStorage(),
							Size: mapBox.CapacitySize, idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
					} else {
						if msg.BoxPassword == 0 {
							globalPipe <- Message{Event: msg.Event, BoxID: mapBox.ID, Error: "need password",
								idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
						} else {
							globalPipe <- Message{Event: "Error", BoxID: mapBox.ID, Error: "wrong password",
								idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
						}
					}
				} else {
					globalPipe <- Message{Event: msg.Event, BoxID: mapBox.ID, Inventory: mapBox.GetStorage(),
						Size: mapBox.CapacitySize, idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
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
			updateBoxInfo(mapBox)
		}

		if msg.Event == "placeItemToBox" {
			err, mapBox = globalGame.PlaceItemToBox(user, msg.BoxID, msg.Slot)
			updateBoxInfo(mapBox)
		}

		if msg.Event == "getItemsFromBox" {
			for _, i := range msg.Slots {
				err, mapBox = globalGame.GetItemFromBox(user, msg.BoxID, i)
				updateBoxInfo(mapBox)
			}
		}

		if msg.Event == "placeItemsToBox" {
			for _, i := range msg.Slots {
				err, mapBox = globalGame.PlaceItemToBox(user, msg.BoxID, i)
				updateBoxInfo(mapBox)
			}
		}

		if err != nil {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
		}
		globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
	}
}

func boxToBox(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)

	if msg.BoxID == msg.ToBoxID {
		return
	}

	if msg.Event == "boxToBoxItem" {
		err, getBox, toBox := globalGame.BoxToBox(user, msg.BoxID, msg.Slot, msg.ToBoxID)
		if err != nil {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
		} else {
			updateBoxInfo(getBox)
			updateBoxInfo(toBox)
		}
	}

	if msg.Event == "boxToBoxItems" {
		for _, i := range msg.Slots {
			err, getBox, toBox := globalGame.BoxToBox(user, msg.BoxID, i, msg.ToBoxID)
			if err != nil {
				globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
			} else {
				updateBoxInfo(getBox)
				updateBoxInfo(toBox)
			}
		}
	}

	globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
}

func updateBoxInfo(box *boxInMap.Box) {
	usersGlobalWs, mx := Clients.GetAll()
	mx.Unlock()

	if box == nil {
		return
	}

	for _, user := range usersGlobalWs {
		boxX, boxY := globalGame.GetXYCenterHex(box.Q, box.R)
		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, boxX, boxY)

		if dist < 175 { // что бы содержимое ящика не видили те кто далеко
			globalPipe <- Message{Event: "UpdateBox", idUserSend: user.GetID(), BoxID: box.ID,
				Inventory: box.GetStorage(), Size: box.CapacitySize, idMap: user.GetSquad().MapID}
		}
	}
}
