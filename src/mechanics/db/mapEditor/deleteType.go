package mapEditor

import (
	"../../../dbConnect"
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
