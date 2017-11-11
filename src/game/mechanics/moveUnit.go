package mechanics

import (
	"../objects"
)

func MoveUnit(idGame int, unit *objects.Unit, toX int, toY int ) (int, int, error) {

	rows, err := db.Query("Select  MAX(queue_attack) FROM action_game_unit WHERE id_game=$1", idGame)
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()

	var queue int

	for rows.Next() {
		err := rows.Scan(&queue)
		if err != nil {
			return 0, 0, err
		}
	}

	// устанавливает фраг готовности пользователя и ставить очередь
	_ , err = db.Query("UPDATE action_game_unit  SET x = $1, y = $2, action = $5, queue_attack = $6  WHERE id=$3 AND id_game=$4", toX, toY, unit.Id, idGame, false, queue + 1)
	if err != nil {
		return 0, 0, err
	} else {
		return toX, toY, nil
	}
}