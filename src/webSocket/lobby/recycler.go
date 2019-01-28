package lobby

import (
	"../../mechanics/factories/storages"
	"../../mechanics/gameObjects/inventory"
	"../../mechanics/lobby"
	"github.com/gorilla/websocket"
)

func placeItemToProcessor(ws *websocket.Conn, msg Message, recycleItems *map[int]*inventory.Slot) {
	user := usersLobbyWs[ws]

	baseStorage, find := storages.Storages.Get(user.GetID(), user.InBaseID)

	if user != nil && find {
		if *recycleItems == nil {
			*recycleItems = make(map[int]*inventory.Slot)
		}

		if msg.Event == "PlaceItemToProcessor" {
			storageSlot, ok := baseStorage.Slots[msg.StorageSlot]
			if ok {
				(*recycleItems)[msg.StorageSlot] = storageSlot
			}
		}

		if msg.Event == "PlaceItemsToProcessor" {
			for _, itemSlot := range msg.StorageSlots {
				storageSlot, ok := baseStorage.Slots[itemSlot]
				if ok {
					(*recycleItems)[itemSlot] = storageSlot
				}
			}
		}

		lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
			PreviewRecycleSlots: lobby.GetRecycleItems(recycleItems)}
	}
}

func removeItemToProcessor(ws *websocket.Conn, msg Message, recycleItems *map[int]*inventory.Slot) {
	user := usersLobbyWs[ws]
	if user != nil {

		if msg.Event == "RemoveItemFromProcessor" {
			delete(*recycleItems, msg.RecyclerSlot)
		}

		lobbyPipe <- Message{Event: "updateRecycler", UserID: user.GetID(), RecycleSlots: *recycleItems,
			PreviewRecycleSlots: lobby.GetRecycleItems(recycleItems)}
	}
}
