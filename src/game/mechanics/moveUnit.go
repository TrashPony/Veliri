package mechanics

import (
	"../objects"
	"database/sql"
)

func MoveUnit(idGame int, unit *objects.Unit, toX int, toY int ) (int, int, error) {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		return 0, 0, err
	}
	// устанавливает фраг готовности пользователя в тру
	_ , err = db.Query("UPDATE action_game_unit  SET x = $1, y = $2 WHERE id=$3 AND id_game=$4", toX, toY, unit.Id, idGame)
	if err != nil {
		return 0, 0, err
	} else {
		return toX, toY, nil
	}
}