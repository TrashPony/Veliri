package inventory

import "../gameObjects/squad"

func AddNewSquad(name string, userID int) (err error, squad *squad.Squad) {
	// TODO проверка на имя
	/*id := 0
	err = dbConnect.GetDBConnect().QueryRow("INSERT INTO squads (name, id_user) VALUES ($1, $2) RETURNING id", name, userID).Scan(&id)

	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	squad = &Squad{ID: int(id), Name: name}
	squad.GetSquadUnits()
	squad.GetSquadMatherShip()
	squad.GetSquadEquip()*/

	return nil, nil
}

