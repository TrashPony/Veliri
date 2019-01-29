package lobby

import (
	"../../mechanics/factories/storages"
	"../../mechanics/lobby"
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

		lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
			PreviewRecycleSlots: lobby.GetRecycleItems(recycleItems)}
	}
}
