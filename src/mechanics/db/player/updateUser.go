package player

import (
	"database/sql"
	"encoding/json"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"log"
)

func UpdateUser(user *player.Player) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE users "+
		"SET mail = $2, credits = $3, training = $4, last_base_id = $5, fraction = $6, avatar = $7, biography = $8, "+
		" scientific_points = $9, attack_points = $10, production_points = $11, title = $12, story_episode = $13 "+
		"WHERE id = $1",
		user.GetID(), user.GetEmail(), user.GetCredits(), user.Training, user.LastBaseID, user.Fraction,
		user.GetAvatar(), user.Biography, user.ScientificPoints, user.AttackPoints, user.ProductionPoints, user.Title,
		user.StoryEpisode)
	if err != nil {
		log.Fatal("update user " + err.Error())
	}

	UpdateNotify(user, tx)
	UpdateUserSkills(user, tx)
	UpdateUserMission(user, tx)
	UpdateUI(user, tx)

	tx.Commit()
}

func UpdateNotify(user *player.Player, tx *sql.Tx) {
	_, err := tx.Exec("DELETE FROM user_notify WHERE id_user = $1", user.GetID())
	if err != nil {
		log.Fatal("delete ui" + err.Error())
	}

	jsonString, err := json.Marshal(user.NotifyQueue)
	_, err = tx.Exec("INSERT INTO user_notify (data, id_user) VALUES ($1, $2)",
		jsonString, user.GetID())
	if err != nil {
		log.Fatal("add new ui" + err.Error())
	}
}

func UpdateUI(user *player.Player, tx *sql.Tx) {
	_, err := tx.Exec("DELETE FROM user_interface WHERE id_user = $1", user.GetID())
	if err != nil {
		log.Fatal("delete ui" + err.Error())
	}

	jsonString, err := json.Marshal(user.UserInterface)
	_, err = tx.Exec("INSERT INTO user_interface (data, id_user) VALUES ($1, $2)",
		jsonString, user.GetID())
	if err != nil {
		log.Fatal("add new ui" + err.Error())
	}
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

func UpdateUserMission(user *player.Player, tx *sql.Tx) {
	_, err := tx.Exec("DELETE FROM user_current_mission WHERE id_user = $1", user.GetID())
	if err != nil {
		log.Fatal("delete all mission" + err.Error())
	}

	for _, mission := range user.Missions {

		jsonMission, _ := json.Marshal(mission)

		_, err := tx.Exec("INSERT INTO user_current_mission (id_user, id_mission, data) VALUES ($1, $2, $3)",
			user.GetID(), mission.ID, string(jsonMission))
		if err != nil {
			log.Fatal("add new missions to user" + err.Error())
		}
	}
}
