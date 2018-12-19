package get

import (
	"../../../dbConnect"
	"../../gameObjects/base"
	"log"
)

func Bases() map[int]*base.Base {
	bases := make(map[int]*base.Base)

	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" base_name," +
		" q," +
		" r," +
		" id_map," +
		" resp_q," +
		" resp_r" +
		"" +
		" FROM bases")
	if err != nil {
		log.Fatal("get all bases " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var gameBase base.Base

		err := rows.Scan(&gameBase.ID, &gameBase.Name, &gameBase.Q, &gameBase.R, &gameBase.MapID, &gameBase.RespQ, &gameBase.RespR)
		if err != nil {
			log.Fatal("get scan all base " + err.Error())
		}
		bases[gameBase.ID] = &gameBase
	}

	return bases
}
