package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/skill"
	"log"
)

func Skills() map[int]skill.Skill {
	allTypes := make(map[int]skill.Skill)

	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" name," +
		" specification," +
		" experience_point," +
		" type," +
		" icon " +
		" " +
		"FROM skills ")
	if err != nil {
		log.Fatal("get all skills " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		skill := skill.Skill{}

		err := rows.Scan(&skill.ID, &skill.Name, &skill.Specification, &skill.ExperiencePoint, &skill.Type, &skill.Icon)
		if err != nil {
			log.Fatal("get scan all skills " + err.Error())
		}

		allTypes[skill.ID] = skill
	}

	return allTypes
}
