package box

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func BoxToBox(user *player.Player, boxID, boxSlot, toBoxID int) (error, *boxInMap.Box, *boxInMap.Box) {
	err, getBox, mx := checkUseBox(user, boxID)
	if err != nil {
		return err, nil, nil
	}
	mx.Unlock() // закрываем сразу т.к. второй ящик всеравно заблокирует работу с ящиками

	err, toBox, mx := checkUseBox(user, toBoxID)
	if err != nil {
		return err, nil, nil
	}
	defer mx.Unlock()

	slot, ok := getBox.GetStorage().Slots[boxSlot]

	if ok && slot.Item != nil && toBox.CapacitySize >= toBox.GetStorage().GetSize()+slot.Size {

		placeOk := toBox.GetStorage().AddItemFromSlot(slot, user.GetID())
		if placeOk {
			slot.RemoveItemBySlot(slot.Quantity)
			go boxes.Boxes.UpdateBox(getBox)
			go boxes.Boxes.UpdateBox(toBox)
			return nil, getBox, toBox
		}
	} else {
		if !ok || slot.Item == nil {
			return errors.New("no find box slot"), nil, nil
		}
		if toBox.CapacitySize < toBox.GetStorage().GetSize()+slot.Size {
			return errors.New("weight exceeded"), nil, nil
		}
	}

	return errors.New("unknown error"), nil, nil
}
