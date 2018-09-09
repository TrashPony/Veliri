package insert

import (
	"../../../dbConnect"
	"../../gameObjects/squad"
	"log"
)

func AddNewSquad(name string, userID int) (err error, newSquad *squad.Squad) {
	// TODO проверка на имя, сделать другие отряды не активными
	id := 0
	err = dbConnect.GetDBConnect().QueryRow("INSERT INTO squads (name, active, id_user, in_game) VALUES ($1, $2, $3, $4) RETURNING id", name, true, userID, false).Scan(&id)

	if err != nil {
		log.Fatal("add new squad " + err.Error())
		return err, nil
	}

	newSquad = &squad.Squad{ID: id, Active: true, Name: name, InGame: false}

	return nil, newSquad
}
