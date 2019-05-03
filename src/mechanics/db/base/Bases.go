package base

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
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

		err := rows.Scan(&gameBase.ID, &gameBase.Name, &gameBase.Q, &gameBase.R, &gameBase.MapID, &transports,
			&defenders, &gameBase.GravityRadius)
		if err != nil {
			log.Fatal("get scan all base " + err.Error())
		}

		respawnBases(&gameBase)
		gameBase.CreateTransports(transports)
		gameBase.CreateDefenders(defenders)

		bases[gameBase.ID] = &gameBase
	}

	return bases
}

func respawnBases(gameBase *base.Base) {
	gameBase.Respawns = make([]*coordinate.Coordinate, 0)

	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT "+
		" q,"+
		" r"+
		" "+
		" FROM bases_respawns WHERE base_id = $1", gameBase.ID)
	if err != nil {
		log.Fatal("get all respawn in base " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var respawn coordinate.Coordinate

		err := rows.Scan(&respawn.Q, &respawn.R)
		if err != nil {
			log.Fatal("scan all respawn in base " + err.Error())
		}

		gameBase.Respawns = append(gameBase.Respawns, &respawn)
	}
}
