package chat

import (
	"github.com/TrashPony/Veliri/src/mechanics/chat"
	"github.com/TrashPony/Veliri/src/mechanics/factories/chats"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
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
}

func AddNewUser(ws *websocket.Conn, login string, id int) {
	newPlayer, ok := players.Users.Get(id)
	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	chat.Clients.AddNewClient(ws, newPlayer) // Регистрируем нового Клиента
	Reader(ws)
}

func UserOnlineChecker() {
	for {
		for _, group := range chats.Groups.GetAllGroups() {
			update := false
			for id := range group.Users {
				chatUser := chat.Clients.GetByID(id)
				if group.Users[id] && chatUser == nil {
					group.Users[id] = false
					update = true
				}

				if !group.Users[id] && chatUser != nil {
					group.Users[id] = true
					update = true
				}
			}
			if update {
				SendMessage("UpdateUsers", nil, 0, group.ID, group, nil, getUsersInChatGroup(group))
			}
		}

		time.Sleep(1 * time.Second)
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
				SendMessage(msg.Event, msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID), chats.Groups.GetAllUserGroups(client), nil)
			}

			if msg.Event == "GetAllGroups" {
				SendMessage(msg.Event, msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID), chats.Groups.GetAllGroups(), nil)
			}

			if msg.Event == "ChangeGroup" {
				if chats.Groups.CheckUserSubscribe(msg.GroupID, client) {

					group := chats.Groups.GetGroup(msg.GroupID)

					SendMessage(msg.Event, msg.Message, client.GetID(), 0, group, nil, getUsersInChatGroup(group))
				}
			}

			if msg.Event == "SubscribeGroup" {
				if chats.Groups.SubscribeGroup(msg.GroupID, client, msg.Password) {
					SendMessage("OpenChat", msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID), chats.Groups.GetAllUserGroups(client), nil)
				}
			}

			if msg.Event == "NewChatMessage" {
				if chats.Groups.CheckUserSubscribe(msg.GroupID, client) {
					// добавляем в историю, отправляем не сообщение текстом а обьектом
					group := chats.Groups.GetGroup(msg.GroupID)

					chatMessage := chatGroup.Message{UserName: client.GetLogin(), AvatarIcon: client.AvatarIcon, Message: msg.MessageText, Time: time.Now().UTC()}
					group.History = append(group.History, &chatMessage)

					SendMessage(msg.Event, &chatMessage, 0, msg.GroupID, nil, nil, nil)
				}
			}
		}
	}
}

func getUsersInChatGroup(group *chatGroup.Group) []*player.ShortUserInfo {
	users := make([]*player.ShortUserInfo, 0)

	for id := range group.Users {
		chatUser := chat.Clients.GetByID(id)
		if chatUser != nil {
			users = append(users, chatUser.GetShortUserInfo(false))
		} else {
			group.Users[id] = false
		}
	}

	return users
}

func SendMessage(event string, senderMessage *chatGroup.Message, userID, GroupID int, group *chatGroup.Group, groups map[int]*chatGroup.Group, users []*player.ShortUserInfo) {
	chatPipe <- chatMessage{Event: event, Message: senderMessage, GroupID: GroupID, UserID: userID, Group: group, Groups: groups, Users: users}
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

			if resp.GroupID == 0 && resp.UserID == client.GetID() { // отправляем только инициатору
				err := ws.WriteJSON(resp)
				if err != nil {
					chat.Clients.DelClientByWS(ws)
				}
			}
		}
		mx.Unlock()
	}
}
