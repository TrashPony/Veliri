package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/blueWorks"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/blueprints"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
	wsInventory "github.com/TrashPony/Veliri/src/webSocket/inventory"
	"time"
)

func openWorkbench(user *player.Player, msg Message) {
	baseStorage, find := storages.Storages.Get(user.GetID(), user.InBaseID)
	if find {
		lobbyPipe <- Message{
			Event:     "WorkbenchStorage",
			UserID:    user.GetID(),
			Storage:   baseStorage,
			BlueWorks: blueWorks.BlueWorks.GetByUserAndBase(user.GetID(), user.InBaseID),
		}
	}
}

func selectBP(user *player.Player, msg Message) {

	baseStorage, findStorage := storages.Storages.Get(user.GetID(), user.InBaseID)
	slot, findSlot := baseStorage.Slots[msg.StorageSlot]
	userBase, _ := bases.Bases.Get(user.InBaseID)

	if findStorage && findSlot && slot.Type == "blueprints" {

		bluePrint, _ := gameTypes.BluePrints.GetByID(slot.ItemID)
		recyclerItems := make([]*inventory.Slot, 0)
		lobby.ParseItems(&recyclerItems, 100+(userBase.GetSumEfficiency()-(user.CurrentSkills["materials_production"].Level*5)), bluePrint, msg.Count)

		for _, slot := range recyclerItems {
			slot.Find = baseStorage.ViewItems(slot.ItemID, slot.Type, slot.Quantity)
		}

		lobbyPipe <- Message{Event: "SelectBP", UserID: user.GetID(), PreviewRecycleSlots: recyclerItems,
			BluePrint: bluePrint, BPItem: gameTypes.BluePrints.GetItemsByBluePrintID(slot.ItemID), Count: msg.Count,
			MaxCount: slot.Quantity, StorageSlot: msg.StorageSlot,
			UserWorkSkillDetailPercent: user.CurrentSkills["materials_production"].Level * 5,
			UserWorkSkillTimePercent:   user.CurrentSkills["production_time"].Level * 5,
			Base:                       userBase,
		}
	} else {
		lobbyPipe <- Message{Event: "SelectBP", UserID: user.GetID(), BluePrint: nil}
	}
}

func craft(user *player.Player, msg Message) {

	baseStorage, findStorage := storages.Storages.Get(user.GetID(), user.InBaseID)
	slot, findSlot := baseStorage.Slots[msg.StorageSlot]
	userBase, _ := bases.Bases.Get(user.InBaseID)

	// TODO проверка
	//msg.UserWorkSkillTimePercent
	//msg.UserWorkSkillDetailPercent
	//msg.Efficiency

	if findStorage && findSlot && slot.Type == "blueprints" && slot.Quantity >= msg.Count {
		bluePrint, _ := gameTypes.BluePrints.GetByID(slot.ItemID)
		mineralTaxPercentage := 100 + (userBase.GetSumEfficiency() - (user.CurrentSkills["materials_production"].Level * 5))

		recyclerItems := make([]*inventory.Slot, 0)
		lobby.ParseItems(&recyclerItems, mineralTaxPercentage, bluePrint, msg.Count)

		for _, slot := range recyclerItems {
			if !baseStorage.ViewItems(slot.ItemID, slot.Type, slot.Quantity) {
				lobbyPipe <- Message{Event: msg.Event, Error: "few items"}
				return
			}
		}

		for i := 0; i < msg.Count; i++ { // для каждого итема новая работа

			// влияние скила на скорость крафта user.CurrentSkills["production_time"].Level * 5
			craftSecondsTime := bluePrint.CraftTime - ((bluePrint.CraftTime * user.CurrentSkills["production_time"].Level * 5) / 100)

			craftStartTime := blueWorks.BlueWorks.GetWorkTime(user.GetID(), user.InBaseID)
			startTime := time.Unix(craftStartTime, 0)
			craftStartTime += int64(craftSecondsTime)
			finishTime := time.Unix(craftStartTime, 0)

			newWork := blueprints.BlueWork{
				BlueprintID:          bluePrint.ID,
				BaseID:               user.InBaseID,
				UserID:               user.GetID(),
				StartTime:            startTime,
				FinishTime:           finishTime,
				TimeTaxPercentage:    -user.CurrentSkills["production_time"].Level * 5,
				MineralTaxPercentage: mineralTaxPercentage,
			}

			blueWorks.BlueWorks.Add(&newWork)
		}

		storages.Storages.RemoveItemBySlot(user.GetID(), user.InBaseID, msg.StorageSlot, msg.Count)
		for _, slot := range recyclerItems {
			storages.Storages.RemoveItem(user.GetID(), user.InBaseID, slot.ItemID, slot.Type, slot.Quantity)
		}

		wsInventory.UpdateStorage(user.GetID())
		selectBP(user, msg)
	}
}

func selectWork(user *player.Player, msg Message) {
	work := blueWorks.BlueWorks.GetByID(msg.ID)
	if work != nil && work.UserID == user.GetID() {

		bp, _ := gameTypes.BluePrints.GetByID(work.BlueprintID)

		percentRemainResource := 0
		maxCount := 0

		if work.GetDonePercent() > 0 {
			maxCount = 1
			percentRemainResource = work.MineralTaxPercentage - ((work.MineralTaxPercentage * work.GetDonePercent()) / 100)
		} else {

			works := blueWorks.BlueWorks.GetSameWorks(
				work.BlueprintID,
				work.MineralTaxPercentage,
				work.TimeTaxPercentage,
				user.GetID(),
				user.InBaseID,
				msg.ToTime/1000,    // промежуток выбраного времени, это гарантирует что выбраные работы удалятся по порядку времени
				msg.StartTime/1000, // промежуток выбраного времени
			)
			maxCount = len(works)
			percentRemainResource = work.MineralTaxPercentage
			percentRemainResource *= msg.Count
		}

		returnItems := make([]*inventory.Slot, 0)
		lobby.ParseItems(&returnItems, percentRemainResource, bp, 1)

		lobbyPipe <- Message{Event: msg.Event, UserID: user.GetID(), PreviewRecycleSlots: returnItems,
			BluePrint: bp, BPItem: gameTypes.BluePrints.GetItemsByBluePrintID(bp.ID), Count: msg.Count,
			StorageSlot: msg.StorageSlot, ID: msg.ID, BlueWork: work, MaxCount: maxCount, StartTime: msg.StartTime,
			ToTime: msg.ToTime}
	}
}

func cancelCraft(user *player.Player, msg Message) {

	work := blueWorks.BlueWorks.GetByID(msg.ID)
	if work != nil && work.UserID == user.GetID() {

		bp, _ := gameTypes.BluePrints.GetByID(work.BlueprintID)

		percentRemainResource := 0

		if work.GetDonePercent() > 0 {
			percentRemainResource = work.MineralTaxPercentage - ((work.MineralTaxPercentage * work.GetDonePercent()) / 100)
			blueWorks.BlueWorks.Remove(work)
		} else {

			// проверить что действительно есть количество работ равное Count, удалить эти работы
			works := blueWorks.BlueWorks.GetSameWorks(
				work.BlueprintID,
				work.MineralTaxPercentage,
				work.TimeTaxPercentage,
				user.GetID(),
				user.InBaseID,
				msg.ToTime/1000,    // промежуток выбраного времени, это гарантирует что выбраные работы удалятся по порядку времени
				msg.StartTime/1000, // промежуток выбраного времени
			)
			if len(works) >= msg.Count {
				percentRemainResource = work.MineralTaxPercentage
				percentRemainResource *= msg.Count

				for _, work := range works {
					blueWorks.BlueWorks.Remove(work)
				}
			} else {
				return
			}
		}

		returnItems := make([]*inventory.Slot, 0)
		lobby.ParseItems(&returnItems, percentRemainResource, bp, 1)

		for _, item := range returnItems {
			storages.Storages.AddSlot(user.GetID(), user.InBaseID, item)
		}
	}
	wsInventory.UpdateStorage(user.GetID())
}
