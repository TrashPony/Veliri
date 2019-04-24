package chats

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/chats"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"strconv"
	"strings"
	"sync"
)

var Groups = NewChatGroupStore()

type groups struct {
	groups map[int]*chatGroup.Group
	// локальные группы чатов, отдельные чаты для каждой игровой зоны от них нельзя отписатся
	localGroups map[string]*chatGroup.Group // base:ID || game:ID || map:ID
	mx          sync.RWMutex
}

func NewChatGroupStore() *groups {
	return &groups{
		groups: chats.Chats(),
	}
}

func (g *groups) GetAllGroups() map[int]*chatGroup.Group {
	return g.groups
}

func (g *groups) GetAllLocalGroups() map[string]*chatGroup.Group {
	return g.localGroups
}

func (g *groups) GetGroup(groupID int) *chatGroup.Group {
	group := g.groups[groupID]
	return group
}

func (g *groups) GetLocalGroup(groupID string) *chatGroup.Group {
	// base:ID || game:ID || map:ID

	if g.localGroups == nil {
		g.localGroups = make(map[string]*chatGroup.Group)
	}

	group := g.localGroups[groupID]
	if group == nil {
		groupName := ""

		if strings.Split(groupID, ":")[0] == "base" {
			id, _ := strconv.Atoi(strings.Split(groupID, ":")[1])
			base, _ := bases.Bases.Get(id)
			groupName = base.Name
		}

		if strings.Split(groupID, ":")[0] == "game" {
			// TODO
		}

		if strings.Split(groupID, ":")[0] == "map" {
			id, _ := strconv.Atoi(strings.Split(groupID, ":")[1])
			mp, _ := maps.Maps.GetByID(id)
			groupName = mp.Name
		}

		group = &chatGroup.Group{Name: groupName, Local: true, Users: make(map[int]bool), History: make([]*chatGroup.Message, 0)}
		g.localGroups[groupID] = group
	}
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
			chats.AddUserInChat(group.ID, user.GetID())
		}
	}

	return true
}

func (g *groups) Unsubscribe(groupID int, user *player.Player) {
	for _, group := range g.groups {
		if group.ID == groupID {
			delete(group.Users, user.GetID())
			chats.RemoveUserInChat(group.ID, user.GetID())
		}
	}
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
