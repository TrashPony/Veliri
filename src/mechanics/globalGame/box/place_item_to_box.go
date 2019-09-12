package box

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func PlaceItemToBox(user *player.Player, boxID, inventorySlot int) (error, *boxInMap.Box) {
	err, mapBox, mx := checkUseBox(user, boxID)
	if err != nil {
		return err, nil
	}
	defer mx.Unlock()

	slot, ok := user.GetSquad().MatherShip.Inventory.Slots[inventorySlot]
	if ok && slot.Item != nil && mapBox.CapacitySize >= mapBox.GetStorage().GetSize()+slot.Size {

		placeOk := mapBox.GetStorage().AddItemFromSlot(slot, user.GetID())

		if placeOk {
			slot.RemoveItemBySlot(slot.Quantity)
			go update.Squad(user.GetSquad(), true)
			go boxes.Boxes.UpdateBox(mapBox)
			return nil, mapBox
		} else {
			return errors.New("unknown error"), nil
		}

	} else {
		if !ok || slot.Item == nil {
			return errors.New("no find box slot"), nil
		}
		if mapBox.CapacitySize < mapBox.GetStorage().GetSize()+slot.Size {
			return errors.New("weight exceeded"), nil
		}
	}
	return errors.New("unknown error"), nil
}
