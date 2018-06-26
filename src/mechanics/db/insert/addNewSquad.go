package insert

import (
	"../../gameObjects/squad"
	"../../../dbConnect"
	"log"
)

func AddNewSquad(name string, userID int) (err error, newSquad *squad.Squad) {
	// TODO проверка на имя
	id := 0
	err = dbConnect.GetDBConnect().QueryRow("INSERT INTO squads (name, id_user, in_game) VALUES ($1, $2, $3) RETURNING id", name, userID, false).Scan(&id)

	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	newSquad = &squad.Squad{ID: id, Name: name, InGame: false}

	return nil, newSquad
}

