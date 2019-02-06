package player

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"log"
)

func UpdateUser(user *player.Player) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE users "+
		"SET mail = $2, credits = $3, experience_point = $4 "+
		"WHERE id = $1",
		user.GetID(), user.GetEmail(), user.GetCredits(), user.GetExperiencePoint())
	if err != nil {
		log.Fatal("update user " + err.Error())
	}

	UpdateUserSkills(user)
}

func UpdateUserSkills(user *player.Player) {

}
