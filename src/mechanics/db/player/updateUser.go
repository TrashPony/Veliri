package player

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"log"
)

func UpdateUser(user *player.Player) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE users "+
		"SET mail = $2, credits = $3, training = $4, last_base_id = $5, fraction = $6, avatar = $7, biography = $8, "+
		" scientific_points = $9, attack_points = $10, production_points = $11, title = $12 "+
		"WHERE id = $1",
		user.GetID(), user.GetEmail(), user.GetCredits(), user.Training, user.LastBaseID, user.Fraction,
		user.AvatarIcon, user.Biography, user.ScientificPoints, user.AttackPoints, user.ProductionPoints, user.Title)
	if err != nil {
		log.Fatal("update user " + err.Error())
	}

	UpdateUserSkills(user, tx)
	tx.Commit()
}

func UpdateUserSkills(user *player.Player, tx *sql.Tx) {
	_, err := tx.Exec("DELETE FROM user_skills WHERE id_user = $1", user.GetID())
	if err != nil {
		log.Fatal("delete all skills" + err.Error())
	}

	for _, currentSkill := range user.CurrentSkills {
		_, err := tx.Exec("INSERT INTO user_skills (lvl, id_skill, id_user) VALUES ($1, $2, $3)",
			currentSkill.Level, currentSkill.ID, user.GetID())
		if err != nil {
			log.Fatal("add new skills to user" + err.Error())
		}
	}
}
