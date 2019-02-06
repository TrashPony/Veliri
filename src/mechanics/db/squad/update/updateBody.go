package update

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"log"
)

func UpdateBody(unit *unit.Unit, squadID int, tx *sql.Tx) {
	body := unit.GetBody()

	/* обновляем оборудование */
	updateEquipping(body.EquippingI, squadID, unit.GetID(), tx)
	updateEquipping(body.EquippingII, squadID, unit.GetID(), tx)
	updateEquipping(body.EquippingIII, squadID, unit.GetID(), tx)
	updateEquipping(body.EquippingIV, squadID, unit.GetID(), tx)
	updateEquipping(body.EquippingV, squadID, unit.GetID(), tx)

	/* обновляем оружие и патроны */
	for _, slot := range body.Weapons {
		if slot.Weapon == nil {
			_, err := tx.Exec("DELETE FROM squad_units_equipping WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3 AND type_slot = $4",
				unit.GetID(), slot.Number, squadID, slot.Type)
			if err != nil {
				log.Fatal("delete unit body weapon slot " + err.Error())
			}
		}

		if slot.InsertToDB && slot.Weapon != nil {
			_, err := tx.Exec("INSERT INTO squad_units_equipping ("+
				"id_squad, "+
				"type, "+
				"id_squad_unit, "+
				"id_equipping, "+
				"slot_in_body, "+
				"type_slot, "+
				"quantity, "+
				"used, "+
				"steps_for_reload, "+
				"hp, "+
				"target ) "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
				squadID,
				"weapon",
				unit.GetID(),
				slot.Weapon.ID,
				slot.Number,
				slot.Type,
				1,
				false,
				0,
				slot.HP,
				"",
			)
			if err != nil {
				log.Fatal("insert unit body weapon slot " + err.Error())
			}
		}

		if !slot.InsertToDB && slot.Weapon != nil {
			_, err := tx.Exec("UPDATE squad_units_equipping "+
				" SET id_equipping = $2, quantity = $7, hp = $8 "+
				" WHERE id_squad = $3 AND id_squad_unit = $4 AND slot_in_body = $5 AND type_slot = $6 AND type = $1",
				"weapon",
				slot.Weapon.ID,
				squadID,
				unit.GetID(),
				slot.Number,
				slot.Type,
				1,
				slot.HP,
			)
			if err != nil {
				log.Fatal("update unit body weapon slot " + err.Error())
			}
		}

		_, err := tx.Exec("DELETE FROM squad_units_equipping "+
			"WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3 AND type = $4 AND type_slot = $5",
			unit.GetID(), slot.Number, squadID, "ammo", slot.Type)
		if err != nil {
			log.Fatal("delete ammo" + err.Error())
		}

		if slot.Ammo != nil {
			_, err = tx.Exec("INSERT INTO squad_units_equipping "+
				" (id_squad, type, id_squad_unit, id_equipping, slot_in_body, type_slot, quantity, used, steps_for_reload, hp, target)"+
				" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
				squadID, "ammo", unit.GetID(), slot.Ammo.ID, slot.Number, slot.Type, slot.AmmoQuantity, false, 0, 1, "")
			if err != nil {
				log.Fatal("insert ammo " + err.Error())
			}
		}
		slot.InsertToDB = false
	}
}

func updateEquipping(Equipping map[int]*detail.BodyEquipSlot, squadID int, unitID int, tx *sql.Tx) {
	for _, slot := range Equipping {
		if slot.Equip == nil {
			_, err := tx.Exec("DELETE FROM squad_units_equipping WHERE id_squad_unit=$1 AND slot_in_body = $2 AND id_squad = $3 AND type_slot = $4",
				unitID, slot.Number, squadID, slot.Type)
			if err != nil {
				log.Fatal("delete unit body equip slot " + err.Error())
			}
		}

		if slot.InsertToDB && slot.Equip != nil {
			_, err := tx.Exec("INSERT INTO squad_units_equipping ("+
				"id_squad, "+
				"type, "+
				"id_squad_unit, "+
				"id_equipping, "+
				"slot_in_body, "+
				"type_slot, "+
				"quantity, "+
				"used, "+
				"steps_for_reload, "+
				"hp, "+
				"target ) "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
				squadID,
				"equip",
				unitID,
				slot.Equip.ID,
				slot.Number,
				slot.Type,
				1, // т.к. к 1 слот эквипа можно положить только 1 итем, возможно потом измениться
				slot.Used,
				slot.StepsForReload,
				slot.HP,
				parseTarget(slot.Target),
			)
			if err != nil {
				log.Fatal("insert unit body equip slot " + err.Error())
			}
		}

		if !slot.InsertToDB && slot.Equip != nil {
			_, err := tx.Exec("UPDATE squad_units_equipping "+
				" SET type = $1, id_equipping = $3, quantity = $7, used = $8, steps_for_reload = $9, hp = $10, target = $11 "+
				" WHERE id_squad = $4 AND id_squad_unit = $5 AND slot_in_body = $6 AND type_slot = $2",
				"equip",
				slot.Equip.TypeSlot,
				slot.Equip.ID,
				squadID,
				unitID,
				slot.Number,
				1, // т.к. к 1 слот эквипа можно положить только 1 итем, возможно потом измениться
				slot.Used,
				slot.StepsForReload,
				slot.HP,
				parseTarget(slot.Target),
			)
			if err != nil {
				log.Fatal("update unit body equip slot " + err.Error())
			}
		}
		slot.InsertToDB = false
	}
}
