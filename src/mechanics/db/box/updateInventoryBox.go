package box

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"log"
)

func Inventory(updateBox *boxInMap.Box) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	if err != nil {
		log.Fatal("update box tx error: " + err.Error())
	}

	for slotNum, slot := range updateBox.GetStorage().Slots {
		if slot.Item == nil {
			_, err := tx.Exec("DELETE FROM box_storage WHERE id_box=$1 AND slot = $2",
				updateBox.ID, slotNum)
			if err != nil {
				log.Fatal("delete item from box storage" + err.Error())
			}
			delete(updateBox.GetStorage().Slots, slotNum)
		}

		if slot.InsertToDB && slot.Item != nil {
			_, err := tx.Exec("INSERT INTO box_storage (id_box, slot, item_type, item_id, quantity, hp, place_user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
				updateBox.ID, slotNum, slot.Type, slot.ItemID, slot.Quantity, slot.HP, slot.PlaceUserID)
			if err != nil {
				log.Fatal("add new item from box storage" + err.Error())
			}
			slot.InsertToDB = false
		}

		if !slot.InsertToDB && slot.Item != nil {
			_, err := tx.Exec("UPDATE box_storage SET quantity = $1, item_type = $2, item_id = $3, hp = $4, place_user_id = $7 WHERE id_box = $5 AND slot = $6",
				slot.Quantity, slot.Type, slot.ItemID, slot.HP, updateBox.ID, slotNum, slot.PlaceUserID)
			if err != nil {
				log.Fatal("update slot from box storage" + err.Error())
			}
		}
	}
	tx.Commit()
}
