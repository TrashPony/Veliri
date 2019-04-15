package chat

import (
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"sync"
)

var Groups = NewChatGroupStore()

type groups struct {
	groups map[int]*Group
	mx     sync.RWMutex
}

func NewChatGroupStore() *groups {
	return &groups{
		groups: make(map[int]*Group),
	}
}

type Group struct {
	ID       int                    `json:"id"`
	Name     string                 `json:"name"`
	Public   bool                   `json:"public"` /* публичный чат любой может войти*/
	Users    map[int]*player.Player `json:"users"`  /* [id] user */
	Password string                 `json:"password"`
	History  []string               `json:"history"`
}

func (g *groups) GetGroup(groupID int) *Group {
	group := g.groups[groupID]
	return group
}

func (g *groups) GetAllUserGroups(user *player.Player) map[int]*Group {
	userGroups := make(map[int]*Group)

	for _, group := range g.groups {
		if g.CheckUserSubscribe(group.ID, user) {
			userGroups[group.ID] = group
		}
	}

	return userGroups
}

func (g *groups) CreateNewGroup() {

}

func (g *groups) SubscribeGroup() {

}

func (g *groups) CheckUserSubscribe(groupID int, user *player.Player) bool {

	group, find := g.groups[groupID]
	if !find {
		return false
	}

	for _, subscribeUser := range group.Users {
		if user.GetID() == subscribeUser.GetID() {
			return true
		}
	}

	return false
}
