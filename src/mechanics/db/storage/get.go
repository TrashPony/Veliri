package storage

import (
	"../../../dbConnect"
	inv "../../gameObjects/inventory"
	"../inventory"
	"log"
)

func UserStorage(userId, baseId int) *inv.Inventory {
	var userInventory inv.Inventory
	userInventory.Slots = make(map[int]*inv.Slot)

	rows, err := dbConnect.GetDBConnect().Query("SELECT slot, item_type, item_id, quantity, hp "+
		"FROM base_storage "+
		"WHERE base_id = $1 AND user_id = $2", baseId, userId)
	if err != nil {
		log.Fatal("get storage inventory " + err.Error())
	}
	defer rows.Close()

	userInventory.Slots = make(map[int]*inv.Slot)

	inventory.FillInventory(&userInventory, rows)

	return &userInventory
}
