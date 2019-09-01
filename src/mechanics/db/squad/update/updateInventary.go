package update

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"log"
)

func InventorySquad(squad *squad.Squad, tx *sql.Tx) {

	if squad.MatherShip.Inventory != nil {
		updateUnitInventory(squad.MatherShip, tx)
	} else {
		createUnitInventory(squad.MatherShip)
	}

	for _, slotUnit := range squad.MatherShip.Units {
		if slotUnit.Unit != nil {

			if slotUnit.Unit.Inventory != nil {
				updateUnitInventory(slotUnit.Unit, tx)
			} else {
				createUnitInventory(slotUnit.Unit)
			}
		}
	}
}

func createUnitInventory(unit *unit.Unit) {
	newInventory := &inventory.Inventory{Slots: make(map[int]*inventory.Slot)}
	newInventory.SetSlotsSize(9999)
	unit.Inventory = newInventory
}

func updateUnitInventory(unit *unit.Unit, tx *sql.Tx) {
	for slotNum, slot := range unit.Inventory.Slots {
		if slot.Item == nil {
			_, err := tx.Exec("DELETE FROM squad_units_inventory WHERE id_unit=$1 AND slot = $2",
				unit.ID, slotNum)
			if err != nil {
				log.Fatal("delete item to inventory" + err.Error())
			}
			delete(unit.Inventory.Slots, slotNum)
		}

		if slot.InsertToDB && slot.Item != nil {
			_, err := tx.Exec("INSERT INTO squad_units_inventory (id_unit, slot, item_type, item_id, quantity, hp, place_user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
				unit.ID, slotNum, slot.Type, slot.ItemID, slot.Quantity, slot.HP, slot.PlaceUserID)
			if err != nil {
				log.Fatal("add new item to inventory" + err.Error())
			}
			slot.InsertToDB = false
		}

		if !slot.InsertToDB && slot.Item != nil {
			_, err := tx.Exec("UPDATE squad_units_inventory SET quantity = $1, item_type = $2, item_id = $3, hp = $6, place_user_id=$7 WHERE id_unit = $4 AND slot = $5",
				slot.Quantity, slot.Type, slot.ItemID, unit.ID, slotNum, slot.HP, slot.PlaceUserID)
			if err != nil {
				log.Fatal("update inventory item squad" + err.Error())
			}
		}
	}
}
