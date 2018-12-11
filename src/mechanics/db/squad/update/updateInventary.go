package update

import (
	"../../../gameObjects/squad"
	"database/sql"
	"log"
)

func InventorySquad(squad *squad.Squad, tx *sql.Tx) {
	for slotNum, slot := range squad.Inventory.Slots {
		if slot.Item == nil {
			_, err := tx.Exec("DELETE FROM squad_inventory WHERE id_squad=$1 AND slot = $2",
				squad.ID, slotNum)
			if err != nil {
				log.Fatal("delete item to inventory" + err.Error())
			}
			delete(squad.Inventory.Slots, slotNum)
		}

		if slot.InsertToDB && slot.Item != nil {
			_, err := tx.Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity, hp) VALUES ($1, $2, $3, $4, $5, $6)",
				squad.ID, slotNum, slot.Type, slot.ItemID, slot.Quantity, slot.HP)
			if err != nil {
				log.Fatal("add new item to inventory" + err.Error())
			}
			slot.InsertToDB = false
		}

		if !slot.InsertToDB && slot.Item != nil {
			_, err := tx.Exec("UPDATE squad_inventory SET quantity = $1, item_type = $2, item_id = $3, hp = $6 WHERE id_squad = $4 AND slot = $5",
				slot.Quantity, slot.Type, slot.ItemID, squad.ID, slotNum, slot.HP)
			if err != nil {
				log.Fatal("update inventory item squad" + err.Error())
			}
		}
	}
}
