package squadInventory

import (
	"../db/squad/update"
	"../factories/gameTypes"
	"../gameObjects/unit"
	"../player"
	"errors"
)

func SetMSBody(user *player.Player, idBody, inventorySlot int) error {
	if user.InBaseID > 0 {
		body := user.GetSquad().Inventory.Slots[inventorySlot]

		if body != nil && body.ItemID == idBody && body.Type == "body" {
			newBody, _ := gameTypes.Bodies.GetByID(idBody)

			if !newBody.MotherShip {
				return errors.New("wrong type body")
			}

			if user.GetSquad().MatherShip == nil {
				user.GetSquad().MatherShip = &unit.Unit{}
			} else {
				if user.GetSquad().MatherShip.Body != nil {
					BodyRemove(user, user.GetSquad().MatherShip)
					user.GetSquad().MatherShip.Body = nil
				}
			}

			user.GetSquad().MatherShip.HP = body.HP                 // устанавливает колво хп как у тела
			user.GetSquad().MatherShip.Power = newBody.MaxPower     // устанавливаем мощьность как у тела
			user.GetSquad().MatherShip.ActionPoints = newBody.Speed // устанавливаем скорость как у тела

			user.GetSquad().Inventory.Slots[inventorySlot].RemoveItemBySlot(1)
			user.GetSquad().MatherShip.Body = newBody

			user.GetSquad().MatherShip.Units = make(map[int]*unit.Slot) // заполняем ячейки юнитов

			for _, slot := range user.GetSquad().MatherShip.Body.EquippingIV {
				unitSlot := unit.Slot{}
				unitSlot.NumberSlot = slot.Number
				unitSlot.Unit = nil
				user.GetSquad().MatherShip.Units[slot.Number] = &unitSlot
			}

			user.GetSquad().MatherShip.CalculateParams()

			go update.Squad(user.GetSquad(), true)
			return nil
		} else {
			return errors.New("wrong inventory slot")
		}
	} else {
		return errors.New("not in base")
	}
}

func SetUnitBody(user *player.Player, idBody, inventorySlot, numberUnitSlot int) error {
	if user.InBaseID > 0 {

		body := user.GetSquad().Inventory.Slots[inventorySlot]

		if user.GetSquad().MatherShip == nil || user.GetSquad().MatherShip.Body == nil {
			return errors.New("no ms")
		}

		unitSlot, ok := user.GetSquad().MatherShip.Body.EquippingIV[numberUnitSlot]
		if !ok {
			return errors.New("wrong slot ms")
		}

		if body != nil && body.ItemID == idBody && body.Type == "body" {
			newBody, _ := gameTypes.Bodies.GetByID(idBody)

			if newBody.MotherShip {
				return errors.New("wrong type body")
			}

			if newBody.StandardSize <= unitSlot.StandardSize {
				unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]

				if ok {
					if unitSlot.Unit == nil {
						unitSlot.Unit = &unit.Unit{}
					} else {
						RemoveUnitBody(user, numberUnitSlot)
						unitSlot.Unit = &unit.Unit{}
					}

					unitSlot.Unit.HP = body.HP                 // устанавливает колво хп как у тела
					unitSlot.Unit.Power = newBody.MaxPower     // устанавливаем мощьность как у тела
					unitSlot.Unit.ActionPoints = newBody.Speed // устанавливаем скорость как у тела

					user.GetSquad().Inventory.Slots[inventorySlot].RemoveItemBySlot(1)
					unitSlot.Unit.Body = newBody

					unitSlot.Unit.CalculateParams()

					go update.Squad(user.GetSquad(), true)
					return nil
				} else {
					return errors.New("wrong slot")
				}
			} else {
				return errors.New("wrong type slot")
			}
		} else {
			return errors.New("wrong inventory slot")
		}
	} else {
		return errors.New("not in base")
	}
}
