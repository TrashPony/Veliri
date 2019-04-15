package chat

import (
	"github.com/TrashPony/Veliri/src/mechanics/chat"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/gorilla/websocket"
)

var chatPipe = make(chan chatMessage)

type chatMessage struct {
	Event    string              `json:"event"`
	UserName string              `json:"user_name"`
	Message  string              `json:"message"`
	GroupID  int                 `json:"group_id"`
	UserID   int                 `json:"user_id"`
	Group    *chat.Group         `json:"group"`
	Groups   map[int]*chat.Group `json:"groups"`
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

		client := localGame.Clients.GetByWs(ws)

		if client != nil {
			if msg.Event == "OpenChat" {
				SendMessage(msg.Event, msg.Message, client.GetID(), 0, chat.Groups.GetGroup(msg.GroupID), chat.Groups.GetAllUserGroups(client))
			}

			if msg.Event == "ChangeGroup" {
				if chat.Groups.CheckUserSubscribe(msg.GroupID, client) {
					SendMessage(msg.Event, msg.Message, client.GetID(), 0, chat.Groups.GetGroup(msg.GroupID), nil)
				}
			}

			if msg.Event == "NewChatMessage" {
				if chat.Groups.CheckUserSubscribe(msg.GroupID, client) {
					SendMessage(msg.Event, msg.Message, 0, msg.GroupID, nil, nil)
				}
			}
		}
	}
}

// todo уменьшить количество арументов и сделать как в field)
func SendMessage(event, senderMessage string, userID, GroupID int, group *chat.Group, groups map[int]*chat.Group) {
	chatPipe <- chatMessage{Event: event, Message: senderMessage, GroupID: GroupID, UserID: userID, Group: group, Groups: groups}
}

func CommonChatSender() {
	for {
		resp := <-chatPipe

		users, mx := chat.Clients.GetAllConnects()
		for ws, client := range users {
			if chat.Groups.CheckUserSubscribe(resp.GroupID, client) && resp.UserID == 0 { // отправляем всей подписоте
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
