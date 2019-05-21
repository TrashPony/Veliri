package other

import (
	"github.com/TrashPony/Veliri/src/mechanics/chat"
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/chatGroup"
	dialogObj "github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/skill"
	"github.com/gorilla/websocket"
)

var respPipe = make(chan Message)

type Message struct {
	service     string
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
	Notify      *player.Notify              `json:"notify"`
	Notifys     map[string]*player.Notify   `json:"notifys"`

	File         string                     `json:"file"`
	Biography    string                     `json:"biography"`
	Player       *player.Player             `json:"player"`
	ID           int                        `json:"id"`
	Skill        skill.Skill                `json:"skill"`
	Error        string                     `json:"error"`
	Count        int                        `json:"count"`
	Maps         map[int]*_map.ShortInfoMap `json:"maps"`
	SearchMaps   []*maps.SearchMap          `json:"search_maps"`
	DialogPage   *dialogObj.Page            `json:"dialog_page"`
	DialogAction string                     `json:"dialog_action"`
	ToPage       int                        `json:"to_page"`
	AskID        int                        `json:"ask_id"`
	Mission      *mission.Mission           `json:"mission"`
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

		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			chat.Clients.DelClientByWS(ws)
			break
		}

		client := chat.Clients.GetByWs(ws)

		// все что связано с чатом выплюнул сюда :\
		if msg.Event == "OpenChat" || msg.Event == "GetAllGroups" || msg.Event == "ChangeGroup" || msg.Event == "SubscribeGroup" ||
			msg.Event == "Unsubscribe" || msg.Event == "CreateNewGroup" || msg.Event == "NewChatMessage" {
			chatReader(client, msg)
		}

		// остальное
		if client != nil {

			if msg.Event == "training" {
				client.Training = msg.Count
				dbPlayer.UpdateUser(client)
			}

			if msg.Event == "LoadAvatar" {
				client.AvatarIcon = msg.File
				dbPlayer.UpdateUser(client)
			}

			if msg.Event == "SetBiography" {
				client.Biography = msg.Biography
				dbPlayer.UpdateUser(client)
			}

			if msg.Event == "OpenUserStat" {
				sendOtherMessage(Message{Event: msg.Event, UserID: client.GetID(), Player: client})
			}

			if msg.Event == "upSkill" {
				userSkill, ok := client.UpSkill(msg.ID)
				if ok {
					sendOtherMessage(Message{Event: "upSkill", UserID: client.GetID(), Player: client, Skill: *userSkill})
					dbPlayer.UpdateUser(client)
				} else {
					sendOtherMessage(Message{Event: "upSkill", UserID: client.GetID(), Error: "no points"})
				}
			}

			if msg.Event == "openMapMenu" {
				userBase, _ := bases.Bases.Get(client.InBaseID)
				sendOtherMessage(Message{Event: msg.Event, UserID: client.GetID(), Maps: maps.Maps.GetAllShortInfoMap(), ID: userBase.MapID})
			}

			if msg.Event == "previewPath" {
				userBase, _ := bases.Bases.Get(client.InBaseID)
				if userBase.MapID != msg.ID {
					searchMaps, _ := maps.Maps.FindGlobalPath(userBase.MapID, msg.ID)
					sendOtherMessage(Message{Event: msg.Event, UserID: client.GetID(), SearchMaps: searchMaps})
				}
			}

			if msg.Event == "OpenDialog" {

			}

			if msg.Event == "Ask" {
				page, err, action, mission := dialog.Ask(client, client.GetOpenDialog(), "base", msg.ToPage, msg.AskID)
				if client.InBaseID > 0 && err == nil {
					sendOtherMessage(Message{Event: "dialog", UserID: client.GetID(), DialogPage: page, DialogAction: action, Mission: mission})
				} else {
					sendOtherMessage(Message{Event: "Error", UserID: client.GetID(), Error: err.Error()})
				}
			}

			if msg.Event == "openDepartmentOfEmployment" {
				userBase, _ := bases.Bases.Get(client.InBaseID)
				page, _ := dialog.GetBaseGreeting(client, userBase)
				sendOtherMessage(Message{Event: msg.Event, UserID: client.GetID(), DialogPage: page})
			}
		}
	}
}

func sendOtherMessage(message Message) {
	message.service = "other"
	respPipe <- message
}

func ReposeSender() {
	for {
		resp := <-respPipe
		if resp.service == "chat" {
			CommonChatSender(&resp)
		}

		if resp.service == "other" {
			users, mx := chat.Clients.GetAllConnects()
			for ws, client := range users {
				if client.GetID() == resp.UserID {
					err := ws.WriteJSON(resp)
					if err != nil {
						chat.Clients.DelClientByWS(ws)
					}
				}
			}
			mx.Unlock()
		}
	}
}
