package auth

import (
	"log"
	"net/http"
	"websocket-master"
	"strconv"
)

var usersWs = make(map[Clients]bool) // тут будут храниться наши подключения
var broadcast = make(chan Message) // широковещательный канал доп. читать в документации
var upgrader = websocket.Upgrader{} // методами приема обычного HTTP-соединения и обновления его до WebSocket

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	var login string
	var id int

	login, id = CheckCookie(w, r) // берем из куки данные по логину и ид пользовтеля

	if login == "" || id == 0  || login == "anonymous" {
		println("Соеденение не разрешено: не авторизован")
		//http.Redirect(w, r, "http://www.google.com", 401)
		return // если человек не авторизован то ему не разрешается соеденение
	}

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

	for {
		var msg Message
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message

		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			log.Printf("error: %v", err)
			for client := range usersWs { // ходим по всем подключениям
				if(client.ws == ws) { // находим подключение с ошибкой
					delete(usersWs, client) // удаляем его из активных подключений
					break
				}
			}
			break
		}
		broadcast <- msg // Отправляет вновь принятое сообщение на широковещательный канал
	}
}

func HandleMessages() {
	for {
		// Берет сообщение из общего канала
		msg := <-broadcast
		// Отправляет его каждому клиенту
		// оп оп тут надо сделать так что бы знать как брать нужного клиента
		for client := range usersWs {
			if client.login == "admin" { // ищем юзера админ и только ему отправляем сообщение
				client.ws.WriteJSON(msg)
			}
			// тут надо будет обработать ошибку
		}
	}
}

type Message struct { // структура описывающие падающие в общий канал сообщения
	Email    string `json:"email"` // обратные ковычки позволяют парсить json значения
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Clients struct { // структура описывающая клиента ws соеденение
	ws *websocket.Conn
	login string
	id int
}