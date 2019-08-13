package dialogEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
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

	utils.CheckDoubleLogin(login, &usersWs)

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
	Event   string                `json:"event"`
	ID      int                   `json:"id"`
	Dialogs map[int]dialog.Dialog `json:"dialogs"`
	Dialog  *dialog.Dialog        `json:"dialog"`
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

		if msg.Event == "GetDialog" {
			// отсылаем 1 диалог по ид
			ws.WriteJSON(&Message{Event: msg.Event, Dialog: gameTypes.Dialogs.GetByID(msg.ID)})
		}

		// все измения/создания страниц, ответов происходит на фронте, потом сюда прилетает полноценный диалог
		// и он заменяется. Главное проверять что бы небыло дубликатов по номерам страниц и ответов
		if msg.Event == "EditDialog" {
			gameTypes.Dialogs.UpdateTypeDialog(msg.Dialog)
			ws.WriteJSON(&Message{Event: msg.Event, Dialogs: gameTypes.Dialogs.GetAll()})
		}

		if msg.Event == "CreateDialog" {
			// вносим в бд, получаем ид
			// вносим в сторедж
			gameTypes.Dialogs.AddNewDialog(msg.Dialog)
		}
	}
}
