package squadInventory

import (
	new "../db/squad"
	"../db/squad/update"
	"../factories/bases"
	"../factories/gameTypes"
	"../gameObjects/unit"
	"../player"
	"errors"
)

func SetMSBody(user *player.Player, idBody, inventorySlot int, source string) error {
	if user.InBaseID > 0 {

		slot := getSlotBySource(user, inventorySlot, source)

		if slot != nil && slot.ItemID == idBody && slot.Type == "body" {
			newBody, _ := gameTypes.Bodies.GetByID(idBody)

			if !newBody.MotherShip {
				return errors.New("wrong type body")
			}

			_, newSquad := new.AddNewSquad(newBody.Name, user.GetID()) // делаем новый отряд

			base, _ := bases.Bases.Get(user.InBaseID)
			newSquad.MapID = base.MapID

			if user.GetSquad() != nil {
				user.GetSquad().Active = false         //  старый отряд делаем не активным
				user.GetSquad().BaseID = user.InBaseID // ид базы где храниться отряд
				update.Squad(user.GetSquad(), true)    // обновляем старый отряд в бд
			}

			GetInventory(user)

			if user.GetSquad().MatherShip == nil {
				user.GetSquad().MatherShip = &unit.Unit{}
			}

			user.GetSquad().MatherShip.HP = slot.HP                 // устанавливает колво хп как у тела
			user.GetSquad().MatherShip.Power = newBody.MaxPower     // устанавливаем мощьность как у тела
			user.GetSquad().MatherShip.ActionPoints = newBody.Speed // устанавливаем скорость как у тела
			user.GetSquad().MatherShip.Body = newBody

			RemoveSlotBySource(user, inventorySlot, source, 1)

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

func SetUnitBody(user *player.Player, idBody, inventorySlot, numberUnitSlot int, source string) error {
	if user.InBaseID > 0 {

		slot := getSlotBySource(user, inventorySlot, source)

		if user.GetSquad().MatherShip == nil || user.GetSquad().MatherShip.Body == nil {
			return errors.New("no ms")
		}

		unitSlot, ok := user.GetSquad().MatherShip.Body.EquippingIV[numberUnitSlot]
		if !ok {
			return errors.New("wrong slot ms")
		}

		if slot != nil && slot.ItemID == idBody && slot.Type == "body" {
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
						RemoveUnitBody(user, numberUnitSlot, true)
						unitSlot.Unit = &unit.Unit{}
					}

					unitSlot.Unit.HP = slot.HP                 // устанавливает колво хп как у тела
					unitSlot.Unit.Power = newBody.MaxPower     // устанавливаем мощьность как у тела
					unitSlot.Unit.ActionPoints = newBody.Speed // устанавливаем скорость как у тела

					RemoveSlotBySource(user, inventorySlot, source, 1)
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
