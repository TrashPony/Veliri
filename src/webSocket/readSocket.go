package webSocket

import (
	"log"
	"websocket-master"
	"net/http"
	"strconv"
)

var lobby = make(chan LobbyMessage) // пайп доп. читать в документации
var usersWs = make(map[Clients]bool) // тут будут храниться наши подключения
var upgrader = websocket.Upgrader{} // методами приема обычного HTTP-соединения и обновления его до WebSocket


func ReadSocket(login string, id int, w http.ResponseWriter, r *http.Request)  {

	ws, err := upgrader.Upgrade(w, r, nil) // запрос GET для перехода на протокол websocket
	if err != nil {
		log.Fatal(err)
	}

	usersWs[Clients{ws, login, id}] = true // Регистрируем нового Клиента


	for client := range usersWs { // просто смотрим кто есть в подключениях
		print("WS Сессия: ")
		print(client.ws)
		println(" login: " + client.login + " id: " + strconv.Itoa(client.id))
	}

	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
	LobbyReader(ws)
}

func Router(ws *websocket.Conn)  {
	for {
		var msg RouterMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, err)
			break
		}
		if msg.Target == "lobby" {
			LobbyReader(ws) // Отправляет вновь принятое сообщение на широковещательный канал
		}
	}
}

func LobbyReader(ws *websocket.Conn)  {
	for {
		var msg LobbyMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message

		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, err)
			break
		}
		lobby <- msg // Отправляет вновь принятое сообщение на широковещательный канал
	}
}

func LobbySender() {
	for {
		// Берет сообщение из общего канала
		msg := <-lobby
		// Отправляет его каждому клиенту
		// оп оп тут надо сделать так что бы знать как брать нужного клиента
		for client := range usersWs {
			//if client.login == "admin" { // ищем юзера админ и только ему отправляем сообщение
			err := client.ws.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.ws.Close()
				delete(usersWs, client)
			}
			//}
		}
	}
}

func DelConn(ws *websocket.Conn, err error)  {
	log.Printf("error: %v", err)
	for client := range usersWs { // ходим по всем подключениям
		if(client.ws == ws) { // находим подключение с ошибкой
			delete(usersWs, client) // удаляем его из активных подключений
			break
		}
	}
}


type RouterMessage struct { // структура описывающие падающие в общий канал сообщения
	Target   string `json:"target"` // обратные ковычки позволяют парсить json значения
}

type LobbyMessage struct { // структура описывающие падающие в общий канал сообщения
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Clients struct { // структура описывающая клиента ws соеденение
	ws *websocket.Conn
	login string
	id int
}
