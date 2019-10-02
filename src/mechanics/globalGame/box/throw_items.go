package box

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"math"
	"math/rand"
	"time"
)

func ThrowItems(user *player.Player, slots []inventory.Slot) (error, bool, *boxInMap.Box) {

	mp, find := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)
	if !find {
		return errors.New("no map"), false, nil
	}

	// берем координату позади отряда смотрим что бы она была пустая
	radRotate := float64(user.GetSquad().MatherShip.Rotate) * math.Pi / 180
	stopX := float64(75) * math.Cos(radRotate) // идем по вектору движения корпуса
	stopY := float64(75) * math.Sin(radRotate)

	forecastX := user.GetSquad().MatherShip.X - int(stopX) // - т.к. нам нужна точка позади
	forecastY := user.GetSquad().MatherShip.Y - int(stopY)

	boxType, _ := gameTypes.Boxes.GetByID(1)
	newBox := &boxInMap.Box{X: forecastX, Y: forecastY, Rotate: rand.Intn(360), MapID: mp.Id, TypeID: 1,
		DestroyTime: time.Now(), Height: boxType.Height, Width: boxType.Width}

	placeFree, oldBox := collisions.CheckBoxCollision(newBox, mp)
	if !placeFree && oldBox == nil {
		return errors.New("place busy"), false, nil
	}

	if oldBox != nil {
		for i, slot := range slots {
			if slot.Item != nil {
				realSlot, ok := user.GetSquad().MatherShip.Inventory.Slots[i]
				if ok {
					addOk := oldBox.GetStorage().AddItemFromSlot(realSlot, user.GetID())
					if addOk {
						realSlot.RemoveItemBySlot(realSlot.Quantity)
					}
				}
			}
		}
		boxes.Boxes.UpdateBox(oldBox)

		return nil, false, oldBox
	} else {

		newBox.GetStorage().Slots = make(map[int]*inventory.Slot)
		newBox.GetStorage().SetSlotsSize(999)

		createBox := false
		for i, slot := range slots {
			if slot.Item != nil {
				realSlot, ok := user.GetSquad().MatherShip.Inventory.Slots[i]
				if ok {
					createBox = true
					addOk := newBox.GetStorage().AddItemFromSlot(realSlot, user.GetID())
					if addOk {
						realSlot.RemoveItemBySlot(realSlot.Quantity)
					}
				}
			}
		}

		if createBox {
			update.Squad(user.GetSquad(), true)
			return nil, createBox, boxes.Boxes.InsertNewBox(newBox)
		} else {
			return errors.New("not find items"), createBox, nil
		}
	}
}
