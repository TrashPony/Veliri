package webSocket

import (
	"log"
	"websocket-master"
	"net/http"
	"strconv"
	"../lobby"
)

var lobbyPipe = make(chan LobbyMessage) // пайп доп. читать в документации
var openGamePipe = make(chan LobbyResponse)

var usersWs = make(map[LobbyClients]bool) // тут будут храниться наши подключения
var wsLogin = ""
var upgrader = websocket.Upgrader{} // методами приема обычного HTTP-соединения и обновления его до WebSocket


func ReadLobbySocket(login string, id int, w http.ResponseWriter, r *http.Request)  {

	ws, err := upgrader.Upgrade(w, r, nil) // запрос GET для перехода на протокол websocket
	if err != nil {
		log.Fatal(err)
	}

	usersWs[LobbyClients{ws, login, id}] = true // Регистрируем нового Клиента


	for client := range usersWs { // просто смотрим кто есть в подключениях
		print("WS Сессия: ")
		print(client.ws)
		println(" login: " + client.login + " id: " + strconv.Itoa(client.id))
	}

	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
	wsLogin = login
	LobbyReader(ws)
}

func LobbyReader(ws *websocket.Conn)  {
	for {
		var msg LobbyMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message

		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, err)
			break
		}



		if msg.Event == "MapSelection"{
			lobby.MapList()
		}

		if msg.Event == "GameSelection"{
			var games []string = lobby.OpenGameList()
			var resp = LobbyResponse{"GameSelection", wsLogin, games[0], games[1], games[2]}
			openGamePipe <- resp
		}

		if msg.Event == "CrateNewGame"{
			lobby.CreateNewGame()
		}

		  // Отправляет сообщение в тред
	}
}

func LobbySender() {
	for {
		// Берет сообщение из общего канала
		msg := <-lobbyPipe
		// Отправляет его каждому клиенту
		// оп оп тут надо сделать так что бы знать как брать нужного клиента
		for client := range usersWs {
			if client.login == msg.UserName {// ищем юзера который отправил сообщение и только ему отправляем
				err := client.ws.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					client.ws.Close()
					delete(usersWs, client)
				}
			}

		}
	}
}

func OpenGameSender() {
	for {
		resp := <-openGamePipe
		for client := range usersWs {
			if client.login == resp.UserName {
				err := client.ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					client.ws.Close()
					delete(usersWs, client)
				}
			}
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

type LobbyMessage struct {
	Event    string `json:"event"`
	MapName  string `json:"map_name"`
	UserName string `json:"user_name"`
	GameName string `json:"game_name"`
}

type LobbyResponse struct {
	Event    	  string `json:"event"`
	UserName	  string
	ResponseNameGame  string `json:"response_name_game"`
	ResponseNameMap   string `json:"response_name_map"`
	ResponseNameUser  string `json:"response_name_user"`
}

type  LobbyClients struct { // структура описывающая клиента ws соеденение
	ws *websocket.Conn
	login string
	id int
}
