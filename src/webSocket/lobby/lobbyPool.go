package lobby

import (
	"../../mechanics/lobby"
	"../../mechanics/player"

	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersLobbyWs = make(map[*websocket.Conn]*player.Player) // тут будут храниться наши подключения
var lobbyPipe = make(chan Response)
var openGames = make(map[int]*lobby.LobbyGames)

func AddNewUser(ws *websocket.Conn, login string, id int) {
	CheckDoubleLogin(login, &usersLobbyWs)

	newPlayer := &player.Player{}
	newPlayer.SetLogin(login)
	newPlayer.SetID(id)

	usersLobbyWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS lobby Сессия: ")                          // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

	go NewLobbyUser(login, &usersLobbyWs)
	go SentOnlineUser(login, &usersLobbyWs)

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	for {
		var msg Message
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersLobbyWs, err)
			break
		}

		if msg.Event == "MapView" {
			var maps = lobby.GetMapList()
			for _, Map := range maps {
				var resp = Response{Event: msg.Event, Map: &Map}
				ws.WriteJSON(resp)
			}
		}

		if msg.Event == "GameView" {
			for _, game := range openGames {
				var resp = Response{Event: msg.Event, Game: game}
				ws.WriteJSON(resp)
			}
		}

		if msg.Event == "CreateLobbyGame" {
			user := usersLobbyWs[ws]
			game := lobby.CreateNewLobbyGame(msg.GameName, msg.MapID, user)

			openGames[game.ID] = &game
			//user.Set = game.Name todo адаптаировать под новую логику

			var resp = Response{Event: msg.Event, UserName: user.GetLogin(), Game: &game}
			ws.WriteJSON(resp)

			RefreshLobbyGames(usersLobbyWs[ws])
		}

		if msg.Event == "JoinToLobbyGame" {
			JoinToLobbyGame(msg, ws)
		}

		if msg.Event == "InitLobby" {
			var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].GetLogin()}
			ws.WriteJSON(resp)
		}

		if msg.Event == "Respawn" {
			user := usersLobbyWs[ws]

			for _, game := range openGames {
				for _, player := range game.Users {
					if player != nil {
						if user.GetLogin() == player.GetLogin() {
							var resp = Response{Event: msg.Event, Respawns: game.Respawns}
							ws.WriteJSON(resp)
						}
					}
				}
			}
		}

		if msg.Event == "Ready" {
			Ready(msg, ws)
		}

		if msg.Event == "StartNewGame" {
			StartNewGame(msg, ws)
		}

		if msg.Event == "DontEndGamesList" {
			games := lobby.GetDontEndGames(usersLobbyWs[ws].GetID())
			var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].GetLogin(), DontEndGames: games}
			ws.WriteJSON(resp)
		}

		if msg.Event == "Logout" {
			ws.Close()
		}
	}
}

func LobbyReposeSender() {
	for {
		resp := <-lobbyPipe
		mutex.Lock()
		for ws, client := range usersLobbyWs {
			if client.GetLogin() == resp.UserName {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					ws.Close()
					delete(usersLobbyWs, ws)
				}
			}
		}
		mutex.Unlock()
	}
}
