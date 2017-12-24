package chat

import (
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"log"
)

var mutex = &sync.Mutex{}

var usersLobbyWs = make(map[*websocket.Conn]*Clients) // тут будут храниться наши подключения
var chatPipe = make(chan chatResponse)

type Clients struct { // структура описывающая клиента ws соеденение
	Login string
	Id    int
}

type chatMessage struct {
	Event    string `json:"event"`
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}

type chatResponse struct {
	Event        string `json:"event"`
	UserName     string `json:"user_name"`
	GameUser     string `json:"game_user"`
	Message		 string `json:"message"`
}

func AddNewUser(ws *websocket.Conn, login string, id int) {
	CheckDoubleLogin(login, &usersLobbyWs)
	usersLobbyWs[ws] = &Clients{Login: login, Id: id} // Регистрируем нового Клиента
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

	Reader(ws)
}


func Reader(ws *websocket.Conn) {
	for {
		var msg chatMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersLobbyWs, err)
			break
		}
		if msg.Event == "NewChatMessage" {
			var resp= chatResponse{Event: msg.Event, Message: msg.Message, GameUser: usersLobbyWs[ws].Login}
			chatPipe <- resp
		}
	}
}

func CommonChatSender()  {
	for {
		resp := <- chatPipe
		mutex.Lock()
		for ws := range usersLobbyWs {
			err := ws.WriteJSON(resp)
			if err != nil {
				log.Printf("error: %v", err)
				ws.Close()
				delete(usersLobbyWs, ws)
			}
		}
		mutex.Unlock()
	}
}

