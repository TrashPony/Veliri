package lobby

import (
	"../../mechanics/factories/gameTypes"
	"../../mechanics/factories/storages"
	"../../mechanics/gameObjects/inventory"
	"../../mechanics/lobby"
	"github.com/gorilla/websocket"
)

func openWorkbench(ws *websocket.Conn, msg Message) {
	user := usersLobbyWs[ws]

	// GET STORAGE INVENTORY
	// todo GET current works
	baseStorage, find := storages.Storages.Get(user.GetID(), user.InBaseID)
	if user != nil && find {
		lobbyPipe <- Message{Event: "WorkbenchStorage", UserID: user.GetID(), Storage: baseStorage}
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
