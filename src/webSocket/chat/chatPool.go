package chat

import (
	"github.com/TrashPony/Veliri/src/mechanics/chat"
	"github.com/TrashPony/Veliri/src/mechanics/factories/chats"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)

var chatPipe = make(chan chatMessage)

type chatMessage struct {
	Event       string                   `json:"event"`
	UserName    string                   `json:"user_name"`
	MessageText string                   `json:"message_text"`
	Message     *chatGroup.Message       `json:"message"`
	GroupID     int                      `json:"group_id"`
	UserID      int                      `json:"user_id"`
	Group       *chatGroup.Group         `json:"group"`
	Groups      map[int]*chatGroup.Group `json:"groups"`
	Password    string                   `json:"password"`
	Users       []*player.ShortUserInfo  `json:"users"`
	User        *player.ShortUserInfo    `json:"user"`
	Local       bool                     `json:"local"`
}

func AddNewUser(ws *websocket.Conn, login string, id int) {
	newPlayer, ok := players.Users.Get(id)
	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	chat.Clients.AddNewClient(ws, newPlayer) // Регистрируем нового Клиента
	Reader(ws)
}

func checkUserOnline(group *chatGroup.Group, id int, local bool) bool {

	chatUser := chat.Clients.GetByID(id)

	// если игрок онайн но игрок уже отключился от сети то говорим что он не онлайн и обновляем у всех список
	if group.Users[id] && chatUser == nil {
		if local {
			delete(group.Users, id)
		} else {
			group.Users[id] = false
		}
		return true
	}

	// если игрока был офлайн и стал онлайн то обновляем у всех список пользователей
	if !group.Users[id] && chatUser != nil {
		group.Users[id] = true
		return true
	}

	return false
}

func UserOnlineChecker() {
	for {
		for _, group := range chats.Groups.GetAllGroups() {
			update := false
			for id := range group.Users {
				update = checkUserOnline(group, id, false)
			}

			if update {
				SendMessage("UpdateUsers", nil, 0, group.ID, group, nil, getUsersInChatGroup(group, true), false, nil)
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func LocalChatChecker() {
	// воркет сделит что бы у игроков был открыть нужный локальный чат, и обновляет в нем игроков

	for {

		for _, localGroup := range chats.Groups.GetAllLocalGroups() {
			update := false
			for id := range localGroup.Users {
				update = checkUserOnline(localGroup, id, true)
			}

			if update {
				SendMessage("UpdateUsers", nil, 0, localGroup.ID, localGroup, nil, getUsersInChatGroup(localGroup, false), true, nil)
			}
		}

		users, mx := chat.Clients.GetAllConnects()
		
		// делаем копию карты что бы не вызвать дедлок или рантайм ошибку конкурентного чтения записи.
		fakeUsers := make(map[*websocket.Conn]*player.Player)
		for key, value := range users {
			fakeUsers[key] = value
		}

		mx.Unlock()

		for _, client := range fakeUsers {

			_, id := getLocalChat(client)

			for localID, localGroup := range chats.Groups.GetAllLocalGroups() {
				// если текущий локальный чат не совподает с прошлым, то удаляем игрока и обновляем у всех список юзеров
				if id != localID && localGroup.CheckUserInGroup(client.GetID()) {
					delete(localGroup.Users, client.GetID())
					SendMessage("UpdateUsers", nil, 0, localGroup.ID, localGroup, nil, getUsersInChatGroup(localGroup, false), true, nil)
				}

				// если текущий пользователь не находится в группе или он там офлайн притом что это его группа то обновляем статус онлайн у всех
				if id == localID && !localGroup.CheckUserInGroup(client.GetID()) {
					localGroup.Users[client.GetID()] = true
					SendMessage("UpdateUsers", nil, 0, localGroup.ID, localGroup, nil, getUsersInChatGroup(localGroup, false), true, nil)
				}
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func Reader(ws *websocket.Conn) {
	for {

		var msg chatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			chat.Clients.DelClientByWS(ws)
			break
		}

		client := chat.Clients.GetByWs(ws)
		if client != nil {
			if msg.Event == "OpenChat" {
				group, _ := getLocalChat(client)
				SendMessage(msg.Event, msg.Message, client.GetID(), 0, group, chats.Groups.GetAllUserGroups(client), nil, false, client.GetShortUserInfo(false))
			}

			if msg.Event == "GetAllGroups" {
				SendMessage(msg.Event, msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID), chats.Groups.GetAllGroups(), nil, false, nil)
			}

			if msg.Event == "ChangeGroup" {
				if chats.Groups.CheckUserSubscribe(msg.GroupID, client) {
					group := chats.Groups.GetGroup(msg.GroupID)
					SendMessage(msg.Event, msg.Message, client.GetID(), 0, group, nil, getUsersInChatGroup(group, true), false, nil)
				} else {
					group, _ := getLocalChat(client)
					SendMessage(msg.Event, msg.Message, client.GetID(), 0, group, nil, getUsersInChatGroup(group, false), false, nil)
				}

			}

			if msg.Event == "SubscribeGroup" {
				group := chats.Groups.GetGroup(msg.GroupID)
				if chats.Groups.SubscribeGroup(msg.GroupID, client, msg.Password) {
					SendMessage("OpenChat", msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID), chats.Groups.GetAllUserGroups(client), nil, false, nil)
					SendMessage("UpdateUsers", nil, 0, msg.GroupID, group, nil, getUsersInChatGroup(group, true), false, nil)
				}
			}

			if msg.Event == "Unsubscribe" {
				group := chats.Groups.GetGroup(msg.GroupID)
				if group != nil {
					chats.Groups.Unsubscribe(msg.GroupID, client)
					SendMessage("UpdateUsers", nil, 0, msg.GroupID, group, nil, getUsersInChatGroup(group, true), false, nil)
					SendMessage("OpenChat", msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID), chats.Groups.GetAllUserGroups(client), nil, false, nil)
				}
			}

			if msg.Event == "CreateNewGroup" {

			}

			if msg.Event == "NewChatMessage" {
				if chats.Groups.CheckUserSubscribe(msg.GroupID, client) {
					// добавляем в историю, отправляем не сообщение текстом а обьектом
					group := chats.Groups.GetGroup(msg.GroupID)

					chatMessage := chatGroup.Message{UserName: client.GetLogin(), AvatarIcon: client.AvatarIcon, Message: msg.MessageText, Time: time.Now().UTC()}
					group.History = append(group.History, &chatMessage)

					SendMessage(msg.Event, &chatMessage, 0, msg.GroupID, nil, nil, nil, false, nil)
				} else {
					// если msg.GroupID == 0 то это сообщение в локальный чат
					group, _ := getLocalChat(client)

					chatMessage := chatGroup.Message{UserName: client.GetLogin(), AvatarIcon: client.AvatarIcon, Message: msg.MessageText, Time: time.Now().UTC()}
					group.History = append(group.History, &chatMessage)

					SendMessage(msg.Event, &chatMessage, 0, msg.GroupID, group, nil, nil, true, nil)
				}
			}
		}
	}
}

func getLocalChat(client *player.Player) (*chatGroup.Group, string) {
	if client.InBaseID > 0 {
		// игрок на базе
		return chats.Groups.GetLocalGroup("base:" + strconv.Itoa(client.InBaseID)), "base:" + strconv.Itoa(client.InBaseID)
	}

	if client.GetSquad().InGame {
		// игрок в игре
		return chats.Groups.GetLocalGroup("game:" + strconv.Itoa(client.GetGameID())), "game:" + strconv.Itoa(client.GetGameID())
	}

	// игрок на глобальной карте
	return chats.Groups.GetLocalGroup("map:" + strconv.Itoa(client.GetSquad().MapID)), "map:" + strconv.Itoa(client.GetSquad().MapID)
}

func getUsersInChatGroup(group *chatGroup.Group, all bool) []*player.ShortUserInfo {
	users := make([]*player.ShortUserInfo, 0)

	for id := range group.Users {
		chatUser := chat.Clients.GetByID(id)

		if chatUser != nil {
			if all || group.Users[id] {
				users = append(users, chatUser.GetShortUserInfo(false))
			}
		} else {
			group.Users[id] = false
		}
	}

	return users
}

func SendMessage(event string, senderMessage *chatGroup.Message, userID, GroupID int, group *chatGroup.Group,
	groups map[int]*chatGroup.Group, users []*player.ShortUserInfo, local bool, user *player.ShortUserInfo) {

	chatPipe <- chatMessage{
		Event:   event,
		Message: senderMessage,
		GroupID: GroupID,
		UserID:  userID,
		Group:   group,
		Groups:  groups,
		Users:   users,
		Local:   local,
		User:    user,
	}
}

func CommonChatSender() {
	for {
		resp := <-chatPipe

		users, mx := chat.Clients.GetAllConnects()
		for ws, client := range users {

			if chats.Groups.CheckUserSubscribe(resp.GroupID, client) && resp.UserID == 0 { // отправляем всей подписоте
				err := ws.WriteJSON(resp)
				if err != nil {
					chat.Clients.DelClientByWS(ws)
				}
			}

			if resp.GroupID == 0 && resp.UserID == client.GetID() && !resp.Local { // отправляем только инициатору
				err := ws.WriteJSON(resp)
				if err != nil {
					chat.Clients.DelClientByWS(ws)
				}
			}

			if resp.Local && resp.Group.CheckUserInGroup(client.GetID()) { // сообщение в локальный чат
				err := ws.WriteJSON(resp)
				if err != nil {
					chat.Clients.DelClientByWS(ws)
				}
			}
		}
		mx.Unlock()
	}
}
