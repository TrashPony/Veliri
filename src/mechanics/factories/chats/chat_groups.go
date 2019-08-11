package chats

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/chats"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Groups = NewChatGroupStore()

// todo вероятный проблемы конкуретного доступа, ибо я забыл везде использовать мьютекс...
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

func (g *groups) GetAllowUserGroups(user *player.Player) map[int]*chatGroup.Group {
	userGroups := make(map[int]*chatGroup.Group)

	for _, group := range g.groups {
		if group.Public && (group.Fraction == "" || group.Fraction == user.Fraction) {
			userGroups[group.ID] = group
		}
	}

	return userGroups
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

func (g *groups) CreateNewGroup(user *player.Player, name, password, avatar, greetings string) (*chatGroup.Group, error) {
	for _, group := range g.groups {
		if group.Name == name {
			return nil, errors.New("name is busy")
		}
	}

	newGroup := &chatGroup.Group{
		Name:         name,
		Greetings:    greetings,
		UserCreate:   true,
		UserIdCreate: user.GetID(),
		Local:        false,
		Public:       true,
		Users:        make(map[int]bool),
		History:      make([]*chatGroup.Message, 0),
	}

	newGroup.SetAvatar(avatar)
	newGroup.SetPassword(password)
	newGroup.ID = chats.AddNewGroup(newGroup)
	g.groups[newGroup.ID] = newGroup
	g.SubscribeGroup(newGroup.ID, user, password)

	return nil, nil
}

func (g *groups) CreateNewPrivateGroup(user1 *player.Player, user2 *player.Player) (*chatGroup.Group, *chatGroup.Message) {
	g.mx.Lock()
	defer g.mx.Unlock()

	// если у игроков уже был приватный чат то отправляем его
	for _, privateGroup := range g.groups {
		if privateGroup.Private &&
			(privateGroup.PrivateKey == strconv.Itoa(user1.GetID())+":"+strconv.Itoa(user2.GetID()) ||
				privateGroup.PrivateKey == strconv.Itoa(user2.GetID())+":"+strconv.Itoa(user1.GetID())) {

			systemMessage := &chatGroup.Message{
				//Message: "игрок " + client.GetLogin() + " покинул этот чат.",
				Time:   time.Now().UTC(),
				System: true,
			}

			if g.CheckUserSubscribe(privateGroup.ID, user1) {
				//user2 не подписан
				chats.AddUserInChat(privateGroup.ID, user2.GetID())
				privateGroup.Users[user2.GetID()] = true
				systemMessage.Message = "игрок " + user2.GetLogin() + " вернулся в чат."
			}

			if g.CheckUserSubscribe(privateGroup.ID, user2) {
				//user1 не подписан
				chats.AddUserInChat(privateGroup.ID, user1.GetID())
				privateGroup.Users[user1.GetID()] = true
				systemMessage.Message = "игрок " + user1.GetLogin() + " вернулся в чат."
			}

			privateGroup.History = append(privateGroup.History, systemMessage)
			return privateGroup, systemMessage
		}
	}

	newPrivateGroup := &chatGroup.Group{
		Name:       user1.GetLogin() + "-" + user2.GetLogin(),
		Local:      false,
		Public:     false,
		Private:    true,
		Users:      make(map[int]bool),
		History:    make([]*chatGroup.Message, 0),
		PrivateKey: strconv.Itoa(user1.GetID()) + ":" + strconv.Itoa(user2.GetID()),
	}

	newPrivateGroup.ID = chats.AddNewGroup(newPrivateGroup)
	g.groups[newPrivateGroup.ID] = newPrivateGroup

	// что бы в приватную группу нельзя было подключится нельзя использовать общий метод SubscribeGroup
	chats.AddUserInChat(newPrivateGroup.ID, user1.GetID())
	chats.AddUserInChat(newPrivateGroup.ID, user2.GetID())
	newPrivateGroup.Users[user1.GetID()] = true
	newPrivateGroup.Users[user2.GetID()] = true

	return newPrivateGroup, nil
}

func (g *groups) SubscribeGroup(groupID int, user *player.Player, password string) (bool, error) {

	if g.CheckUserSubscribe(groupID, user) {
		return false, errors.New("already")
	}

	for _, group := range g.groups {
		if group.ID == groupID {

			if group.Fraction != "" && group.Fraction != user.Fraction {
				return false, errors.New("wrong fraction")
			}

			// если пароль есть и он не верный то ошибка
			if group.Secure && group.GetPassword() != password {
				return false, errors.New("wrong password")
			}

			// через этот метод нельзя подписатся на приватные группы
			if group.Private {
				return false, errors.New("private group")
			}

			group.Users[user.GetID()] = true // при подключение игрок онайн
			chats.AddUserInChat(group.ID, user.GetID())
			return true, nil
		}
	}

	return false, errors.New("group not find")
}

func (g *groups) Unsubscribe(groupID int, user *player.Player) {
	for _, group := range g.groups {
		// если это общая группа то просто отписываемся
		if group.ID == groupID {
			delete(group.Users, user.GetID())
			chats.RemoveUserInChat(group.ID, user.GetID())
		}

		// если это приватная группа и из нее все вышли, то удаляем ее
		if group.ID == groupID && group.Private && len(group.Users) == 0 {
			g.RemoveGroup(group.ID)
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

func (g *groups) RemoveGroup(groupID int) {
	chats.RemoveGroup(groupID)
	delete(g.groups, groupID)
}
