package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/box"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func placeNewBox(user *player.Player, msg Message) {
	// устанавливать ящики может только мп
	err, newBox := box.PlaceNewBox(user, msg.Slot, msg.BoxPassword)
	if err != nil {
		go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
	} else {
		go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
		go SendMessage(Message{Event: "NewBox", Box: newBox, X: user.GetSquad().MatherShip.X, Y: user.GetSquad().MatherShip.Y, IDMap: user.GetSquad().MatherShip.MapID})
	}
}

func openBox(user *player.Player, msg Message) {

	mapBox, mx := boxes.Boxes.Get(msg.BoxID)
	mx.Unlock()

	if mapBox != nil {

		dist := game_math.GetBetweenDist(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, mapBox.X, mapBox.Y)

		if dist < 75 {
			if mapBox.Protect {
				if mapBox.GetPassword() == msg.BoxPassword || mapBox.GetPassword() == 0 {
					go SendMessage(Message{Event: msg.Event, BoxID: mapBox.ID, Inventory: mapBox.GetStorage(),
						Size: mapBox.CapacitySize, IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
				} else {
					if msg.BoxPassword == 0 {
						go SendMessage(Message{Event: msg.Event, BoxID: mapBox.ID, Error: "need password",
							IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
					} else {
						go SendMessage(Message{Event: "Error", BoxID: mapBox.ID, Error: "wrong password",
							IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
					}
				}
			} else {
				go SendMessage(Message{Event: msg.Event, BoxID: mapBox.ID, Inventory: mapBox.GetStorage(),
					Size: mapBox.CapacitySize, IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
			}
		}
	}
}

func useBox(user *player.Player, msg Message) {
	var err error
	var mapBox *boxInMap.Box

	if msg.Event == "getItemFromBox" {
		err, mapBox = box.GetItemFromBox(user, msg.BoxID, msg.Slot)
		updateBoxInfo(mapBox)
	}

	if msg.Event == "placeItemToBox" {
		err, mapBox = box.PlaceItemToBox(user, msg.BoxID, msg.Slot)
		updateBoxInfo(mapBox)
	}

	if msg.Event == "getItemsFromBox" {
		for _, i := range msg.Slots {
			err, mapBox = box.GetItemFromBox(user, msg.BoxID, i)
			updateBoxInfo(mapBox)
		}
	}

	if msg.Event == "placeItemsToBox" {
		for _, i := range msg.Slots {
			err, mapBox = box.PlaceItemToBox(user, msg.BoxID, i)
			updateBoxInfo(mapBox)
		}
	}

	if err != nil {
		go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
	}
	go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
}

func boxToBox(user *player.Player, msg Message) {

	if msg.BoxID == msg.ToBoxID {
		return
	}

	if msg.Event == "boxToBoxItem" {
		err, getBox, toBox := box.BoxToBox(user, msg.BoxID, msg.Slot, msg.ToBoxID)
		if err != nil {
			go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
		} else {
			updateBoxInfo(getBox)
			updateBoxInfo(toBox)
		}
	}

	if msg.Event == "boxToBoxItems" {
		for _, i := range msg.Slots {
			err, getBox, toBox := box.BoxToBox(user, msg.BoxID, i, msg.ToBoxID)
			if err != nil {
				go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
			} else {
				updateBoxInfo(getBox)
				updateBoxInfo(toBox)
			}
		}
	}

	go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
}

func updateBoxInfo(box *boxInMap.Box) {

	if box == nil {
		return
	}

	users, rLock := globalGame.Clients.GetAll()
	defer rLock.Unlock()
	for _, user := range users {

		dist := game_math.GetBetweenDist(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, box.X, box.Y)

		if dist < 175 { // что бы содержимое ящика не видили те кто далеко
			go SendMessage(Message{Event: "UpdateBox", IDUserSend: user.GetID(), BoxID: box.ID,
				Inventory: box.GetStorage(), Size: box.CapacitySize, IDMap: user.GetSquad().MatherShip.MapID})
		}
	}
}
