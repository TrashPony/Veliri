package chats

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/chats"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"sync"
)

var Groups = NewChatGroupStore()

type groups struct {
	groups map[int]*chatGroup.Group
	mx     sync.RWMutex
}

func NewChatGroupStore() *groups {
	return &groups{
		groups: chats.Chats(),
	}
}

func (g *groups) GetAllGroups() map[int]*chatGroup.Group {
	return g.groups
}

func (g *groups) GetGroup(groupID int) *chatGroup.Group {
	group := g.groups[groupID]
	return group
}

func (g *groups) GetAllUserGroups(user *player.Player) map[int]*chatGroup.Group {
	userGroups := make(map[int]*chatGroup.Group)

	for _, group := range g.groups {
		if g.CheckUserSubscribe(group.ID, user) {
			userGroups[group.ID] = group
		}
	}

	return userGroups
}

func (g *groups) CreateNewGroup() {

}

func (g *groups) SubscribeGroup(groupID int, user *player.Player, password string) bool {

	if g.CheckUserSubscribe(groupID, user) {
		return false
	}

	for _, group := range g.groups {
		if group.ID == groupID && (group.Public || group.Password == password) {
			group.Users[user.GetID()] = true // при подключение игрок онайн
			// TODO обновление в бд
		}
	}

	return true
}

func (g *groups) CheckUserSubscribe(groupID int, user *player.Player) bool {

	group, find := g.groups[groupID]
	if !find {
		return false
	}

	for id := range group.Users {
		if user.GetID() == id {
			return true
		}
	}

	return false
}
