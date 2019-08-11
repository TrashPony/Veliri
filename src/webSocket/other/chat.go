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
				chats.Groups.GetAllowUserGroups(client), nil, false, nil, nil, nil, nil)
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

		if msg.Event == "CreateNewChatGroup" {
			// todo проверка полей на наличие опасных символов
			newGroup, err := chats.Groups.CreateNewGroup(client, msg.Name, msg.Password, msg.File, msg.Greetings)
			if err != nil {
				sendOtherMessage(Message{Event: "Error", UserID: client.GetID(), Error: err.Error()})
				return
			}

			// открываем чат у создателя канала
			SendMessage("OpenChat", msg.Message, client.GetID(), 0, newGroup,
				chats.Groups.GetAllUserGroups(client), nil, false, nil, nil, nil, nil)
		}

		if msg.Event == "CreateNewPrivateGroup" && msg.UserID != client.GetID() {
			// создаем приватный чат
			toUser := chat.Clients.GetByID(msg.UserID)
			privateGroup, sysMessage := chats.Groups.CreateNewPrivateGroup(client, toUser)

			// открываем чат у инициалитора общения
			SendMessage("OpenChat", msg.Message, client.GetID(), 0, privateGroup,
				chats.Groups.GetAllUserGroups(client), nil, false, nil, nil, nil, nil)

			// говорим второму что у него есть новый чат, но не открываем его
			SendMessage("OpenChat", msg.Message, toUser.GetID(), 0, nil,
				chats.Groups.GetAllUserGroups(toUser), nil, false, nil, nil, nil, nil)

			if sysMessage != nil {
				SendMessage("NewChatMessage", sysMessage, 0, privateGroup.ID, nil, nil, nil, false, nil, nil, nil, nil)
			}
		}

		if msg.Event == "SubscribeGroup" {
			group := chats.Groups.GetGroup(msg.GroupID)
			ok, err := chats.Groups.SubscribeGroup(msg.GroupID, client, msg.Password)
			if err != nil {
				sendOtherMessage(Message{Event: "Error", UserID: client.GetID(), Error: err.Error()})
			}

			if ok {
				// открываем чат у игрока
				SendMessage("OpenChat", msg.Message, client.GetID(), 0, chats.Groups.GetGroup(msg.GroupID),
					chats.Groups.GetAllUserGroups(client), nil, false, nil, nil, nil, nil)
				// оповещаем группу что у них новый юзер
				SendMessage("UpdateUsers", nil, 0, msg.GroupID, group, nil,
					getUsersInChatGroup(group, true), false, nil, nil, nil, nil)
			}
		}

		if msg.Event == "Unsubscribe" {
			localGroup, _ := getLocalChat(client)
			group := chats.Groups.GetGroup(msg.GroupID)
			if group != nil {

				if group.Private {
					systemMessage := &chatGroup.Message{
						Message: "игрок " + client.GetLogin() + " покинул этот чат.",
						Time:    time.Now().UTC(),
						System:  true,
					}
					group.History = append(group.History, systemMessage)
					SendMessage("NewChatMessage", systemMessage, 0, msg.GroupID, nil, nil, nil, false, nil, nil, nil, nil)
				}

				chats.Groups.Unsubscribe(msg.GroupID, client)

				// оповещаем всех в чате что игровы вышел из группы.
				SendMessage("UpdateUsers", nil, 0, msg.GroupID, group, nil,
					getUsersInChatGroup(group, true), false, nil, nil, nil, nil)

				// у игрока который вышел открывается локальный чат
				SendMessage("OpenChat", msg.Message, client.GetID(), 0, localGroup,
					chats.Groups.GetAllUserGroups(client), nil, false, nil, nil, nil, nil)

			}
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
				if msg.GroupID == 0 {
					group, _ := getLocalChat(client)

					chatMessage := chatGroup.Message{UserName: client.GetLogin(), UserID: strconv.Itoa(client.GetID()), Message: msg.MessageText, Time: time.Now().UTC()}
					group.History = append(group.History, &chatMessage)

					SendMessage(msg.Event, &chatMessage, 0, msg.GroupID, group, nil, nil, true, nil, nil, nil, nil)
				}
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
