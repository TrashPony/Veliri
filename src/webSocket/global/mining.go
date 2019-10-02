package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"time"
)

func startMining(miner *unit.Unit, msg Message) {
	reservoir := maps.Maps.GetReservoirByXY(msg.X, msg.Y, miner.MapID)
	if reservoir == nil {
		go SendMessage(Message{Event: "Error", Error: "no reservoir", IDUserSend: miner.OwnerID, IDMap: miner.MapID})
		return
	}

	miningEquip := miner.Body.GetEquip(msg.TypeSlot, msg.Slot)
	if miningEquip == nil || miningEquip.Equip == nil && miningEquip.Equip.Applicable == reservoir.Type {
		go SendMessage(Message{Event: "Error", Error: "no equip", IDUserSend: miner.OwnerID, IDMap: miner.MapID})
		return
	}

	if miner.Body.CapacitySize < miner.Inventory.GetSize()+reservoir.Resource.Size {
		go SendMessage(Message{Event: "Error", Error: "inventory is full", IDUserSend: miner.OwnerID, IDMap: miner.MapID})
		return
	}

	dist := game_math.GetBetweenDist(miner.X, miner.Y, reservoir.X, reservoir.Y)
	if int(dist) < miningEquip.Equip.Radius && !miningEquip.Equip.MiningChecker {

		go SendMessage(Message{Event: msg.Event, ShortUnit: miner.GetShortInfo(), Seconds: miningEquip.Equip.Reload,
			TypeSlot: msg.TypeSlot, Slot: msg.Slot, X: reservoir.X, Y: reservoir.Y, IDMap: miner.MapID})

		miningEquip.Equip.MiningChecker = true
		miningEquip.Equip.CreateMining()

		go Mining(miner, miningEquip.Equip, reservoir, msg)
	} else {
		if int(dist) > miningEquip.Equip.Radius {
			go SendMessage(Message{Event: "Error", Error: "not enough distance", IDUserSend: miner.OwnerID, IDMap: miner.MapID})
			return
		}
		if miningEquip.Equip.MiningChecker {
			go SendMessage(Message{Event: "Error", Error: "extractor work", IDUserSend: miner.OwnerID, IDMap: miner.MapID})
			return
		}
	}
}

func Mining(miner *unit.Unit, miningEquip *equip.Equip, reservoir *resource.Map, msg Message) {
	exit := false

	for {

		// переменная для проверки времени цикла
		miningEquip.CurrentReload = miningEquip.Reload

		// проверка на полный трюм
		if miner.Body.CapacitySize < miner.Inventory.GetSize()+reservoir.Resource.Size {
			go SendMessage(Message{Event: "Error", Error: "inventory is full", IDUserSend: miner.OwnerID, IDMap: miner.MapID})
			miningEquip.MiningChecker = false
			return
		}

		for miningEquip.CurrentReload > 0 {
			select {
			case exitNow := <-miningEquip.GetMining():
				if exitNow {
					// игрок сам отменить копание
					go SendMessage(Message{Event: "stopMining", ShortUnit: miner.GetShortInfo(), Seconds: miningEquip.Reload,
						TypeSlot: msg.TypeSlot, Slot: msg.Slot, IDMap: miner.MapID})
					exit = true
				}
			default:

				if globalGame.Clients.GetById(miner.OwnerID) == nil {
					// игрок вышел
					go SendMessage(Message{Event: "stopMining", ShortUnit: miner.GetShortInfo(), Seconds: miningEquip.Reload,
						TypeSlot: msg.TypeSlot, Slot: msg.Slot, IDMap: miner.MapID})
					exit = true
				}

				dist := game_math.GetBetweenDist(miner.X, miner.Y, reservoir.X, reservoir.Y)

				if int(dist) > miningEquip.Radius {
					// игрок уехал слишком далеко
					go SendMessage(Message{Event: "stopMining", ShortUnit: miner.GetShortInfo(), Seconds: miningEquip.Reload,
						TypeSlot: msg.TypeSlot, Slot: msg.Slot, IDMap: miner.MapID})
					exit = true
				}

				miningEquip.CurrentReload--
				time.Sleep(time.Second)
			}
		}

		if exit {
			miningEquip.MiningChecker = false
			return
		}

		// проверка сколько итемов влезет в трюм
		// Region определяет кубоменты доставаемые итемом
		var countRes int
		if miner.Body.CapacitySize >= miner.Inventory.GetSize()+float32(miningEquip.Region) {
			countRes = int(float32(miningEquip.Region) / reservoir.Resource.Size)
		} else {
			countRes = int((miner.Body.CapacitySize - miner.Inventory.GetSize()) / reservoir.Resource.Size)
		}

		if reservoir.Count < countRes {
			miner.Inventory.AddItem(reservoir.Resource, "resource", reservoir.Resource.TypeID,
				reservoir.Count, 1, reservoir.Resource.Size, 1, false, miner.OwnerID)

			reservoir.Count = 0
		} else {
			miner.Inventory.AddItem(reservoir.Resource, "resource", reservoir.Resource.TypeID,
				countRes, 1, reservoir.Resource.Size, 1, false, miner.OwnerID)

			reservoir.Count -= countRes
		}

		//update.Squad(user.GetSquad(), true) todo

		go SendMessage(Message{Event: "UpdateInventory", IDUserSend: miner.OwnerID, IDMap: miner.MapID})
		go SendMessage(Message{Event: "updateReservoir", X: reservoir.X, Y: reservoir.Y, Count: reservoir.Count, IDMap: miner.MapID})

		if reservoir.Count == 0 {
			// если руда капается в несколько руд, то пусть остановяться все лазеры )
			go SendMessage(Message{Event: "stopMining", ShortUnit: miner.GetShortInfo(), Seconds: miningEquip.Reload,
				TypeSlot: msg.TypeSlot, Slot: msg.Slot, IDMap: miner.MapID})

			maps.Maps.RemoveReservoirByQR(reservoir.X, reservoir.Y, reservoir.MapID)
			go SendMessage(Message{Event: "destroyReservoir", ShortUnit: miner.GetShortInfo(), X: reservoir.X,
				Y: reservoir.Y, IDMap: miner.MapID})

			miningEquip.MiningChecker = false
			return
		} else {
			go SendMessage(Message{Event: msg.Event, ShortUnit: miner.GetShortInfo(), Seconds: miningEquip.Reload,
				TypeSlot: msg.TypeSlot, Slot: msg.Slot, X: reservoir.X, Y: reservoir.Y, IDMap: miner.MapID})
		}
	}
}
