package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
	"time"
)

func startMining(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)
	if user != nil {
		reservoir := maps.Maps.GetReservoirByQR(msg.Q, msg.R, user.GetSquad().MapID)
		if reservoir == nil {
			go SendMessage(Message{Event: "Error", Error: "no reservoir", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
			return
		}

		miningEquip := user.GetSquad().MatherShip.Body.GetEquip(msg.TypeSlot, msg.Slot)
		if miningEquip == nil || miningEquip.Equip == nil && miningEquip.Equip.Applicable == reservoir.Type {
			go SendMessage(Message{Event: "Error", Error: "no equip", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
			return
		}

		if user.GetSquad().MatherShip.Body.CapacitySize <= user.GetSquad().Inventory.GetSize()+reservoir.Resource.Size {
			go SendMessage(Message{Event: "Error", Error: "inventory is full", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
			return
		}

		x, y := globalGame.GetXYCenterHex(reservoir.Q, reservoir.R)
		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
		if int(dist) < miningEquip.Equip.Radius*100 && !miningEquip.Equip.MiningChecker {

			go SendMessage(Message{Event: msg.Event, OtherUser: user.GetShortUserInfo(true), Seconds: miningEquip.Equip.Reload,
				TypeSlot: msg.TypeSlot, Slot: msg.Slot, Q: reservoir.Q, R: reservoir.R, IDMap: user.GetSquad().MapID})

			miningEquip.Equip.MiningChecker = true
			miningEquip.Equip.CreateMining()

			go Mining(ws, user, miningEquip.Equip, reservoir, msg)
		} else {
			if int(dist) > miningEquip.Equip.Radius*100 {
				go SendMessage(Message{Event: "Error", Error: "not enough distance", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
				return
			}
			if miningEquip.Equip.MiningChecker {
				go SendMessage(Message{Event: "Error", Error: "extractor work", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
				return
			}
		}
	}
}

func Mining(ws *websocket.Conn, user *player.Player, miningEquip *equip.Equip, reservoir *resource.Map, msg Message) {
	exit := false
	for {

		timeCount := 0 // переменная для проверки времени цикла

		for miningEquip.Reload > timeCount {
			select {
			case exitNow := <-miningEquip.GetMining():
				if exitNow {
					// игрок сам отменить копание
					go SendMessage(Message{Event: "stopMining", OtherUser: user.GetShortUserInfo(true), Seconds: miningEquip.Reload,
						TypeSlot: msg.TypeSlot, Slot: msg.Slot, IDMap: user.GetSquad().MapID})
					exit = true
				}
			default:

				if ws == nil || globalGame.Clients.GetByWs(ws) == nil {
					// игрок вышел
					go SendMessage(Message{Event: "stopMining", OtherUser: user.GetShortUserInfo(true), Seconds: miningEquip.Reload,
						TypeSlot: msg.TypeSlot, Slot: msg.Slot, IDMap: user.GetSquad().MapID})
					exit = true
				}

				x, y := globalGame.GetXYCenterHex(reservoir.Q, reservoir.R)
				dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)

				if int(dist) > miningEquip.Radius*100 {
					// игрок уехал слишком далеко
					go SendMessage(Message{Event: "stopMining", OtherUser: user.GetShortUserInfo(true), Seconds: miningEquip.Reload,
						TypeSlot: msg.TypeSlot, Slot: msg.Slot, IDMap: user.GetSquad().MapID})
					exit = true
				}

				timeCount++
				time.Sleep(time.Second)
			}
		}

		if exit {
			miningEquip.MiningChecker = false
			return
		}

		// проверка на полный трюм
		if user.GetSquad().MatherShip.Body.CapacitySize <= user.GetSquad().Inventory.GetSize()+reservoir.Resource.Size {
			go SendMessage(Message{Event: "Error", Error: "inventory is full", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
			miningEquip.MiningChecker = false
			return
		}

		// проверка сколько итемов влезет в трюм
		// Region определяет кубоменты доставаемые итемом
		var countRes int
		if user.GetSquad().MatherShip.Body.CapacitySize >= user.GetSquad().Inventory.GetSize()+float32(miningEquip.Region) {
			countRes = int(float32(miningEquip.Region) / reservoir.Resource.Size)
		} else {
			countRes = int((user.GetSquad().MatherShip.Body.CapacitySize - user.GetSquad().Inventory.GetSize()) / reservoir.Resource.Size)
		}

		if reservoir.Count < countRes {
			user.GetSquad().Inventory.AddItem(reservoir.Resource, "resource", reservoir.Resource.TypeID,
				reservoir.Count, 1, reservoir.Resource.Size, 1)

			reservoir.Count = 0
		} else {
			user.GetSquad().Inventory.AddItem(reservoir.Resource, "resource", reservoir.Resource.TypeID,
				countRes, 1, reservoir.Resource.Size, 1)

			reservoir.Count -= countRes
		}

		update.Squad(user.GetSquad(), true)

		go SendMessage(Message{Event: "UpdateInventory", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
		go SendMessage(Message{Event: "updateReservoir", Q: reservoir.Q, R: reservoir.R, Count: reservoir.Count,
			IDMap: user.GetSquad().MapID})

		if reservoir.Count == 0 {
			// если руда капается в несколько руд, то пусть остановяться все лазеры )
			go SendMessage(Message{Event: "stopMining", OtherUser: user.GetShortUserInfo(true), Seconds: miningEquip.Reload,
				TypeSlot: msg.TypeSlot, Slot: msg.Slot, IDMap: user.GetSquad().MapID})

			maps.Maps.RemoveReservoirByQR(reservoir.Q, reservoir.R, reservoir.MapID)
			go SendMessage(Message{Event: "destroyReservoir", OtherUser: user.GetShortUserInfo(true), Q: reservoir.Q,
				R: reservoir.R, IDMap: user.GetSquad().MapID})

			miningEquip.MiningChecker = false
			return
		} else {
			go SendMessage(Message{Event: msg.Event, OtherUser: user.GetShortUserInfo(true), Seconds: miningEquip.Reload,
				TypeSlot: msg.TypeSlot, Slot: msg.Slot, Q: reservoir.Q, R: reservoir.R, IDMap: user.GetSquad().MapID})
		}
	}
}
