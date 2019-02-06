package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
	"github.com/TrashPony/Veliri/src/webSocket/storage"
	"github.com/gorilla/websocket"
)

func placeItemToProcessor(ws *websocket.Conn, msg Message, recycleItems *map[int]*lobby.RecycleItem) {
	user := usersLobbyWs[ws]

	baseStorage, find := storages.Storages.Get(user.GetID(), user.InBaseID)

	if user != nil && find {
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

func removeItemToProcessor(ws *websocket.Conn, msg Message, recycleItems *map[int]*lobby.RecycleItem) {
	user := usersLobbyWs[ws]
	if user != nil {

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
}

func recycle(ws *websocket.Conn, msg Message, recycleItems *map[int]*lobby.RecycleItem) {
	user := usersLobbyWs[ws]
	if user != nil {
		err := lobby.Recycle(user, recycleItems)

		if err != nil {
			lobbyPipe <- Message{Event: msg.Event, Error: err.Error()}
			return
		}

		lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
			PreviewRecycleSlots: lobby.GetRecycleItems(recycleItems)}

		storage.Updater(user.GetID())
	}
}
