package globalGame

import (
	"../../mechanics/db/squad/update"
	"../../mechanics/factories/boxes"
	"../../mechanics/factories/gameTypes"
	"../../mechanics/factories/maps"
	"../../mechanics/gameObjects/inventory"
	"../../mechanics/player"
	"../gameObjects/boxInMap"
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

func PlaceNewBox(user *player.Player, numberSlot, password int) (error, *boxInMap.Box) {
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

	oldBox, mx := boxes.Boxes.GetByQR(hexCoordinate.Q, hexCoordinate.R, mp.Id)
	mx.Unlock()

	if oldBox != nil {
		return errors.New("place busy"), nil
	}

	slot, find := user.GetSquad().Inventory.Slots[numberSlot]
	if find && slot != nil && slot.Item != nil && slot.Type == "boxes" {
		typeBox, _ := gameTypes.Boxes.GetByID(slot.ItemID)
		if typeBox != nil {

			slot.RemoveItemBySlot(1)

			newBox := boxInMap.Box{Q: hexCoordinate.Q, R: hexCoordinate.R, Rotate: rand.Intn(360), MapID: mp.Id,
				TypeID: typeBox.TypeID, DestroyTime: time.Now()}

			newBox.GetStorage().Slots = make(map[int]*inventory.Slot)

			if typeBox.Protect {
				newBox.SetPassword(password)
			}

			update.Squad(user.GetSquad(), true)
			return nil, boxes.Boxes.InsertNewBox(&newBox)
		} else {
			return errors.New("box type not find"), nil
		}
	} else {
		return errors.New("inventory slot not find"), nil
	}
	return errors.New("unknown error"), nil
}

func ThrowItems(user *player.Player, slots []inventory.Slot) (error, bool, *boxInMap.Box) {

	mp, find := maps.Maps.GetByID(user.GetSquad().MapID)
	if !find {
		return errors.New("no map"), false, nil
	}

	// берем координату позади отряда смотрим что бы она была пустая
	radRotate := float64(user.GetSquad().MatherShip.Rotate) * math.Pi / 180
	stopX := float64(130) * math.Cos(radRotate) // идем по вектору движения корпуса
	stopY := float64(130) * math.Sin(radRotate)

	forecastX := float64(user.GetSquad().GlobalX) - stopX // - т.к. нам нужна точка позади
	forecastY := float64(user.GetSquad().GlobalY) - stopY

	hexCoordinate := GetQRfromXY(int(forecastX), int(forecastY), mp)

	if hexCoordinate.Move {
		oldBox, mx := boxes.Boxes.GetByQR(hexCoordinate.Q, hexCoordinate.R, mp.Id)

		if oldBox != nil {
			for i, slot := range slots {
				if slot.Item != nil {
					realSlot, ok := user.GetSquad().Inventory.Slots[i]
					if ok {
						addOk := oldBox.GetStorage().AddItemFromSlot(realSlot)
						if addOk {
							realSlot.RemoveItemBySlot(realSlot.Quantity)
						}
					}
				}
			}
			mx.Unlock()
			boxes.Boxes.UpdateBox(oldBox)

			return nil, false, oldBox
		} else {
			mx.Unlock()
			newBox := boxInMap.Box{Q: hexCoordinate.Q, R: hexCoordinate.R, Rotate: rand.Intn(360), MapID: mp.Id, TypeID: 1,
				DestroyTime: time.Now()}

			newBox.GetStorage().Slots = make(map[int]*inventory.Slot)

			createBox := false
			for i, slot := range slots {
				if slot.Item != nil {
					realSlot, ok := user.GetSquad().Inventory.Slots[i]
					if ok {
						createBox = true
						addOk := newBox.GetStorage().AddItemFromSlot(realSlot)
						if addOk {
							realSlot.RemoveItemBySlot(realSlot.Quantity)
						}
					}
				}
			}

			if createBox {
				update.Squad(user.GetSquad(), true)
				return nil, createBox, boxes.Boxes.InsertNewBox(&newBox)
			} else {
				return errors.New("not find items"), createBox, nil
			}
		}
	} else {
		return errors.New("not allow place"), false, nil
	}

	return errors.New("unknown error"), false, nil
}

func checkUseBox(user *player.Player, boxID int) (error, *boxInMap.Box, *sync.Mutex) {
	mapBox, mx := boxes.Boxes.Get(boxID)
	if mapBox != nil {
		boxX, boxY := GetXYCenterHex(mapBox.Q, mapBox.R)

		dist := GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, boxX, boxY)
		if dist < 150 {
			return nil, mapBox, mx
		} else {
			mx.Unlock()
			return errors.New("no min dist"), nil, nil
		}
	} else {
		mx.Unlock()
		return errors.New("no find box"), nil, nil
	}
}

func GetItemFromBox(user *player.Player, boxID, boxSlot int) (error, *boxInMap.Box) {
	err, mapBox, mx := checkUseBox(user, boxID)
	defer mx.Unlock()

	if err != nil {
		return err, nil
	}

	slot, ok := mapBox.GetStorage().Slots[boxSlot]

	if ok && slot.Item != nil && user.GetSquad().MatherShip.Body.CapacitySize >= user.GetSquad().Inventory.GetSize()+slot.Size {

		placeOk := user.GetSquad().Inventory.AddItemFromSlot(slot)
		if placeOk {
			slot.RemoveItemBySlot(slot.Quantity)
			update.Squad(user.GetSquad(), true)
			boxes.Boxes.UpdateBox(mapBox) // в дефер что бы он отработал после закрытия мьютекса
			return nil, mapBox
		} else {
			return errors.New("unknown error"), nil
		}

	} else {
		if !ok || slot.Item == nil {
			return errors.New("no find box slot"), nil
		}
		if user.GetSquad().MatherShip.Body.CapacitySize < user.GetSquad().Inventory.GetSize()+slot.Size {
			return errors.New("weight exceeded"), nil
		}
	}

	return errors.New("unknown error"), nil
}

func PlaceItemToBox(user *player.Player, boxID, inventorySlot int) (error, *boxInMap.Box) {
	err, mapBox, mx := checkUseBox(user, boxID)
	defer mx.Unlock()

	if err != nil {
		return err, nil
	}

	slot, ok := user.GetSquad().Inventory.Slots[inventorySlot]
	if ok && slot.Item != nil && mapBox.CapacitySize >= mapBox.GetStorage().GetSize()+slot.Size {

		placeOk := mapBox.GetStorage().AddItemFromSlot(slot)

		if placeOk {
			slot.RemoveItemBySlot(slot.Quantity)
			update.Squad(user.GetSquad(), true)
			boxes.Boxes.UpdateBox(mapBox)
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
