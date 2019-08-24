package storage

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	inv "github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"log"
)

func Inventory(inventory *inv.Inventory, userId, baseId int) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	if err != nil {
		log.Fatal("update storage tx error: " + err.Error())
	}

	for slotNum, slot := range inventory.Slots {
		if slot.Item == nil {
			_, err := tx.Exec("DELETE FROM base_storage WHERE base_id=$1 AND user_id=$2 AND slot = $3",
				baseId, userId, slotNum)
			if err != nil {
				log.Fatal("delete item to storage" + err.Error())
			}
			delete(inventory.Slots, slotNum)
		}

		if slot.InsertToDB && slot.Item != nil {
			_, err := tx.Exec("INSERT INTO base_storage (base_id, user_id, slot, item_type, item_id, quantity, hp, place_user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
				baseId, userId, slotNum, slot.Type, slot.ItemID, slot.Quantity, slot.HP, slot.PlaceUserID)
			if err != nil {
				log.Fatal("add new item to storage" + err.Error())
			}
			slot.InsertToDB = false
		}

		if !slot.InsertToDB && slot.Item != nil {
			_, err := tx.Exec("UPDATE base_storage SET quantity = $1, item_type = $2, item_id = $3, hp = $4, place_user_id = $8 WHERE base_id = $5 AND user_id=$6 AND slot = $7",
				slot.Quantity, slot.Type, slot.ItemID, slot.HP, baseId, userId, slotNum, slot.PlaceUserID)
			if err != nil {
				log.Fatal("update storage item" + err.Error())
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal("update storage: " + err.Error())
	}
}
