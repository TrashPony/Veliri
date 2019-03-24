package chat

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersChatWs = make(map[*websocket.Conn]*player.Player) // тут будут храниться наши подключения
var chatPipe = make(chan chatResponse)

type chatMessage struct {
	Event    string `json:"event"`
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}

type chatResponse struct {
	Event    string `json:"event"`
	UserName string `json:"user_name"`
	GameUser string `json:"game_user"`
	Message  string `json:"message"`
}

func AddNewUser(ws *websocket.Conn, login string, id int) {

	mutex.Lock()

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	usersChatWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS chat Сессия: ") // просто смотрим новое подключение
	println(" login: " + newPlayer.GetLogin() + " id: " + strconv.Itoa(newPlayer.GetID()))

	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция завершается

	mutex.Unlock()

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	for {
		var msg chatMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil {          // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			// TODO concurrent map writes
			DelConn(ws, &usersChatWs, err)
			break
		}
		if msg.Event == "NewChatMessage" {
			var resp = chatResponse{Event: msg.Event, Message: msg.Message, GameUser: usersChatWs[ws].GetLogin()}
			chatPipe <- resp
		}
	}
}

func CommonChatSender() {
	for {
		resp := <-chatPipe
		mutex.Lock()
		for ws := range usersChatWs {
			err := ws.WriteJSON(resp)
			if err != nil {
				log.Printf("error: %v", err)
				ws.Close()
				delete(usersChatWs, ws)
			}
		}
		mutex.Unlock()
	}
}
