package chat

import (
	"github.com/TrashPony/Veliri/src/mechanics/chat"
	"github.com/TrashPony/Veliri/src/mechanics/factories/chats"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/gorilla/websocket"
	"time"
)

var chatPipe = make(chan chatMessage)

type chatMessage struct {
	Event       string                      `json:"event"`
	UserName    string                      `json:"user_name"`
	MessageText string                      `json:"message_text"`
	Message     *chatGroup.Message          `json:"message"`
	GroupID     int                         `json:"group_id"`
	UserID      int                         `json:"user_id"`
	Group       *chatGroup.Group            `json:"group"`
	Groups      map[int]*chatGroup.Group    `json:"groups"`
	Password    string                      `json:"password"`
	Users       []*player.ShortUserInfo     `json:"users"`
	User        *player.ShortUserInfo       `json:"user"`
	Local       bool                        `json:"local"`
	Missions    map[string]*mission.Mission `json:"missions"`
}

func AddNewUser(ws *websocket.Conn, login string, id int) {
	newPlayer, ok := players.Users.Get(id)
	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	chat.Clients.AddNewClient(ws, newPlayer) // Регистрируем нового Клиента
	Reader(ws)
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
				SendMessage(msg.Event, msg.Message, client.GetID(), 0, group, chats.Groups.GetAllUserGroups(client),
					nil, false, client.GetShortUserInfo(false, true), client.Missions)
			}

			if msg.Event == "GetAllGroups" {
				SendMessage(msg.Event, msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID),
					chats.Groups.GetAllGroups(), nil, false, nil, nil)
			}

			if msg.Event == "ChangeGroup" {
				if chats.Groups.CheckUserSubscribe(msg.GroupID, client) {
					group := chats.Groups.GetGroup(msg.GroupID)
					SendMessage(msg.Event, msg.Message, client.GetID(), 0, group, nil,
						getUsersInChatGroup(group, true), false, nil, nil)
				} else {
					group, _ := getLocalChat(client)
					SendMessage(msg.Event, msg.Message, client.GetID(), 0, group, nil,
						getUsersInChatGroup(group, false), false, nil, nil)
				}

			}

			if msg.Event == "SubscribeGroup" {
				group := chats.Groups.GetGroup(msg.GroupID)
				if chats.Groups.SubscribeGroup(msg.GroupID, client, msg.Password) {
					SendMessage("OpenChat", msg.Message, client.GetID(), 0,
						chats.Groups.GetGroup(msg.GroupID), chats.Groups.GetAllUserGroups(client), nil, false, nil, nil)
					SendMessage("UpdateUsers", nil, 0, msg.GroupID, group, nil,
						getUsersInChatGroup(group, true), false, nil, nil)
				}
			}

			if msg.Event == "Unsubscribe" {
				group := chats.Groups.GetGroup(msg.GroupID)
				if group != nil {
					chats.Groups.Unsubscribe(msg.GroupID, client)
					SendMessage("UpdateUsers", nil, 0, msg.GroupID, group, nil,
						getUsersInChatGroup(group, true), false, nil, nil)
					SendMessage("OpenChat", msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID),
						chats.Groups.GetAllUserGroups(client), nil, false, nil, nil)
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

					SendMessage(msg.Event, &chatMessage, 0, msg.GroupID, nil, nil, nil, false, nil, nil)
				} else {
					// если msg.GroupID == 0 то это сообщение в локальный чат
					group, _ := getLocalChat(client)

					chatMessage := chatGroup.Message{UserName: client.GetLogin(), AvatarIcon: client.AvatarIcon, Message: msg.MessageText, Time: time.Now().UTC()}
					group.History = append(group.History, &chatMessage)

					SendMessage(msg.Event, &chatMessage, 0, msg.GroupID, group, nil, nil, true, nil, nil)
				}
			}
		}
	}
}

func SendMessage(event string, senderMessage *chatGroup.Message, userID, GroupID int, group *chatGroup.Group,
	groups map[int]*chatGroup.Group, users []*player.ShortUserInfo, local bool, user *player.ShortUserInfo, missions map[string]*mission.Mission) {

	chatPipe <- chatMessage{
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
