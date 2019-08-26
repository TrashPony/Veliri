package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func DeleteType(idType int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM map_constructor WHERE id_type=$1", idType)
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbConnect.GetDBConnect().Exec("DELETE FROM coordinate_type WHERE id=$1", idType)
	if err != nil {
		log.Fatal(err)
	}
}
