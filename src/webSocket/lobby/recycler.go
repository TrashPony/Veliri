package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
	wsInventory "github.com/TrashPony/Veliri/src/webSocket/inventory"
)

func placeItemToProcessor(user *player.Player, msg Message, recycleItems *map[int]*lobby.RecycleItem) {

	baseStorage, find := storages.Storages.Get(user.GetID(), user.InBaseID)
	userBase, findBase := bases.Bases.Get(user.InBaseID)

	if find && findBase {
		if *recycleItems == nil {
			*recycleItems = make(map[int]*lobby.RecycleItem)
		}

		if msg.Event == "PlaceItemToProcessor" {
			storageSlot, ok := baseStorage.Slots[msg.StorageSlot]
			if ok {
				(*recycleItems)[msg.StorageSlot] = &lobby.RecycleItem{Slot: storageSlot, Recycled: false}
			}
		}

		if msg.Event == "PlaceItemsToProcessor" {
			for _, itemSlot := range msg.StorageSlots {
				storageSlot, ok := baseStorage.Slots[itemSlot]
				if ok {
					(*recycleItems)[itemSlot] = &lobby.RecycleItem{Slot: storageSlot, Recycled: false}
				}
			}
		}

		resultItems, _ := lobby.GetRecycleItems(recycleItems, user, userBase)

		lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
			PreviewRecycleSlots: resultItems}
	}
}

func removeItemToProcessor(user *player.Player, msg Message, recycleItems *map[int]*lobby.RecycleItem) {

	userBase, _ := bases.Bases.Get(user.InBaseID)

	if msg.Event == "RemoveItemFromProcessor" {
		delete(*recycleItems, msg.RecyclerSlot)
	}

	if msg.Event == "RemoveItemsFromProcessor" {
		for _, itemSlot := range msg.StorageSlots {
			delete(*recycleItems, itemSlot)
		}
	}

	resultItems, _ := lobby.GetRecycleItems(recycleItems, user, userBase)

	lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
		PreviewRecycleSlots: resultItems}
}

func recycle(user *player.Player, msg Message, recycleItems *map[int]*lobby.RecycleItem) {

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
}
