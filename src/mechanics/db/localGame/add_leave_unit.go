package localGame

import (
	"encoding/json"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"log"
)

func AddLeaveUnit(unit *unit.Unit, clientID, gameID int) {

	jsonUnit, err := json.Marshal(unit)

	_, err = dbConnect.GetDBConnect().Exec("INSERT INTO game_leave_unit (unit, id_user, id_game) VALUES ($1, $2, $3)",
		jsonUnit, clientID, gameID)
	if err != nil {
		println("add leave unit")
		log.Fatal(err)
	}
}

func DeleteAllLeaveUnit(gameID int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM game_leave_unit WHERE id_game=$1", gameID)
	if err != nil {
		log.Fatal("delete all leave unit" + err.Error())
	}
}
