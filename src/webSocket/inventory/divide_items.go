package inventory

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/gorilla/websocket"
)

func divideItems(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if msg.Storage {
		userStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)

		slot, ok := userStorage.Slots[msg.InventorySlot]
		if ok && slot.Quantity > msg.Count {
			if storages.Storages.AddItem(user.GetID(), user.InBaseID, slot.Item, slot.Type, slot.ItemID, msg.Count,
				slot.HP, slot.Size/float32(slot.Quantity), slot.MaxHP, true) {

				storages.Storages.RemoveItemBySlot(user.GetID(), user.InBaseID, msg.InventorySlot, msg.Count)
			}
		}

	} else {
		if user.GetSquad() != nil {

			slot, ok := user.GetSquad().Inventory.Slots[msg.InventorySlot]
			if ok && slot.Quantity > msg.Count {
				if user.GetSquad().Inventory.AddItem(slot.Item, slot.Type, slot.ItemID, msg.Count, slot.HP,
					slot.Size/float32(slot.Quantity), slot.MaxHP, true) {
					slot.RemoveItemBySlot(msg.Count)
				}
			}

			go update.Squad(user.GetSquad(), true)
		}
	}
	UpdateSquad("UpdateSquad", user, nil, ws, msg)
}
