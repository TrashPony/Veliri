package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamic_map_object"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/find_path"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
	"time"
)

func LoadGame(user *player.Player, msg Message) {

	mp, find := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)

	if find {

		// обнуляем все параметры глобальной игры
		user.GetSquad().MatherShip.Afterburner = false
		user.GetSquad().MatherShip.MoveChecker = false
		user.GetSquad().MatherShip.ActualPath = nil
		user.GetSquad().MatherShip.HighGravity = move.GetGravity(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, user.GetSquad().MatherShip.MapID)

		GetPlaceSquad(user, mp)

		if user.MemoryDynamicObjects == nil {
			user.MemoryDynamicObjects = make(map[int]map[int]*dynamic_map_object.Object)
		}

		go SendMessage(Message{
			Event:          msg.Event,
			Map:            mp,
			User:           user,
			Squad:          user.GetSquad(),
			Bases:          bases.Bases.GetBasesByMap(mp.Id),
			IDUserSend:     user.GetID(),
			Credits:        user.GetCredits(),
			IDMap:          user.GetSquad().MatherShip.MapID,
			Bot:            user.Bot,
			HighGravity:    move.GetGravity(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, user.GetSquad().MatherShip.MapID),
			DynamicObjects: user.MemoryDynamicObjects,
		})

		// запускаем реактор машины
		if user.GetSquad().RecoveryPowerWork {
			user.GetSquad().RecoveryPowerExit = true
			for user.GetSquad().RecoveryPowerWork {
				time.Sleep(100 * time.Millisecond)
			}
		}
		go RecoveryPowerWorker(user)

		// отслеживание целей
		if user.GetSquad().GunWorkerWork {
			user.GetSquad().GunWorkerExit = true
			for user.GetSquad().GunWorkerWork {
				time.Sleep(100 * time.Millisecond)
			}
		}
		go GunWorker(user)

		// работа обзора и радара отряда
		if user.GetSquad().RadarWorkerWork {
			user.GetSquad().RadarWorkerExit = true
			for user.GetSquad().RadarWorkerWork {
				time.Sleep(100 * time.Millisecond)
			}
		}
		go RadarWorker(user, mp)

		// находим аномалии
		equipSlot := user.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
		anomalies, err := globalGame.GetVisibleAnomaly(user, equipSlot)
		if err == nil {
			go SendMessage(Message{Event: "AnomalySignal", IDUserSend: user.GetID(), Anomalies: anomalies, IDMap: user.GetSquad().MatherShip.MapID, Bot: user.Bot})
		}
	} else {
		go SendMessage(Message{Event: "Error", Error: "no allow", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID, Bot: user.Bot})
	}
}

func GetPlaceSquad(user *player.Player, mp *_map.Map) {
	GetPlaceCoordinate(user.GetSquad().MatherShip, mp)
	for _, unitSlot := range user.GetSquad().MatherShip.Units {
		if unitSlot != nil && unitSlot.Unit != nil && unitSlot.Unit.OnMap {
			GetPlaceCoordinate(unitSlot.Unit, mp)
		}
	}
}

func GetPlaceCoordinate(placeUnit *unit.Unit, mp *_map.Map) {
	x, y, _ := find_path.SearchEndPoint(
		float64(placeUnit.X),
		float64(placeUnit.Y),
		float64(placeUnit.X),
		float64(placeUnit.Y),
		placeUnit,
		mp,
		globalGame.Clients.GetAllShortUnits(mp.Id),
	)

	placeUnit.X = int(x)
	placeUnit.Y = int(y)
}
