package globalGame

import (
	"../../mechanics/db/squad/update"
	"../../mechanics/factories/boxes"
	"../../mechanics/factories/maps"
	"../../mechanics/gameObjects/box"
	"../../mechanics/gameObjects/inventory"
	"../../mechanics/player"
	"errors"
	"math"
	"math/rand"
	"time"
)

func ThrowItems(user *player.Player, slots []inventory.Slot) (error, *box.Box) {

	mp, find := maps.Maps.GetByID(user.GetSquad().MapID)
	if !find {
		return errors.New("no map"), nil
	}

	// берем координату позади отряда смотрим что бы она была пустая
	radRotate := float64(user.GetSquad().MatherShip.Rotate) * math.Pi / 180
	stopX := float64(130) * math.Cos(radRotate) // идем по вектору движения корпуса
	stopY := float64(130) * math.Sin(radRotate)

	forecastX := float64(user.GetSquad().GlobalX) - stopX // - т.к. нам нужна точка позади
	forecastY := float64(user.GetSquad().GlobalY) - stopY

	hexCoordinate := GetQRfromXY(int(forecastX), int(forecastY), mp)
	if hexCoordinate.Move {
		newBox := box.Box{Q: hexCoordinate.Q, R: hexCoordinate.R, Rotate: rand.Intn(360), MapID: mp.Id, TypeID: 1,
			DestroyTime: time.Now()}

		newBox.GetStorage().Slots = make(map[int]*inventory.Slot)

		createBox := false
		for i, slot := range slots {
			if slot.Item != nil {
				realSlot, ok := user.GetSquad().Inventory.Slots[i]
				if ok {
					addOk := newBox.GetStorage().AddItemFromSlot(realSlot)
					if addOk {
						createBox = true
						realSlot.RemoveItemBySlot(realSlot.Quantity)
					}
				}
			}
		}

		if createBox {
			update.Squad(user.GetSquad(), true)
			return nil, boxes.Boxes.InsertNewBox(&newBox)
		} else {
			return errors.New("not find items"), nil
		}
	} else {
		return errors.New("not allow place"), nil
	}

	return errors.New("unknown error"), nil
}
