package other

import (
	"github.com/TrashPony/Veliri/src/mechanics/chat"
	"github.com/TrashPony/Veliri/src/mechanics/factories/chats"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"strconv"
	"time"
)

func chatReader(client *player.Player, msg Message) {
	if client != nil {
		if msg.Event == "OpenChat" {
			group, _ := getLocalChat(client)
			SendMessage(msg.Event, msg.Message, client.GetID(), 0, group, chats.Groups.GetAllUserGroups(client),
				nil, false, client.GetShortUserInfo(false), client.Missions, nil, client.NotifyQueue)
		}

		if msg.Event == "GetAllGroups" {
			SendMessage(msg.Event, msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID),
				chats.Groups.GetAllGroups(), nil, false, nil, nil, nil, nil)
		}

		if msg.Event == "ChangeGroup" {
			if chats.Groups.CheckUserSubscribe(msg.GroupID, client) {
				group := chats.Groups.GetGroup(msg.GroupID)
				SendMessage(msg.Event, msg.Message, client.GetID(), 0, group, nil,
					getUsersInChatGroup(group, true), false, nil, nil, nil, nil)
			} else {
				group, _ := getLocalChat(client)
				SendMessage(msg.Event, msg.Message, client.GetID(), 0, group, nil,
					getUsersInChatGroup(group, false), false, nil, nil, nil, nil)
			}

		}

		if msg.Event == "SubscribeGroup" {
			group := chats.Groups.GetGroup(msg.GroupID)
			if chats.Groups.SubscribeGroup(msg.GroupID, client, msg.Password) {
				SendMessage("OpenChat", msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID),
					chats.Groups.GetAllUserGroups(client), nil, false, nil, nil, nil, nil)
				SendMessage("UpdateUsers", nil, 0, msg.GroupID, group, nil,
					getUsersInChatGroup(group, true), false, nil, nil, nil, nil)
			}
		}

		if msg.Event == "Unsubscribe" {
			group := chats.Groups.GetGroup(msg.GroupID)
			if group != nil {
				chats.Groups.Unsubscribe(msg.GroupID, client)
				SendMessage("UpdateUsers", nil, 0, msg.GroupID, group, nil,
					getUsersInChatGroup(group, true), false, nil, nil, nil, nil)
				SendMessage("OpenChat", msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID),
					chats.Groups.GetAllUserGroups(client), nil, false, nil, nil, nil, nil)
			}
		}

		if msg.Event == "CreateNewGroup" {

		}

		if msg.Event == "NewChatMessage" {
			if chats.Groups.CheckUserSubscribe(msg.GroupID, client) {
				// добавляем в историю, отправляем не сообщение текстом а обьектом
				group := chats.Groups.GetGroup(msg.GroupID)

				chatMessage := chatGroup.Message{UserName: client.GetLogin(), UserID: strconv.Itoa(client.GetID()), Message: msg.MessageText, Time: time.Now().UTC()}
				group.History = append(group.History, &chatMessage)

				SendMessage(msg.Event, &chatMessage, 0, msg.GroupID, nil, nil, nil, false, nil, nil, nil, nil)
			} else {
				// если msg.GroupID == 0 то это сообщение в локальный чат
				group, _ := getLocalChat(client)

				chatMessage := chatGroup.Message{UserName: client.GetLogin(), UserID: strconv.Itoa(client.GetID()), Message: msg.MessageText, Time: time.Now().UTC()}
				group.History = append(group.History, &chatMessage)

				SendMessage(msg.Event, &chatMessage, 0, msg.GroupID, group, nil, nil, true, nil, nil, nil, nil)
			}
		}
	}
}

func SendMessage(event string, senderMessage *chatGroup.Message, userID, GroupID int, group *chatGroup.Group,
	groups map[int]*chatGroup.Group, users []*player.ShortUserInfo, local bool, user *player.ShortUserInfo,
	missions map[string]*mission.Mission, notify *player.Notify, notifys map[string]*player.Notify) {

	respPipe <- Message{
		service:  "chat",
		Event:    event,
		Message:  senderMessage,
		GroupID:  GroupID,
		UserID:   userID,
		Group:    group,
		Groups:   groups,
		Users:    users,
		Local:    local,
		User:     user,
		Missions: missions,
		Notify:   notify,
		Notifys:  notifys,
	}
}

func CommonChatSender(resp *Message) {

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
