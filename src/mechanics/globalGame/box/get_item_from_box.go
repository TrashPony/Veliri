package box

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func GetItemFromBox(user *player.Player, boxID, boxSlot int) (error, *boxInMap.Box) {
	err, mapBox, mx := checkUseBox(user, boxID)
	if err != nil {
		return err, nil
	}
	defer mx.Unlock()

	slot, ok := mapBox.GetStorage().Slots[boxSlot]

	if ok && slot.Item != nil && user.GetSquad().MatherShip.Body.CapacitySize >= user.GetSquad().MatherShip.Inventory.GetSize()+slot.Size {
		placeOk := user.GetSquad().MatherShip.Inventory.AddItemFromSlot(slot, user.GetID())
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
		if user.GetSquad().MatherShip.Body.CapacitySize < user.GetSquad().MatherShip.Inventory.GetSize()+slot.Size {
			return errors.New("weight exceeded"), nil
		}
	}

	return errors.New("unknown error"), nil
}
