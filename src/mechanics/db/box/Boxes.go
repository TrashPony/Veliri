package box

import (
	"../../../dbConnect"
	"../../gameObjects/box"
	inv "../../gameObjects/inventory"
	"log"
)

func Boxes() map[int]*box.Box {
	boxes := make(map[int]*box.Box)

	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" password," +
		" destroy_time," +
		" id_map," +
		" id_box_type," +
		" q," +
		" r," +
		" rotate" +
		" " +
		" FROM box_in_map")
	if err != nil {
		log.Fatal("get all box " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var gameBox box.Box
		var password int

		err := rows.Scan(&gameBox.ID, &password, &gameBox.DestroyTime, &gameBox.MapID, &gameBox.TypeID,
			&gameBox.Q, &gameBox.R, &gameBox.Rotate)
		if err != nil {
			log.Fatal("get scan all box " + err.Error())
		}

		getTypeBox(&gameBox)
		getBoxStorage(&gameBox)
		gameBox.SetPassword(password)

		boxes[gameBox.ID] = &gameBox
	}

	return boxes
}

func getTypeBox(gameBox *box.Box) {
	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		" type,"+
		" capacity_size,"+
		" fold_size,"+
		" protect,"+
		" protect_lvl"+
		" "+
		"FROM box_type "+
		"WHERE id = $1", gameBox.TypeID)
	if err != nil {
		log.Fatal("get type box " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&gameBox.Type, &gameBox.CapacitySize, &gameBox.FoldSize, &gameBox.Protect, &gameBox.ProtectLvl)
		if err != nil {
			log.Fatal("get scan type box " + err.Error())
		}
	}
}

func getBoxStorage(gameBox *box.Box) {
	gameBox.GetStorage().Slots = make(map[int]*inv.Slot)

	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		" slot,"+
		" item_type,"+
		" item_id,"+
		" quantity,"+
		" hp"+
		" "+
		"FROM box_storage "+
		"WHERE id_box = $1", gameBox.ID)
	if err != nil {
		log.Fatal("get box inventory " + err.Error())
	}
	defer rows.Close()

	gameBox.GetStorage().FillInventory(rows)
}
