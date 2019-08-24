package inventory

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/gorilla/websocket"
)

func combineItems(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	var srcSlot *inventory.Slot
	var dstSlot *inventory.Slot

	if user.GetSquad() == nil || user.GetSquad().MatherShip.Body == nil || user.InBaseID == 0 {
		return
	}

	if msg.Source == "squadInventory" {
		srcSlot, _ = user.GetSquad().Inventory.Slots[msg.SrcSlot]
	}

	if msg.Source == "storage" {
		userStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)
		srcSlot, _ = userStorage.Slots[msg.SrcSlot]
	}

	if msg.Destination == "squadInventory" {
		dstSlot, _ = user.GetSquad().Inventory.Slots[msg.DstSlot]
	}

	if msg.Destination == "storage" {
		userStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)
		dstSlot, _ = userStorage.Slots[msg.DstSlot]
	}

	if srcSlot != nil && dstSlot != nil && srcSlot.Type == dstSlot.Type && srcSlot.ItemID == srcSlot.ItemID {

		add := false

		if msg.Destination == "squadInventory" {
			// если источник и приемник находятся в 1 инвентаре то нарватся на перигруз невозможно
			if msg.Source == "squadInventory" {
				add = true
				user.GetSquad().Inventory.Slots[msg.DstSlot].AddItemBySlot(srcSlot.Quantity, user.GetID())
			} else {
				// проверка на перегруз
				if user.GetSquad().MatherShip.Body.CapacitySize >= user.GetSquad().Inventory.GetSize()+dstSlot.Size {
					add = true
					user.GetSquad().Inventory.Slots[msg.DstSlot].AddItemBySlot(srcSlot.Quantity, user.GetID())
				} else {
					add = false
				}
			}
		}

		if msg.Destination == "storage" {
			add = true
			storages.Storages.AddItemBySlot(user.GetID(), user.InBaseID, msg.DstSlot, srcSlot.Quantity)
		}

		if msg.Source == "squadInventory" && add {
			user.GetSquad().Inventory.Slots[msg.SrcSlot].RemoveItemBySlot(srcSlot.Quantity)
		}

		if msg.Source == "storage" && add {
			storages.Storages.RemoveItemBySlot(user.GetID(), user.InBaseID, msg.SrcSlot, srcSlot.Quantity)
		}
	}
	UpdateSquad("UpdateSquad", user, nil, ws, msg)
}
