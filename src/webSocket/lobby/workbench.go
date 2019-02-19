package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/blueWorks"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/blueprints"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
	"github.com/TrashPony/Veliri/src/webSocket/storage"
	"github.com/gorilla/websocket"
	"time"
)

func openWorkbench(ws *websocket.Conn, msg Message) {
	user := usersLobbyWs[ws]

	baseStorage, find := storages.Storages.Get(user.GetID(), user.InBaseID)
	if user != nil && find {
		lobbyPipe <- Message{
			Event:     "WorkbenchStorage",
			UserID:    user.GetID(),
			Storage:   baseStorage,
			BlueWorks: blueWorks.BlueWorks.GetByUserAndBase(user.GetID(), user.InBaseID),
		}
	}
}

func selectBP(ws *websocket.Conn, msg Message) {
	user := usersLobbyWs[ws]

	if user != nil {

		baseStorage, findStorage := storages.Storages.Get(user.GetID(), user.InBaseID)
		slot, findSlot := baseStorage.Slots[msg.StorageSlot]

		if findStorage && findSlot && slot.Type == "blueprints" {

			bluePrint, _ := gameTypes.BluePrints.GetByID(slot.ItemID)

			recyclerItems := make([]*inventory.Slot, 0)
			lobby.ParseItems(&recyclerItems, 100, bluePrint, msg.Count)

			for _, slot := range recyclerItems {
				slot.Find = baseStorage.ViewItems(slot.ItemID, slot.Type, slot.Quantity)
			}

			lobbyPipe <- Message{Event: msg.Event, UserID: user.GetID(), PreviewRecycleSlots: recyclerItems,
				BluePrint: bluePrint, BPItem: gameTypes.BluePrints.GetItemsByBluePrintID(slot.ItemID), Count: msg.Count,
				MaxCount: slot.Quantity, StorageSlot: msg.StorageSlot}
		}
	}
}

func craft(ws *websocket.Conn, msg Message) {
	user := usersLobbyWs[ws]

	if user != nil {
		baseStorage, findStorage := storages.Storages.Get(user.GetID(), user.InBaseID)
		slot, findSlot := baseStorage.Slots[msg.StorageSlot]

		if findStorage && findSlot && slot.Type == "blueprints" && slot.Quantity >= msg.Count {
			bluePrint, _ := gameTypes.BluePrints.GetByID(slot.ItemID)

			recyclerItems := make([]*inventory.Slot, 0)
			lobby.ParseItems(&recyclerItems, 100, bluePrint, msg.Count)

			for _, slot := range recyclerItems {
				if !baseStorage.ViewItems(slot.ItemID, slot.Type, slot.Quantity) {
					lobbyPipe <- Message{Event: msg.Event, Error: "few items"}
					return
				}
			}

			for i := 0; i < msg.Count; i++ { // для каждого итема новая работа

				nowSecond := blueWorks.BlueWorks.GetWorkTime(user.GetID(), user.InBaseID)
				nowSecond += int64(bluePrint.CraftTime)
				finishTime := time.Unix(nowSecond, 0)

				newWork := blueprints.BlueWork{
					BlueprintID:             bluePrint.ID,
					BaseID:                  user.InBaseID,
					UserID:                  user.GetID(),
					FinishTime:              finishTime,
					TimeSavingPercentage:    0,
					MineralSavingPercentage: 0,
				}

				blueWorks.BlueWorks.Add(&newWork)
			}

			storages.Storages.RemoveItemBySlot(user.GetID(), user.InBaseID, msg.StorageSlot, msg.Count)
			for _, slot := range recyclerItems {
				storages.Storages.RemoveItem(user.GetID(), user.InBaseID, slot.ItemID, slot.Type, slot.Quantity)
			}

			storage.Updater(user.GetID())
		}
	}
}

func selectWork(ws *websocket.Conn, msg Message) {
	user := usersLobbyWs[ws]

	if user != nil {
		if msg.ID > 0 {
			work := blueWorks.BlueWorks.GetByID(msg.ID)
			if work != nil && work.UserID == user.GetID() {

				bp, _ := gameTypes.BluePrints.GetByID(work.BlueprintID)
				percentRemainResource := 100 - (work.GetDonePercent() + work.MineralSavingPercentage)

				returnItems := make([]*inventory.Slot, 0)
				lobby.ParseItems(&returnItems, percentRemainResource, bp, 1)

				lobbyPipe <- Message{Event: msg.Event, UserID: user.GetID(), PreviewRecycleSlots: returnItems,
					BluePrint: bp, BPItem: gameTypes.BluePrints.GetItemsByBluePrintID(bp.ID), Count: bp.Count,
					StorageSlot: msg.StorageSlot, ID: msg.ID, BlueWork: work}
			}
		} else {
			works := blueWorks.BlueWorks.GetSameWorks(
				msg.BluePrintID,
				msg.MineralSaving,
				msg.TimeSaving,
				user.GetID(),
				user.InBaseID,
				msg.ToTime/1000,
				msg.StartTime/1000,
			)

			returnItems := make([]*inventory.Slot, 0)
			bp, _ := gameTypes.BluePrints.GetByID(msg.BluePrintID)

			count := 0
			for _, work := range works {
				if work != nil && work.UserID == user.GetID() && msg.Count > 0 {
					percentRemainResource := 100 - work.MineralSavingPercentage
					lobby.ParseItems(&returnItems, percentRemainResource, bp, 1)
					msg.Count--
					count++
					//я очень хочу спать и очень не хочу думать Х(
				}
			}

			lobbyPipe <- Message{Event: "SelectWork", UserID: user.GetID(), PreviewRecycleSlots: returnItems,
				BluePrint: bp, BPItem: gameTypes.BluePrints.GetItemsByBluePrintID(bp.ID), Count: bp.Count * count,
				StorageSlot: msg.StorageSlot, ID: msg.ID, BluePrintID: msg.BluePrintID, MineralSaving: msg.MineralSaving,
				TimeSaving: msg.TimeSaving, ToTime: msg.ToTime, StartTime: msg.StartTime, MaxCount: len(works)}
		}
	}
}

func cancelCraft(ws *websocket.Conn, msg Message) {
	user := usersLobbyWs[ws]

	if user != nil {
		if msg.ID > 0 {
			work := blueWorks.BlueWorks.GetByID(msg.ID)
			if work != nil && work.UserID == user.GetID() {

				bp, _ := gameTypes.BluePrints.GetByID(work.BlueprintID)
				percentRemainResource := 100 - (work.GetDonePercent() + work.MineralSavingPercentage)

				returnItems := make([]*inventory.Slot, 0)
				lobby.ParseItems(&returnItems, percentRemainResource, bp, 1)

				for _, item := range returnItems {
					storages.Storages.AddSlot(user.GetID(), user.InBaseID, item)
				}

				blueWorks.BlueWorks.Remove(work)
			}
		} else {

			works := blueWorks.BlueWorks.GetSameWorks(
				msg.BluePrintID,
				msg.MineralSaving,
				msg.TimeSaving,
				user.GetID(),
				user.InBaseID,
				msg.ToTime/1000,
				msg.StartTime/1000,
			)

			returnItems := make([]*inventory.Slot, 0)
			bp, _ := gameTypes.BluePrints.GetByID(msg.BluePrintID)

			for _, work := range works {
				if work != nil && work.UserID == user.GetID() && msg.Count > 0 {
					percentRemainResource := 100 - work.MineralSavingPercentage
					lobby.ParseItems(&returnItems, percentRemainResource, bp, 1)
					blueWorks.BlueWorks.Remove(work)
					msg.Count--
				}
			}

			for _, item := range returnItems {
				storages.Storages.AddSlot(user.GetID(), user.InBaseID, item)
			}

			msg.Count = len(works)
			selectWork(ws, msg)
		}
		storage.Updater(user.GetID())
	}
}
