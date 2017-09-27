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

var usersLobbyWs = make(map[LobbyClients]bool) // тут будут храниться наши подключения
var upgrader = websocket.Upgrader{} // методами приема обычного HTTP-соединения и обновления его до WebSocket


func ReadLobbySocket(login string, id int, w http.ResponseWriter, r *http.Request)  {

	ws, err := upgrader.Upgrade(w, r, nil) // запрос GET для перехода на протокол websocket
	if err != nil {
		log.Fatal(err)
	}

	usersLobbyWs[LobbyClients{ws, login, id}] = true // Регистрируем нового Клиента

	print("WS lobby Сессия: ") // просто смотрим новое подключение
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
			DelLobbyConn(ws, err)
			break
		}

		if msg.Event == "MapView"{
			var maps = lobby.MapList()
			var resp = LobbyResponse{"MapView", LoginLobbyWs(ws), "", maps, "", ""}
			openGamePipe <- resp // Отправляет сообщение в тред
		}

		if msg.Event == "GameView"{
			var games []string = lobby.OpenGameList()
			var resp = LobbyResponse{"GameView", LoginLobbyWs(ws), games[0], games[1], games[2], ""}
			openGamePipe <- resp
		}

		if msg.Event == "DontEndGames"{
			var gameName string = lobby.DontEndGames(LoginLobbyWs(ws))
			var resp = LobbyResponse{"DontEndGames", LoginLobbyWs(ws), gameName, "", "", ""}
			openGamePipe <- resp
		}

		if msg.Event == "ConnectGame"{
			// ваще он должен попадать в лоби но сейчас он сразу регистрирует новую игру
			user2 , success := lobby.ConnectGame(msg.GameName, LoginLobbyWs(ws))
			var resp = LobbyResponse{"GameView", LoginLobbyWs(ws), strconv.FormatBool(success), "", "" , user2}
			openGamePipe <- resp
		}

		if msg.Event == "CreateNewGame"{
			lobby.CreateNewGame(msg.GameName, msg.MapName, LoginLobbyWs(ws))
			var resp = LobbyResponse{"CreateNewGame", LoginLobbyWs(ws), "", "", "", ""}
			openGamePipe <- resp
		}

		if msg.Event == "StartNewGame"{
			// а вот этот метод должен после того как все подтвердили создавать новую игру в бд и перекидывать на новую страницу
			//lobby.StartNewGame(msg.MapName, msg.UserName)
		}
	}
}

func LobbyReposeSender() {
	for {
		resp := <-openGamePipe
		for client := range usersLobbyWs {
			if client.login == resp.UserName {
				err := client.ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					lobby.DelNewGame(client.login)
					client.ws.Close()
					delete(usersLobbyWs, client)
				}
			}
		}
	}
}

func LoginLobbyWs(ws *websocket.Conn) (string)  {
	for client := range usersLobbyWs { // ходим по всем подключениям
		if(client.ws == ws) {
			return client.login
		}
	}
	return ""
}

func DelLobbyConn(ws *websocket.Conn, err error)  {
	log.Printf("error: %v", err)
	for client := range usersLobbyWs { // ходим по всем подключениям
		if(client.ws == ws) { // находим подключение с ошибкой
			lobby.DelNewGame(client.login)
			delete(usersLobbyWs, client) // удаляем его из активных подключений
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
	ResponseNameUser2 string `json:"response_name_user_2"`
}

type  LobbyClients struct { // структура описывающая клиента ws соеденение
	ws *websocket.Conn
	login string
	id int
}
