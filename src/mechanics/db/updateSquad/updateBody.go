package updateSquad

import (
	"../../gameObjects/detail"
	"log"
	"database/sql"
)

type BodyEquipper interface {
	GetBody() *detail.Body
	GetID() int
}

func UpdateBody(unit BodyEquipper, squadID int, tableName string, tx *sql.Tx) {
	body := unit.GetBody()

	/* обновляем оборудование */
	updateEquipping(body.EquippingI, squadID, tableName, unit.GetID(), tx)
	updateEquipping(body.EquippingII, squadID, tableName, unit.GetID(), tx)
	updateEquipping(body.EquippingIII, squadID, tableName, unit.GetID(), tx)
	updateEquipping(body.EquippingIV, squadID, tableName, unit.GetID(), tx)
	updateEquipping(body.EquippingV, squadID, tableName, unit.GetID(), tx)

	/* обновляем оружие и патроны */
	for _, slot := range body.Weapons {
		if slot.Weapon == nil {
			_, err := tx.Exec("DELETE FROM "+tableName+" WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3 AND type_slot = $4",
				unit.GetID(), slot.Number, squadID, slot.Type)
			if err != nil {
				log.Fatal("delete unit body weapon slot " + err.Error())
			}
		}

		if slot.InsertToDB && slot.Weapon != nil {
			_, err := tx.Exec("INSERT INTO " + tableName + " (id_squad, type, id_squad_unit, id_equipping, slot_in_body, type_slot, quantity ) "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7)",
				squadID, "weapon", unit.GetID(), slot.Weapon.ID, slot.Number, slot.Type, 1)
			if err != nil {
				log.Fatal("insert unit body weapon slot " + err.Error())
			}
		}

		if !slot.InsertToDB && slot.Weapon != nil {
			_, err := tx.Exec("UPDATE " + tableName+
				" SET id_equipping = $2, quantity = $7 "+
				" WHERE id_squad = $3 AND id_squad_unit = $4 AND slot_in_body = $5 AND type_slot = $6 AND type = $1",
				"weapon", slot.Weapon.ID, squadID, unit.GetID(), slot.Number, slot.Type, 1)
			if err != nil {
				log.Fatal("update unit body weapon slot " + err.Error())
			}
		}

		_, err := tx.Exec("DELETE FROM " + tableName + " "+
			"WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3 AND type = $4 AND type_slot = $5",
			unit.GetID(), slot.Number, squadID, "ammo", slot.Type)
		if err != nil {
			log.Fatal("delete ammo" + err.Error())
		}

		if slot.Ammo != nil {
			_, err = tx.Exec("INSERT INTO " + tableName +
				" (id_squad, type, id_squad_unit, id_equipping, slot_in_body, type_slot, quantity)"+
				" VALUES ($1, $2, $3, $4, $5, $6, $7)",
				squadID, "ammo", unit.GetID(), slot.Ammo.ID, slot.Number, slot.Type, slot.AmmoQuantity)
			if err != nil {
				log.Fatal("insert ammo " + err.Error())
			}
		}

		slot.InsertToDB = false
	}
}

func updateEquipping(Equipping map[int]*detail.BodyEquipSlot, squadID int, tableName string, unitID int, tx *sql.Tx) {
	for _, slot := range Equipping {
		if slot.Equip == nil {
			_, err := tx.Exec("DELETE FROM "+tableName+" WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3 AND type_slot = $4",
				unitID, slot.Number, squadID, slot.Type)
			if err != nil {
				log.Fatal("delete unit body equip slot " + err.Error())
			}
		}

		if slot.InsertToDB && slot.Equip != nil {
			_, err := tx.Exec("INSERT INTO " + tableName + " (id_squad, type, id_squad_unit, id_equipping, slot_in_body, type_slot, quantity ) "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7)",
				squadID, "equip", unitID, slot.Equip.ID, slot.Number, slot.Type, 1)
			if err != nil {
				log.Fatal("insert unit body equip slot " + err.Error())
			}
		}

		if !slot.InsertToDB && slot.Equip != nil {
			_, err := tx.Exec("UPDATE " + tableName+
				" SET type = $1, id_equipping = $3, quantity = $7 "+
				" WHERE id_squad = $4 AND id_squad_unit = $5 AND slot_in_body = $6 AND type_slot = $2",
				"equip", slot.Equip.TypeSlot, slot.Equip.ID, squadID, unitID, slot.Number, 1)
			if err != nil {
				log.Fatal("update unit body equip slot " + err.Error())
			}
		}
		slot.InsertToDB = false
	}
}
