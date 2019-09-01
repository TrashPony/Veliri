package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
	wsInventory "github.com/TrashPony/Veliri/src/webSocket/inventory"
)

func placeItemToProcessor(user *player.Player, msg Message, recycleItems *map[string]map[int]*lobby.RecycleItem) {

	baseStorage, find := storages.Storages.Get(user.GetID(), user.InBaseID)
	userBase, findBase := bases.Bases.Get(user.InBaseID)

	if find && findBase {

		if *recycleItems == nil {
			*recycleItems = make(map[string]map[int]*lobby.RecycleItem)
		}

		if (*recycleItems)[msg.ItemSource] == nil {
			(*recycleItems)[msg.ItemSource] = make(map[int]*lobby.RecycleItem)
		}

		if msg.ItemSource == "storage" {
			if msg.Event == "PlaceItemToProcessor" {
				// todo не передовать элементы как единичку, "ну что это блять" (с)
				storageSlot, ok := baseStorage.Slots[msg.StorageSlot]
				if ok {
					(*recycleItems)[msg.ItemSource][msg.StorageSlot] = &lobby.RecycleItem{Slot: storageSlot, Recycled: false, Source: msg.ItemSource}
				}
			}

			if msg.Event == "PlaceItemsToProcessor" {
				for _, itemSlot := range msg.StorageSlots {
					storageSlot, ok := baseStorage.Slots[itemSlot]
					if ok {
						(*recycleItems)[msg.ItemSource][itemSlot] = &lobby.RecycleItem{Slot: storageSlot, Recycled: false, Source: msg.ItemSource}
					}
				}
			}
		} else {
			if msg.ItemSource == "squadInventory" {
				if msg.Event == "PlaceItemToProcessor" {
					inventorySlot, ok := user.GetSquad().MatherShip.Inventory.Slots[msg.StorageSlot]
					if ok {
						(*recycleItems)[msg.ItemSource][msg.StorageSlot] = &lobby.RecycleItem{Slot: inventorySlot, Recycled: false, Source: msg.ItemSource}
					}
				}
				if msg.Event == "PlaceItemsToProcessor" {
					for _, itemSlot := range msg.StorageSlots {
						inventorySlot, ok := user.GetSquad().MatherShip.Inventory.Slots[itemSlot]
						if ok {
							(*recycleItems)[msg.ItemSource][itemSlot] = &lobby.RecycleItem{Slot: inventorySlot, Recycled: false, Source: msg.ItemSource}
						}
					}
				}
			}
		}

		resultItems, _ := lobby.GetRecycleItems(recycleItems, user, userBase)

		lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
			PreviewRecycleSlots: resultItems, UserRecycleSkill: 25 - user.CurrentSkills["processing"].Level*5}
	}
}

func removeItemToProcessor(user *player.Player, msg Message, recycleItems *map[string]map[int]*lobby.RecycleItem) {

	userBase, _ := bases.Bases.Get(user.InBaseID)

	if msg.Event == "RemoveItemFromProcessor" {
		if (*recycleItems)[msg.ItemSource] != nil {
			delete((*recycleItems)[msg.ItemSource], msg.RecyclerSlot)
		}
	}

	if msg.Event == "RemoveItemsFromProcessor" {
		//for _, itemSlot := range msg.StorageSlots {
		//	//delete(*recycleItems, itemSlot) todo source
		//}
	}

	resultItems, _ := lobby.GetRecycleItems(recycleItems, user, userBase)

	lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
		PreviewRecycleSlots: resultItems}
}

func recycle(user *player.Player, msg Message, recycleItems *map[string]map[int]*lobby.RecycleItem) {

	userBase, _ := bases.Bases.Get(user.InBaseID)

	err := lobby.Recycle(user, recycleItems, userBase)

	if err != nil {
		lobbyPipe <- Message{Event: msg.Event, Error: err.Error()}
		return
	}

	resultItems, _ := lobby.GetRecycleItems(recycleItems, user, userBase)

	lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
		PreviewRecycleSlots: resultItems}

	wsInventory.UpdateStorage(user.GetID())
	wsInventory.UpdateInventory(user.GetID())
}
