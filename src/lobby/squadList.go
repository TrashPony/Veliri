package lobby

import "log"

func AddNewSquad(name string, userID int) (err error){
	_, err = db.Exec("INSERT INTO squads (name, id_user) VALUES ($1, $2)", name, userID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func GetListSquads(userID int)  (squadNames []string, err error) {

	rows, err := db.Query("Select name FROM squads WHERE id_user=$1", userID)
	if err != nil {
		log.Fatal(err)
	}

	squadNames = make([]string,0)
	var name string

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		squadNames = append(squadNames, name)
	}

	return
}