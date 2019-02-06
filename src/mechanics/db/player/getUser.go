package player

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"log"
)

func User(id int, login string) *player.Player {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id, name, mail, credits, experience_point "+
		"FROM users "+
		"WHERE id=$1 AND name=$2", id, login)
	if err != nil {
		log.Fatal("get user " + err.Error())
	}
	defer rows.Close()

	newUser := player.Player{}

	for rows.Next() {

		var id, credits, experiencePoint int
		var name, mail string

		err := rows.Scan(&id, &name, &mail, &credits, &experiencePoint)
		if err != nil {
			log.Fatal("get user " + err.Error())
		}

		newUser.SetID(id)
		newUser.SetLogin(name)
		newUser.SetEmail(mail)
		newUser.SetCredits(credits)
		newUser.SetExperiencePoint(experiencePoint)

		getUserSkills(&newUser)
		getUserBase(&newUser)
	}

	return &newUser
}

func getUserBase(user *player.Player) {
	rows, err := dbConnect.GetDBConnect().Query("SELECT base_id "+
		"FROM base_users "+
		"WHERE user_id=$1", user.GetID())
	if err != nil {
		log.Fatal("get base user " + err.Error())
	}
	defer rows.Close()

	id := 0

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			println("get base user " + err.Error())
		}
		user.InBaseID = id
	}
}

func getUserSkills(user *player.Player) {

}
