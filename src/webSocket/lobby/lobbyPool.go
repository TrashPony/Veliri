package lobby

import (
	"../../mechanics/db/get"
	"../../mechanics/db/insert"
	"../../mechanics/lobby"
	"../../mechanics/player"
	"../../mechanics/players"
	"../utils"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var mutex = &sync.Mutex{}

var usersLobbyWs = make(map[*websocket.Conn]*player.Player) // тут будут храниться наши подключения
var lobbyPipe = make(chan Response)
var openGames = make(map[int]*lobby.Game)

func AddNewUser(ws *websocket.Conn, login string, id int) {

	utils.CheckDoubleLogin(login, &usersLobbyWs)

	newPlayer, ok := players.Users.Get(id)

	if !ok {
		newPlayer = players.Users.Add(id, login)
	}

	usersLobbyWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS lobby Сессия: ") // просто смотрим новое подключение
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

	insert.UserIntoBase(id, 1) // пока есть только 1 база
	newPlayer.InBaseID = 1

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	for {
		var msg Message

		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil {          // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersLobbyWs, err)
			break
		}

		if msg.Event == "MapView" {
			var maps = get.InfoMapList()
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
			game := lobby.CreateNewLobbyGame(msg.GameName, msg.MapID, user, len(openGames))

			openGames[game.ID] = &game
			user.SetGameID(game.ID)

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
				for _, gamePlayer := range game.Users {
					if gamePlayer != nil {
						if user.GetLogin() == gamePlayer.GetLogin() {
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
			user := usersLobbyWs[ws]
			usersLobbyWs[ws].GetLogin()
			games := get.GetNotFinishedGames(user.GetID())

			var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].GetLogin(), DontEndGames: games}
			ws.WriteJSON(resp)
		}

		if msg.Event == "Logout" {
			ws.Close()
		}

		if msg.Event == "GetSquad" {
			var resp = Response{Event: "GetSquad", Squad: usersLobbyWs[ws].GetSquad()}
			ws.WriteJSON(resp)
		}
	}
}

func ReposeSender() {
	for {
		resp := <-lobbyPipe
		mutex.Lock()
		for ws, client := range usersLobbyWs {
			if client.GetLogin() == resp.UserName {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Fatal(err)
					DelConn(ws, &usersLobbyWs, err)
				}
			}
		}
		mutex.Unlock()
	}
}
