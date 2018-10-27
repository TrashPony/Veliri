package get

import (
	"../../../dbConnect"
	"../../player"
	"log"
)

func User(id int, login string) *player.Player {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id, name, mail, credits FROM users WHERE id=$1 AND name=$2", id, login)
	if err != nil {
		log.Fatal("get user " + err.Error())
	}
	defer rows.Close()

	newUser := player.Player{}

	for rows.Next() {

		var id, credits int
		var name, mail string

		err := rows.Scan(&id, &name, &mail, &credits)
		if err != nil {
			log.Fatal("get user " + err.Error())
		}

		newUser.SetID(id)
		newUser.SetLogin(name)
		newUser.SetEmail(mail)
		newUser.SetCredits(credits)
	}

	return &newUser
}
