package get

import (
	"../../../dbConnect"
	inv "../../gameObjects/inventory"
	"log"
)

func UserStorage(userId int) *inv.Inventory {
	// если юзер сидит в базе то он полюбому будет в этой таблице, 1 юзхер может быть одноверменно в 1 базе
	// если в этой таблице один пользователь указан много раз то это ошибка
	rows, err := dbConnect.GetDBConnect().Query("SELECT base_id FROM base_users WHERE user_id = $1 LIMIT 1 ", userId)
	if err != nil {
		log.Fatal("get base storage " + err.Error())
	}
	defer rows.Close()

	baseID := 0

	for rows.Next() {
		rows.Scan(&baseID)
	}

	if baseID != 0 {
		var inventory inv.Inventory
		inventory.Slots = make(map[int]*inv.Slot)

		rows, err := dbConnect.GetDBConnect().Query("SELECT slot, item_type, item_id, quantity, hp "+
			"FROM base_storage "+
			"WHERE base_id = $1 AND user_id = $2", baseID, userId)
		if err != nil {
			log.Fatal("get storage inventory " + err.Error())
		}
		defer rows.Close()

		inventory.Slots = make(map[int]*inv.Slot)

		FillInventory(&inventory, rows)

		return &inventory
	} else {
		return nil
	}
}
