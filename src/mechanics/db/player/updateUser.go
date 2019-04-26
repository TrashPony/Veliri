package player

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"log"
)

func UpdateUser(user *player.Player) {

	_, err := dbConnect.GetDBConnect().Exec("UPDATE users "+
		"SET mail = $2, credits = $3, training = $4, last_base_id = $5, fraction = $6, avatar = $7, biography = $8, "+
		" scientific_points = $9, attack_points = $10, production_points = $11, title = $12 "+
		"WHERE id = $1",
		user.GetID(), user.GetEmail(), user.GetCredits(), user.Training, user.LastBaseID, user.Fraction,
		user.AvatarIcon, user.Biography, user.ScientificPoints, user.AttackPoints, user.ProductionPoints, user.Title)
	if err != nil {
		log.Fatal("update user " + err.Error())
	}

	UpdateUserSkills(user)
}

func UpdateUserSkills(user *player.Player) {

}
