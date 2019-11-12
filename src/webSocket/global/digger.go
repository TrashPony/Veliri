package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
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
	// TODO добавить кратер
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
				// todo обьект с диалогом
			}
		}

		go SendMessage(Message{
			Event:     msg.Event,
			ShortUnit: user.GetSquad().MatherShip.GetShortInfo(),
			X:         msg.X, Y: msg.Y,
			TypeSlot: msg.TypeSlot, Slot: msg.Slot,
			Box:       box,
			Reservoir: res,
			// todo DynamicObject: &dynamicObject,
			Name:          diggerSlot.Equip.Name,
			IDMap:         user.GetSquad().MatherShip.MapID,
			NeedCheckView: true,
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
			Box:       nil,
			Reservoir: nil,
			// todo DynamicObject: &dynamicObject,
			Name:          diggerSlot.Equip.Name,
			IDMap:         user.GetSquad().MatherShip.MapID,
			NeedCheckView: true,
		})
	}
}
