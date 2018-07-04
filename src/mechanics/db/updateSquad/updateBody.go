package updateSquad

import (
	"../../gameObjects/detail"
	"../../../dbConnect"
	"log"
)

type BodyEquipper interface {
	GetBody() *detail.Body
	GetID() int
}

func UpdateBody(unit BodyEquipper, squadID int, tableName string)  {
	body := unit.GetBody()

	/* обновляем оборудование */
	for _, slot := range body.Equipping {
		if slot.Equip == nil {
			_, err := dbConnect.GetDBConnect().Exec("DELETE FROM " + tableName + " WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3",
				unit.GetID(), slot.Number, squadID)
			if err != nil {
				log.Fatal("delete unit body equip slot " + err.Error())
			}
		}

		if slot.InsertToDB && slot.Equip != nil {
			_, err := dbConnect.GetDBConnect().Exec("INSERT INTO " + tableName + " (id_squad, type, id_squad_unit, id_equipping, slot_in_body) VALUES ($1, $2, $3, $4, $5)",
				squadID, "equip", unit.GetID(), slot.Equip.ID, slot.Number)
			if err != nil {
				log.Fatal("insert unit body equip slot " + err.Error())
			}
		}

		if !slot.InsertToDB && slot.Equip != nil {
			_, err := dbConnect.GetDBConnect().Exec("UPDATE " + tableName + " SET type = $1, id_equipping = $2 WHERE id_squad = $3 AND id_squad_unit = $4 AND slot_in_body = $5",
				"equip", slot.Equip.ID, squadID, unit.GetID(), slot.Number)
			if err != nil {
				log.Fatal("update unit body equip slot " + err.Error())
			}
		}
	}

	/* обновляем оружие и патроны */
	for _, slot := range body.Weapons{
		if slot.Weapon == nil {
			_, err := dbConnect.GetDBConnect().Exec("DELETE FROM " + tableName + " WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3",
				unit.GetID(), slot.Number, squadID)
			if err != nil {
				log.Fatal("delete unit body weapon slot " + err.Error())
			}
		}

		if slot.InsertToDB && slot.Weapon != nil {
			_, err := dbConnect.GetDBConnect().Exec("INSERT INTO " + tableName + " (id_squad, type, id_squad_unit, id_equipping, slot_in_body) VALUES ($1, $2, $3, $4, $5)",
				squadID, "weapon", unit.GetID(), slot.Weapon.ID, slot.Number)
			if err != nil {
				log.Fatal("insert unit body weapon slot " + err.Error())
			}

			if slot.Ammo != nil {
				_, err := dbConnect.GetDBConnect().Exec("INSERT INTO " + tableName + " (id_squad, type, id_squad_unit, id_equipping, slot_in_body) VALUES ($1, $2, $3, $4, $5)",
					squadID, "ammo", unit.GetID(), slot.Ammo.ID, slot.Number)
				if err != nil {
					log.Fatal("insert unit body ammo slot " + err.Error())
				}
			}
		}

		if !slot.InsertToDB && slot.Weapon != nil {
			_, err := dbConnect.GetDBConnect().Exec("UPDATE " + tableName + " SET type = $1, id_equipping = $2 WHERE id_squad = $3 AND id_squad_unit = $4 AND slot_in_body = $5",
				"weapon", slot.Weapon.ID, squadID, unit.GetID(), slot.Number)
			if err != nil {
				log.Fatal("update unit body weapon slot " + err.Error())
			}

			_, err = dbConnect.GetDBConnect().Exec("DELETE FROM " + tableName + " WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3 AND type = $4",
				unit.GetID(), slot.Number, squadID, "ammo")
			if err != nil {
				log.Fatal("delete unit body weapon slot " + err.Error())
			}

			if slot.Ammo != nil {
				_, err := dbConnect.GetDBConnect().Exec("INSERT INTO " + tableName + " (id_squad, type, id_squad_unit, id_equipping, slot_in_body) VALUES ($1, $2, $3, $4, $5)",
					squadID, "ammo", unit.GetID(), slot.Ammo.ID, slot.Number)
				if err != nil {
					log.Fatal("insert unit body ammo slot " + err.Error())
				}
			}
		}
	}
}