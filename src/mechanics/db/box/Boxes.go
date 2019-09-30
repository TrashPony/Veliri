package box

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	inv "github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"log"
)

func Boxes() map[int]*boxInMap.Box {
	boxes := make(map[int]*boxInMap.Box)

	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" password," +
		" destroy_time," +
		" id_map," +
		" id_box_type," +
		" x," +
		" y," +
		" rotate," +
		" current_hp," +
		" owned_by_map_object" +
		" " +
		" FROM box_in_map")
	if err != nil {
		log.Fatal("get all box " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var gameBox boxInMap.Box
		var password int

		err := rows.Scan(&gameBox.ID, &password, &gameBox.DestroyTime, &gameBox.MapID, &gameBox.TypeID,
			&gameBox.X, &gameBox.Y, &gameBox.Rotate, &gameBox.HP, &gameBox.OwnedByMapObject)
		if err != nil {
			log.Fatal("get scan all box " + err.Error())
		}

		if !gameBox.OwnedByMapObject {
			getTypeBox(&gameBox)
		} else {
			gameBox.CapacitySize = 100
		}

		getBoxStorage(&gameBox)
		gameBox.SetPassword(password)

		boxes[gameBox.ID] = &gameBox
	}

	return boxes
}

func getTypeBox(gameBox *boxInMap.Box) {
	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		" type,"+
		" capacity_size,"+
		" fold_size,"+
		" protect,"+
		" protect_lvl,"+
		" underground,"+
		" hp"+
		" "+
		"FROM box_type "+
		"WHERE id = $1", gameBox.TypeID)
	if err != nil {
		log.Fatal("get type box " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&gameBox.Type, &gameBox.CapacitySize, &gameBox.FoldSize, &gameBox.Protect,
			&gameBox.ProtectLvl, &gameBox.Underground, &gameBox.HP)
		if err != nil {
			log.Fatal("get scan type box " + err.Error())
		}
	}
}

func getBoxStorage(gameBox *boxInMap.Box) {
	gameBox.GetStorage().Slots = make(map[int]*inv.Slot)
	gameBox.GetStorage().SetSlotsSize(999)

	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		" slot,"+
		" item_type,"+
		" item_id,"+
		" quantity,"+
		" hp,"+
		" place_user_id"+
		" "+
		"FROM box_storage "+
		"WHERE id_box = $1", gameBox.ID)
	if err != nil {
		log.Fatal("get box inventory " + err.Error())
	}
	defer rows.Close()

	gameBox.GetStorage().FillInventory(rows)
}
