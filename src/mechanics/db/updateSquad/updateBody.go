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

func UpdateBody(unit BodyEquipper, squadID int, tableName string) {
	body := unit.GetBody()

	/* обновляем оборудование */
	updateEquipping(body.EquippingI, squadID, tableName, unit.GetID())
	updateEquipping(body.EquippingII, squadID, tableName, unit.GetID())
	updateEquipping(body.EquippingIII, squadID, tableName, unit.GetID())
	updateEquipping(body.EquippingIV, squadID, tableName, unit.GetID())
	updateEquipping(body.EquippingV, squadID, tableName, unit.GetID())

	/* обновляем оружие и патроны */
	for i, slot := range body.Weapons {
		if slot.Weapon == nil {
			_, err := dbConnect.GetDBConnect().Exec("DELETE FROM "+tableName+" WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3",
				unit.GetID(), slot.Number, squadID)
			if err != nil {
				log.Fatal("delete unit body weapon slot " + err.Error())
			}

			delete(body.Weapons, i)
		}

		if slot.InsertToDB && slot.Weapon != nil {
			_, err := dbConnect.GetDBConnect().Exec("INSERT INTO "+tableName+" (id_squad, type, id_squad_unit, id_equipping, slot_in_body) VALUES ($1, $2, $3, $4, $5)",
				squadID, "weapon", unit.GetID(), slot.Weapon.ID, slot.Number)
			if err != nil {
				log.Fatal("insert unit body weapon slot " + err.Error())
			}

			if slot.Ammo != nil {
				_, err := dbConnect.GetDBConnect().Exec("INSERT INTO "+tableName+" (id_squad, type, id_squad_unit, id_equipping, slot_in_body) VALUES ($1, $2, $3, $4, $5)",
					squadID, "ammo", unit.GetID(), slot.Ammo.ID, slot.Number)
				if err != nil {
					log.Fatal("insert unit body ammo slot " + err.Error())
				}
			}
		}

		if !slot.InsertToDB && slot.Weapon != nil {
			_, err := dbConnect.GetDBConnect().Exec("UPDATE "+tableName+" SET type = $1, id_equipping = $2 WHERE id_squad = $3 AND id_squad_unit = $4 AND slot_in_body = $5",
				"weapon", slot.Weapon.ID, squadID, unit.GetID(), slot.Number)
			if err != nil {
				log.Fatal("update unit body weapon slot " + err.Error())
			}

			_, err = dbConnect.GetDBConnect().Exec("DELETE FROM "+tableName+" WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3 AND type = $4",
				unit.GetID(), slot.Number, squadID, "ammo")
			if err != nil {
				log.Fatal("delete unit body weapon slot " + err.Error())
			}

			if slot.Ammo != nil {
				_, err := dbConnect.GetDBConnect().Exec("INSERT INTO "+tableName+" (id_squad, type, id_squad_unit, id_equipping, slot_in_body) VALUES ($1, $2, $3, $4, $5)",
					squadID, "ammo", unit.GetID(), slot.Ammo.ID, slot.Number)
				if err != nil {
					log.Fatal("insert unit body ammo slot " + err.Error())
				}
			}
		}
	}
}

func updateEquipping(Equipping map[int]*detail.BodyEquipSlot, squadID int, tableName string, unitID int) {
	for i, slot := range Equipping {
		if slot.Equip == nil {
			_, err := dbConnect.GetDBConnect().Exec("DELETE FROM "+tableName+" WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3",
				unitID, slot.Number, squadID)
			if err != nil {
				log.Fatal("delete unit body equip slot " + err.Error())
			}
			delete(Equipping, i)
		}

		if slot.InsertToDB && slot.Equip != nil {
			_, err := dbConnect.GetDBConnect().Exec("INSERT INTO "+tableName+" (id_squad, type, type_slot, id_squad_unit, id_equipping, slot_in_body) VALUES ($1, $2, $3, $4, $5)",
				squadID, "equip", slot.Equip.TypeSlot, unitID, slot.Equip.ID, slot.Number)
			if err != nil {
				log.Fatal("insert unit body equip slot " + err.Error())
			}
		}

		if !slot.InsertToDB && slot.Equip != nil {
			_, err := dbConnect.GetDBConnect().Exec("UPDATE " + tableName + " SET type = $1, type_slot = $2, id_equipping = $3 "+
				" WHERE id_squad = $4 AND id_squad_unit = $5 AND slot_in_body = $6",
				"equip", slot.Equip.TypeSlot, slot.Equip.ID, squadID, unitID, slot.Number)
			if err != nil {
				log.Fatal("update unit body equip slot " + err.Error())
			}
		}
	}
}
