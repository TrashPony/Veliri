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
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
	"math/rand"
	"time"
)

func PlaceNewBox(user *player.Player, numberSlot, password int) (error, *boxInMap.Box) {
	mp, find := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)
	if !find {
		return errors.New("no map"), nil
	}

	// берем координату позади отряда смотрим что бы она была пустая
	radRotate := float64(user.GetSquad().MatherShip.Rotate) * math.Pi / 180
	stopX := float64(65) * math.Cos(radRotate) // идем по вектору движения корпуса
	stopY := float64(65) * math.Sin(radRotate)

	forecastX := float64(user.GetSquad().MatherShip.X) - stopX // - т.к. нам нужна точка позади
	forecastY := float64(user.GetSquad().MatherShip.Y) - stopY

	q, r := game_math.GetQRfromXY(int(forecastX), int(forecastY))
	hexCoordinate, find := mp.OneLayerMap[q][r]
	if !find {
		return errors.New("wrong place"), nil
	}

	oldBox, mx := boxes.Boxes.GetByQR(hexCoordinate.Q, hexCoordinate.R, mp.Id)
	mx.Unlock()

	if oldBox != nil {
		return errors.New("place busy"), nil
	}

	slot, find := user.GetSquad().MatherShip.Inventory.Slots[numberSlot]
	if find && slot != nil && slot.Item != nil && slot.Type == "boxes" {
		typeBox, _ := gameTypes.Boxes.GetByID(slot.ItemID)
		if typeBox != nil {

			slot.RemoveItemBySlot(1)

			newBox := boxInMap.Box{Q: hexCoordinate.Q, R: hexCoordinate.R, Rotate: rand.Intn(360), MapID: mp.Id,
				TypeID: typeBox.TypeID, DestroyTime: time.Now()}

			newBox.GetStorage().Slots = make(map[int]*inventory.Slot)
			newBox.GetStorage().SetSlotsSize(999)

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
}
