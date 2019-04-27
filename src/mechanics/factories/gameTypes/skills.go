package gameTypes

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/skill"
	"github.com/getlantern/deepcopy"
)

type skillTypes struct {
	skills map[int]skill.Skill
}

var Skills = newSkillTypes()

func newSkillTypes() *skillTypes {
	return &skillTypes{skills: get.Skills()}
}

func (s *skillTypes) GetByID(id int) *skill.Skill {
	var newSkill skill.Skill
	skill, _ := s.skills[id]
	deepcopy.Copy(&newSkill, &skill)
	return &newSkill
}

func (s *skillTypes) GetAllSkillTypes() map[int]skill.Skill {
	return s.skills
}
