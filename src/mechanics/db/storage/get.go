package storage

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	inv "github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"log"
)

func UserStorage(userId, baseId int) *inv.Inventory {
	var userInventory inv.Inventory
	userInventory.Slots = make(map[int]*inv.Slot)
	userInventory.SetSlotsSize(999)

	rows, err := dbConnect.GetDBConnect().Query("SELECT slot, item_type, item_id, quantity, hp "+
		"FROM base_storage "+
		"WHERE base_id = $1 AND user_id = $2", baseId, userId)
	if err != nil {
		log.Fatal("get storage inventory " + err.Error())
	}
	defer rows.Close()

	userInventory.FillInventory(rows)

	return &userInventory
}
