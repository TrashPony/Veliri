package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamicMapObject"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"math/rand"
	"time"
)

// TODO рефакторинг
func useDigger(user *player.Player, msg Message) {
	mp, _ := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)

	stopMove(user.GetSquad().MatherShip, true)

	diggerSlot := user.GetSquad().MatherShip.Body.GetEquip(msg.TypeSlot, msg.Slot)
	if diggerSlot == nil || diggerSlot.Equip == nil || diggerSlot.Equip.Applicable != "digger" ||
		diggerSlot.Equip.CurrentReload > 0 || user.GetSquad().MatherShip.Power < diggerSlot.Equip.UsePower {
		go SendMessage(Message{Event: "Error", Error: "no equip", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
		return
	}

	users, rLock := globalGame.Clients.GetAll()
	defer rLock.Unlock()

	user.GetSquad().MatherShip.Power -= diggerSlot.Equip.UsePower
	// перезаряджка
	diggerSlot.Equip.CurrentReload = diggerSlot.Equip.Reload
	go func() {
		for diggerSlot.Equip.CurrentReload > 0 {
			time.Sleep(1 * time.Second)
			diggerSlot.Equip.CurrentReload--
		}
	}()

	mpCoordinate, _ := mp.GetCoordinate(msg.X, msg.Y)

	var dynamicObject dynamicMapObject.DynamicObject
	dynamicObject.TextureBackground = "crater_2"
	dynamicObject.BackgroundScale = 75
	dynamicObject.BackgroundRotate = rand.Intn(360)
	mpCoordinate.DynamicObject = &dynamicObject

	// todo проверить что координата свободна от игрока
	anomaly := maps.Maps.GetMapAnomaly(mp.Id, msg.X, msg.Y)

	if anomaly != nil {
		maps.Maps.RemoveMapAnomaly(mp.Id, anomaly)

		box, res, AnomalyText := anomaly.GetLoot()
		// TODO запуст горутины на уничтожение

		if box != nil {

			box.MapID = mp.Id
			box.X = anomaly.GetX()
			box.X = anomaly.GetY()
			box.Rotate = rand.Intn(360)

			boxes.Boxes.InsertNewBox(box)
		}

		if res != nil {
			if mpCoordinate != nil {

				res.X = anomaly.GetX()
				res.Y = anomaly.GetY()
				res.Rotate = rand.Intn(360)
				res.MapID = mp.Id

				mp.AddResourceInMap(res)
			}
		}

		if AnomalyText != nil {
			if mpCoordinate != nil {
				dynamicObject.TextureObject = "infoAnomaly"
				dynamicObject.Dialog = AnomalyText
				dynamicObject.Destroyed = true
				dynamicObject.DestroyTime = time.Now()
				dynamicObject.ObjectScale = 20
				dynamicObject.ObjectRotate = rand.Intn(360)
				dynamicObject.Shadow = 50
				dynamicObject.Move = true
				dynamicObject.View = true
				dynamicObject.Attack = true
			}
		}

		go SendMessage(Message{
			Event:     msg.Event,
			ShortUnit: user.GetSquad().MatherShip.GetShortInfo(),
			X:         msg.X, Y: msg.Y,
			TypeSlot: msg.TypeSlot, Slot: msg.Slot,
			Box:           box,
			Reservoir:     res,
			DynamicObject: &dynamicObject,
			Name:          diggerSlot.Equip.Name,
			IDMap:         user.GetSquad().MatherShip.MapID,
		})

		for _, otherUser := range users {
			equipSlot := otherUser.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
			anomalies, err := globalGame.GetVisibleAnomaly(otherUser, equipSlot)
			if err == nil {
				go SendMessage(Message{Event: "AnomalySignal", IDUserSend: otherUser.GetID(),
					Anomalies: anomalies, IDMap: user.GetSquad().MatherShip.MapID})
			} else {
				go SendMessage(Message{Event: "RemoveAnomalies", IDUserSend: otherUser.GetID(),
					IDMap: user.GetSquad().MatherShip.MapID})
			}
		}

	} else {
		go SendMessage(Message{
			Event:     msg.Event,
			ShortUnit: user.GetSquad().MatherShip.GetShortInfo(),
			X:         msg.X, Y: msg.Y,
			TypeSlot: msg.TypeSlot, Slot: msg.Slot,
			Box:           nil,
			Reservoir:     nil,
			DynamicObject: &dynamicObject,
			Name:          diggerSlot.Equip.Name,
			IDMap:         user.GetSquad().MatherShip.MapID,
		})
	}
}
