package updateSquad

import "../../gameObjects/unit"

func UpdateBodyUnit(unit *unit.Unit)  {
	body := unit.GetBody()

	for _, slot := range body.Equip {
		if slot.Equip == nil {

		}

		if slot.InsertToDB && slot != nil {

		}

		if !slot.InsertToDB && slot != nil {

		}

		if body.Weapon == nil && slot.Weapon {

		}
	}
}