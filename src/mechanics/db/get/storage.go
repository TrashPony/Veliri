package get

import (
	"../../../dbConnect"
	inv "../../gameObjects/inventory"
	"log"
)

func UserStorage(userId, baseId int) *inv.Inventory {
	var inventory inv.Inventory
	inventory.Slots = make(map[int]*inv.Slot)

	rows, err := dbConnect.GetDBConnect().Query("SELECT slot, item_type, item_id, quantity, hp "+
		"FROM base_storage "+
		"WHERE base_id = $1 AND user_id = $2", baseId, userId)
	if err != nil {
		log.Fatal("get storage inventory " + err.Error())
	}
	defer rows.Close()

	inventory.Slots = make(map[int]*inv.Slot)

	FillInventory(&inventory, rows)

	return &inventory
}
