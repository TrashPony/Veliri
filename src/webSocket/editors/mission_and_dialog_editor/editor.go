package mission_and_dialog_editor

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/missions"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/webSocket/utils"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersWs = make(map[*websocket.Conn]*player.Player)

func AddNewUser(ws *websocket.Conn, login string, id int) {

	mutex.Lock()
	//utils.CheckDoubleLogin(login, &usersWs)
	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	usersWs[ws] = newPlayer
	println("WS mapEditor Сессия:  login: " + login + " id: " + strconv.Itoa(id))

	mutex.Unlock()

	Reader(ws)
}

type Message struct {
	Event    string                `json:"event"`
	ID       int                   `json:"id"`
	Dialogs  map[int]dialog.Dialog `json:"dialogs"`
	Dialog   *dialog.Dialog        `json:"dialog"`
	IdPage   int                   `json:"id_page"`
	Fraction string                `json:"fraction"`
	File     string                `json:"file"`
	Name     string                `json:"name"`

	Missions     map[int]*mission.Mission `json:"missions"`
	Mission      *mission.Mission         `json:"mission"`
	Slot         int                      `json:"slot"`
	ItemType     string                   `json:"item_type"`
	ItemID       int                      `json:"item_id"`
	ItemQuantity int                      `json:"item_quantity"`
}

func Reader(ws *websocket.Conn) {
	for {
		var msg Message
		err := ws.ReadJSON(&msg)

		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			println(err.Error())
			utils.DelConn(ws, &usersWs, err)
			break
		}

		if msg.Event == "OpenEditor" {
			// отсылаем все диалоги которые есть в игре
			ws.WriteJSON(&Message{Event: msg.Event, Dialogs: gameTypes.Dialogs.GetAll()})
		}

		if msg.Event == "GetMissionDialog" {

			if msg.ID == 0 {
				ws.WriteJSON(&Message{Event: "OpenEditor", Dialogs: gameTypes.Dialogs.GetAll()})
				continue
			}

			dialogs := make(map[int]dialog.Dialog)
			gameMission := missions.Missions.GetByID(msg.ID)

			if gameMission != nil {
				dialogs[gameMission.StartDialogID] = *gameTypes.Dialogs.GetByID(gameMission.StartDialogID)
				dialogs[gameMission.NotFinishedDialogId] = *gameTypes.Dialogs.GetByID(gameMission.NotFinishedDialogId)

				if gameMission.Actions != nil {
					for _, action := range gameMission.Actions {
						dialogs[action.DialogID] = *gameTypes.Dialogs.GetByID(action.DialogID)
						dialogs[action.AlternativeDialogId] = *gameTypes.Dialogs.GetByID(action.AlternativeDialogId)
					}
				}
			}

			ws.WriteJSON(&Message{Event: "OpenEditor", Dialogs: dialogs})
		}

		if msg.Event == "GetDialog" {
			// отсылаем 1 диалог по ид
			ws.WriteJSON(&Message{Event: msg.Event, Dialog: gameTypes.Dialogs.GetByID(msg.ID)})
		}

		if msg.Event == "SaveDialog" {
			gameTypes.Dialogs.UpdateTypeDialog(msg.Dialog)
			ws.WriteJSON(&Message{Event: "GetDialog", Dialog: gameTypes.Dialogs.GetByID(msg.Dialog.ID)})
		}

		if msg.Event == "CreateDialog" {
			// вносим в бд, получаем ид
			// вносим в сторедж
			gameTypes.Dialogs.AddNewDialog(msg.Name)
			ws.WriteJSON(&Message{Event: "OpenEditor", Dialogs: gameTypes.Dialogs.GetAll()})
		}

		if msg.Event == "DeleteDialog" {
			gameTypes.Dialogs.DeleteDialog(msg.ID)
			ws.WriteJSON(&Message{Event: "OpenEditor", Dialogs: gameTypes.Dialogs.GetAll()})
		}

		if msg.Event == "SetPicture" {
			// метод для установки картинка в страницу
			gameTypes.Dialogs.SetPicture(msg.ID, msg.IdPage, msg.Fraction, msg.File)
		}

		if msg.Event == "AddPage" {
			gameTypes.Dialogs.AddPage(msg.Dialog.ID)
			gameTypes.Dialogs.UpdateTypeDialog(msg.Dialog)
			ws.WriteJSON(&Message{Event: "GetDialog", Dialog: gameTypes.Dialogs.GetByID(msg.Dialog.ID)})
		}

		if msg.Event == "GetAllMissions" {
			ws.WriteJSON(&Message{Event: msg.Event, Missions: missions.Missions.GetAllMissType()})
		}

		if msg.Event == "SaveMissions" {
			missions.Missions.SaveTypeMission(msg.Mission)
			ws.WriteJSON(&Message{Event: "GetAllMissions", Missions: missions.Missions.GetAllMissType()})
		}

		if msg.Event == "DeleteMission" {
			missions.Missions.DeleteMission(msg.Mission)
			ws.WriteJSON(&Message{Event: "GetAllMissions", Missions: missions.Missions.GetAllMissType()})
		}

		if msg.Event == "AddMission" {
			missions.Missions.AddMission(msg.Name)
			ws.WriteJSON(&Message{Event: "GetAllMissions", Missions: missions.Missions.GetAllMissType()})
		}

		if msg.Event == "RemoveMissionRewardItem" {
			missions.Missions.RemoveItem(msg.ID, msg.Slot)
			ws.WriteJSON(&Message{Event: "GetAllMissions", Missions: missions.Missions.GetAllMissType()})
		}

		if msg.Event == "RemoveActionItem" {
			missions.Missions.ActionRemoveItem(msg.ID, msg.Slot)
			ws.WriteJSON(&Message{Event: "GetAllMissions", Missions: missions.Missions.GetAllMissType()})
		}

		if msg.Event == "AddMissionRewardItem" {
			missions.Missions.AddItem(msg.ID, msg.ItemID, msg.ItemQuantity, msg.ItemType)
			ws.WriteJSON(&Message{Event: "GetAllMissions", Missions: missions.Missions.GetAllMissType()})
		}

		if msg.Event == "AddActionItem" {
			missions.Missions.ActionAddItem(msg.ID, msg.ItemID, msg.ItemQuantity, msg.ItemType)
			ws.WriteJSON(&Message{Event: "GetAllMissions", Missions: missions.Missions.GetAllMissType()})
		}
	}
}
