package base

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
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
		" resp_r," +
		" transport_count," +
		" defender_count," +
		" gravity_radius" +
		" " +
		" FROM bases")
	if err != nil {
		log.Fatal("get all bases " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var gameBase base.Base
		var transports int
		var defenders int

		err := rows.Scan(&gameBase.ID, &gameBase.Name, &gameBase.Q, &gameBase.R, &gameBase.MapID, &gameBase.RespQ,
			&gameBase.RespR, &transports, &defenders, &gameBase.GravityRadius)
		if err != nil {
			log.Fatal("get scan all base " + err.Error())
		}

		gameBase.CreateTransports(transports)
		gameBase.CreateDefenders(defenders)

		bases[gameBase.ID] = &gameBase
	}

	return bases
}
