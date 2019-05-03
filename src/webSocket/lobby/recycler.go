package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
	wsInventory "github.com/TrashPony/Veliri/src/webSocket/inventory"
)

func placeItemToProcessor(user *player.Player, msg Message, recycleItems *map[int]*lobby.RecycleItem) {

	baseStorage, find := storages.Storages.Get(user.GetID(), user.InBaseID)

	if find {
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

		lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
			PreviewRecycleSlots: lobby.GetRecycleItems(recycleItems)}
	}
}

func removeItemToProcessor(user *player.Player, msg Message, recycleItems *map[int]*lobby.RecycleItem) {

	if msg.Event == "RemoveItemFromProcessor" {
		delete(*recycleItems, msg.RecyclerSlot)
	}

	if msg.Event == "RemoveItemsFromProcessor" {
		for _, itemSlot := range msg.StorageSlots {
			delete(*recycleItems, itemSlot)
		}
	}

	lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
		PreviewRecycleSlots: lobby.GetRecycleItems(recycleItems)}
}

func recycle(user *player.Player, msg Message, recycleItems *map[int]*lobby.RecycleItem) {
	err := lobby.Recycle(user, recycleItems)

	if err != nil {
		lobbyPipe <- Message{Event: msg.Event, Error: err.Error()}
		return
	}

	lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
		PreviewRecycleSlots: lobby.GetRecycleItems(recycleItems)}

	wsInventory.UpdateStorage(user.GetID())
}
