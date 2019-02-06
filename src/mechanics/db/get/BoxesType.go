package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/box"
	"log"
)

func Boxes() map[int]box.Box {
	allTypes := make(map[int]box.Box)

	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" name," +
		" type," +
		" capacity_size," +
		" fold_size," +
		" protect," +
		" protect_lvl," +
		" underground" +
		" " +
		"FROM box_type ")
	if err != nil {
		log.Fatal("get all type box " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		typeBox := box.Box{}
		err := rows.Scan(&typeBox.TypeID, &typeBox.Name, &typeBox.Type, &typeBox.CapacitySize, &typeBox.FoldSize,
			&typeBox.Protect, &typeBox.ProtectLvl, &typeBox.Underground)
		if err != nil {
			log.Fatal("get scan all type box " + err.Error())
		}

		allTypes[typeBox.TypeID] = typeBox
	}

	return allTypes
}
