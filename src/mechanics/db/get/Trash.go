package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/trashItem"
	"log"
)

func TrashItems() map[int]trashItem.TrashItem {
	allTypes := make(map[int]trashItem.TrashItem)

	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" name," +
		" size," +
		" description " +
		" " +
		"FROM trash_type ")
	if err != nil {
		log.Fatal("get all trash " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		typeTrash := trashItem.TrashItem{}
		err := rows.Scan(&typeTrash.ID, &typeTrash.Name, &typeTrash.Size, &typeTrash.Description)
		if err != nil {
			log.Fatal("scan all trash" + err.Error())
		}

		allTypes[typeTrash.ID] = typeTrash
	}

	return allTypes
}
