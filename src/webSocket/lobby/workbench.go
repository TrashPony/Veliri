package lobby

import (
	"../../mechanics/factories/blueWorks"
	"../../mechanics/factories/gameTypes"
	"../../mechanics/factories/storages"
	"../../mechanics/gameObjects/blueprints"
	"../../mechanics/gameObjects/inventory"
	"../../mechanics/lobby"
	"../storage"
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
				BluePrint: bluePrint, BPItem: gameTypes.BluePrints.GetItems(slot.ItemID), Count: msg.Count,
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

func cancelCraft(ws *websocket.Conn, msg Message) {
	user := usersLobbyWs[ws]

	if user != nil {
		work := blueWorks.BlueWorks.GetByID(msg.ID)
		if work.UserID == user.GetID() {
			bp, _ := gameTypes.BluePrints.GetByID(work.BlueprintID)

			realMineralsPercent := 100 - work.MineralSavingPercentage

			// TODO проверять сколько времени в % соотношение крафт делался и отнять эти проценту из realMineralsPercent
			craftItems := make([]*inventory.Slot, 0)
			lobby.ParseItems(&craftItems, realMineralsPercent, bp, 1)
			// TODO возвращать итемы для крафта

			// удалить работы
			blueWorks.BlueWorks.Remove(work)
		}
	}
}
