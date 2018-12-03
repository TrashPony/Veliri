package mapEditor

import (
	"../../../dbConnect"
	"log"
)

func ReplaceType(oldID, newID int) {
	//todo если новый тип имеет отличный радиус надо его подставлять + смотреть это на всех картах
	_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET id_type = $1 WHERE id_type = $2",
		newID, oldID)
	if err != nil {
		log.Fatal("offset r coordinates map " + err.Error())
	}

	DeleteType(oldID)
}
