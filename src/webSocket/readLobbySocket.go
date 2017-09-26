package webSocket

import (
	"log"
	"websocket-master"
	"net/http"
	"strconv"
	"../lobby"
)

// пайп доп. читать в документации
var openGamePipe = make(chan LobbyResponse)

var usersWs = make(map[LobbyClients]bool) // тут будут храниться наши подключения
var upgrader = websocket.Upgrader{} // методами приема обычного HTTP-соединения и обновления его до WebSocket


func ReadLobbySocket(login string, id int, w http.ResponseWriter, r *http.Request)  {

	ws, err := upgrader.Upgrade(w, r, nil) // запрос GET для перехода на протокол websocket
	if err != nil {
		log.Fatal(err)
	}




	usersWs[LobbyClients{ws, login, id}] = true // Регистрируем нового Клиента



	print("WS Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))


	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

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
			var maps string = lobby.MapList()
			var resp = LobbyResponse{"MapSelection", LoginWs(ws), "", maps, ""}
			openGamePipe <- resp // Отправляет сообщение в тред
		}

		if msg.Event == "GameSelection"{
			var games []string = lobby.OpenGameList()
			var resp = LobbyResponse{"GameSelection", LoginWs(ws), games[0], games[1], games[2]}
			openGamePipe <- resp
		}

		if msg.Event == "CreateNewGame"{
			lobby.CreateNewGame(msg.GameName, msg.MapName, LoginWs(ws))
			var resp = LobbyResponse{"CreateNewGame", LoginWs(ws), "", "", ""}
			openGamePipe <- resp
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
					lobby.DelNewGame(client.login)
					client.ws.Close()
					delete(usersWs, client)
				}
			}
		}
	}
}

func LoginWs(ws *websocket.Conn) (string)  {
	for client := range usersWs { // ходим по всем подключениям
		if(client.ws == ws) {
			return client.login
		}
	}
	return ""
}

func DelConn(ws *websocket.Conn, err error)  {
	log.Printf("error: %v", err)
	for client := range usersWs { // ходим по всем подключениям
		if(client.ws == ws) { // находим подключение с ошибкой
			lobby.DelNewGame(client.login)
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
