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

func PlaceNewBox(user *player.Player, numberSlot, password int) (error, *boxInMap.Box) {
	mp, find := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)
	if !find {
		return errors.New("no map"), nil
	}

	// берем координату позади отряда смотрим что бы она была пустая
	radRotate := float64(user.GetSquad().MatherShip.Rotate) * math.Pi / 180
	stopX := float64(75) * math.Cos(radRotate) // идем по вектору движения корпуса
	stopY := float64(75) * math.Sin(radRotate)

	forecastX := user.GetSquad().MatherShip.X - int(stopX) // - т.к. нам нужна точка позади
	forecastY := user.GetSquad().MatherShip.Y - int(stopY)

	slot, find := user.GetSquad().MatherShip.Inventory.Slots[numberSlot]
	if find && slot != nil && slot.Item != nil && slot.Type == "boxes" {

		typeBox, _ := gameTypes.Boxes.GetByID(slot.ItemID)
		if typeBox != nil {

			newBox := &boxInMap.Box{X: forecastX, Y: forecastY, Rotate: rand.Intn(360), MapID: mp.Id,
				TypeID: typeBox.TypeID, DestroyTime: time.Now()}

			placeFree, _ := collisions.CheckBoxCollision(newBox, mp, 0)
			if !placeFree {
				return errors.New("place busy"), nil
			}

			slot.RemoveItemBySlot(1)

			newBox.GetStorage().Slots = make(map[int]*inventory.Slot)
			newBox.GetStorage().SetSlotsSize(999)

			if typeBox.Protect {
				newBox.SetPassword(password)
			}

			update.Squad(user.GetSquad(), true)
			return nil, boxes.Boxes.InsertNewBox(newBox)
		} else {
			return errors.New("box type not find"), nil
		}
	} else {
		return errors.New("inventory slot not find"), nil
	}
}
