package Squad

import "log"

func AddNewSquad(name string, userID int) (err error, squad *Squad) {
	// TODO проверка на имя
	res, err := db.Exec("INSERT INTO squads (name, id_user) VALUES ($1, $2)", name, userID)
	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	squad = &Squad{ID: int(id), Name:name}

	return nil, squad
}
