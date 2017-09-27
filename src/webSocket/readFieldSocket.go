package webSocket

import (
	"net/http"
	"log"
	"strconv"
	"websocket-master"
)

var FieldPipe = make(chan FieldResponse)

var usersFieldWs = make(map[FieldClients]bool) // тут будут храниться наши подключения

func ReadFieldSocket(login string, id int, w http.ResponseWriter, r *http.Request)  {

	ws, err := upgrader.Upgrade(w, r, nil) // запрос GET для перехода на протокол websocket
	if err != nil {
		log.Fatal(err)
	}

	usersFieldWs[FieldClients{ws, login, id}] = true // Регистрируем нового Клиента

	print("WS field Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))

	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

	FieldReader(ws)
}

func FieldReader(ws *websocket.Conn)  {
	for {
		var msg FieldMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelFieldConn(ws, err)
			break
		}
	}
}

func LoginFieldWs(ws *websocket.Conn) (string)  {
	for client := range usersFieldWs { // ходим по всем подключениям
		if(client.ws == ws) {
			return client.login
		}
	}
	return ""
}

func FieldReposeSender() {
	for {
		resp := <-FieldPipe
		for client := range usersFieldWs {
			if client.login == resp.UserName {
				err := client.ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					client.ws.Close()
					delete(usersFieldWs, client)
				}
			}
		}
	}
}

func DelFieldConn(ws *websocket.Conn, err error)  {
	log.Printf("error: %v", err)
	for client := range usersFieldWs { // ходим по всем подключениям
		if(client.ws == ws) { // находим подключение с ошибкой
			delete(usersFieldWs, client) // удаляем его из активных подключений
			break
		}
	}
}

type FieldMessage struct {
	Event    string `json:"event"`
	MapName  string `json:"map_name"`
	UserName string `json:"user_name"`
	GameName string `json:"game_name"`
}

type FieldResponse struct {
	Event    	  string `json:"event"`
	UserName	  string
	ResponseNameGame  string `json:"response_name_game"`
	ResponseNameMap   string `json:"response_name_map"`
	ResponseNameUser  string `json:"response_name_user"`
	ResponseNameUser2 string `json:"response_name_user_2"`
}

type FieldClients struct { // структура описывающая клиента ws соеденение
	ws *websocket.Conn
	login string
	id int
}