package mechanics

import (
	"../objects"
	"database/sql"
)

func SetTarget(unit objects.Unit, target string, idGame int)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {

	}

	_ , err = db.Query("UPDATE action_game_unit  SET target = $1 WHERE id=$2 AND id_game=$3", target, unit.Id, idGame)
	if err != nil {

	}
}