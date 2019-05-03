package squad_inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func ItemToStorage(user *player.Player, inventorySlot int) error {
	if user.InBaseID > 0 {
		slot := user.GetSquad().Inventory.Slots[inventorySlot]

		if slot == nil {
			return errors.New("no find slot")
		}

		ok := storages.Storages.AddItem(user.GetID(), user.InBaseID, slot.Item, slot.Type, slot.ItemID, slot.Quantity,
			slot.HP, slot.Size/float32(slot.Quantity), slot.MaxHP, false)

		if ok {
			user.GetSquad().Inventory.Slots[inventorySlot].RemoveItemBySlot(slot.Quantity)
		}

		go update.Squad(user.GetSquad(), true)
		return nil
	} else {
		return errors.New("user not in base")
	}
}
