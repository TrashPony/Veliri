package updateSquad

import (
	"../../gameObjects/squad"
	"../../../dbConnect"
	"log"
)

func InventorySquad(squad *squad.Squad) {
	for slotNum, slot := range squad.Inventory {
		if slot.Item == nil {
			_, err := dbConnect.GetDBConnect().Exec("DELETE FROM squad_inventory WHERE id_squad=$1 AND slot = $2",
				squad.ID, slotNum)
			if err != nil {
				log.Fatal("delete item to inventory" + err.Error())
			}
		}

		if slot.InsertToDB && slot.Item != nil {
			_, err := dbConnect.GetDBConnect().Exec("INSERT INTO squad_inventory (id_squad, slot, item_type, item_id, quantity) VALUES ($1, $2, $3, $4, $5)",
				squad.ID, slotNum, slot.Type, slot.ItemID, slot.Quantity)
			if err != nil {
				log.Fatal("add new item to inventory" + err.Error())
			}
		}

		if !slot.InsertToDB && slot.Item != nil {
			_, err := dbConnect.GetDBConnect().Exec("UPDATE squad_inventory SET quantity = $1, item_type = $2, item_id = $3 WHERE id_squad = $4 AND slot = $5",
				slot.Quantity, slot.Type, slot.ItemID, squad.ID, slotNum)
			if err != nil {
				log.Fatal("update inventory item squad" + err.Error())
			}
		}
	}
}
