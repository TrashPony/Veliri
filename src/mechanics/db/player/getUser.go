package player

import (
	"encoding/json"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/skill"
	"log"
)

func User(id int, login string) *player.Player {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id, name, mail, credits, training, last_base_id,"+
		" fraction, avatar, biography, scientific_points, attack_points, production_points, title, story_episode "+
		"FROM users "+
		"WHERE id=$1 AND name=$2", id, login)
	if err != nil {
		log.Fatal("get user " + err.Error())
	}
	defer rows.Close()

	newUser := player.Player{}

	for rows.Next() {

		var id, credits int
		var name, mail, avatar string

		err := rows.Scan(&id, &name, &mail, &credits, &newUser.Training, &newUser.LastBaseID,
			&newUser.Fraction, &avatar, &newUser.Biography, &newUser.ScientificPoints, &newUser.AttackPoints,
			&newUser.ProductionPoints, &newUser.Title, &newUser.StoryEpisode)
		if err != nil {
			log.Fatal("get user " + err.Error())
		}

		newUser.NotifyQueue = make(map[string]*player.Notify, 0)

		newUser.SetAvatar(avatar)
		newUser.SetID(id)
		newUser.SetLogin(name)
		newUser.SetEmail(mail)
		newUser.SetCredits(credits)

		getUserNotify(&newUser)
		getUserSkills(&newUser)
		getUserBase(&newUser)
		getUserMission(&newUser)
		getUI(&newUser)
	}

	return &newUser
}

func getUserNotify(user *player.Player) {
	rows, err := dbConnect.GetDBConnect().Query("SELECT data "+
		"FROM user_notify "+
		"WHERE id_user=$1", user.GetID())
	if err != nil {
		log.Fatal("get notify " + err.Error())
	}
	defer rows.Close()

	var notifyJson []byte

	for rows.Next() {
		err := rows.Scan(&notifyJson)
		if err != nil {
			println("get scan notify " + err.Error())
		}

		err = json.Unmarshal(notifyJson, &user.NotifyQueue)
		if err != nil {
			println("unmarshal notify " + err.Error())
		}
	}
}

func getUI(user *player.Player) {
	rows, err := dbConnect.GetDBConnect().Query("SELECT data "+
		"FROM user_interface "+
		"WHERE id_user=$1", user.GetID())
	if err != nil {
		log.Fatal("get ui " + err.Error())
	}
	defer rows.Close()

	var uiJson []byte

	for rows.Next() {
		err := rows.Scan(&uiJson)
		if err != nil {
			println("get scan ui " + err.Error())
		}

		err = json.Unmarshal(uiJson, &user.UserInterface)
		if err != nil {
			println("unmarshal ui " + err.Error())
		}
	}
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
	user.CurrentSkills = make(map[string]*skill.Skill)

	rows, err := dbConnect.GetDBConnect().Query("SELECT lvl, id_skill FROM user_skills WHERE id_user=$1", user.GetID())
	if err != nil {
		log.Fatal("get base user " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id, lvl int
		err := rows.Scan(&lvl, &id)
		if err != nil {
			println("get skills user " + err.Error())
		}

		currentSkill := gameTypes.Skills.GetByID(id)
		currentSkill.Level = lvl

		user.CurrentSkills[currentSkill.Name] = currentSkill
	}

	// докидываем скилы нулевчого лвла
	for id, skillType := range gameTypes.Skills.GetAllSkillTypes() {
		if _, ok := user.CurrentSkills[skillType.Name]; !ok {
			currentSkill := gameTypes.Skills.GetByID(id)
			currentSkill.Level = 0 // оно и так должно быть ноль но на всякий случай
			user.CurrentSkills[currentSkill.Name] = currentSkill
		}
	}
}

func getUserMission(user *player.Player) {
	user.Missions = make(map[string]*mission.Mission)

	rows, err := dbConnect.GetDBConnect().Query("SELECT data FROM user_current_mission WHERE id_user=$1", user.GetID())
	if err != nil {
		log.Fatal("get user mission " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var userMission mission.Mission
		var data string

		err := rows.Scan(&data)
		if err != nil {
			println("scan user mission" + err.Error())
		}

		err = json.Unmarshal([]byte(data), &userMission)
		if err != nil {
			println("json unmarshal user mission" + err.Error())
		}

		user.Missions[userMission.UUID] = &userMission
	}
}
