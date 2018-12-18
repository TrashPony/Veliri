package lobby

import (
	"../../mechanics/db/localGame"
	"../../mechanics/factories/maps"
	"../../mechanics/factories/players"
	"../../mechanics/lobby"
	"../../mechanics/player"
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

	if newPlayer.InBaseID == 0 { // если игрок находиться не на базе то говорим ему что он загружал глобальную игру
		ws.WriteJSON(Response{Event: "OutBase"})
		return
	}

	usersLobbyWs[ws] = newPlayer // Регистрируем нового Клиента

	print("WS lobby Сессия: ") // просто смотрим новое подключение
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик

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
			var allMaps = maps.Maps.GetAllMap()
			for _, Map := range allMaps {
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
			games := localGame.GetNotFinishedGames(user.GetID())

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

		if msg.Event == "OutBase" {
			err := lobby.OutBase(usersLobbyWs[ws])
			if err != nil {
				ws.WriteJSON(Response{Event: "Error", Error: err.Error()})
			} else {
				ws.WriteJSON(Response{Event: msg.Event})
			}
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
