package webSocket

import (
	"log"
	"websocket-master"
	"../DB_info"
	"strconv"
)

// пайп доп. читать в документации

func LobbyReader(ws *websocket.Conn)  {
	for {
		var msg LobbyMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersLobbyWs, err)
			break
		}

		if msg.Event == "MapView" {
			var maps= DB_info.GetMapList()
			for _, Map := range maps {
				var resp = LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), NameMap: Map.Name, NumOfPlayers: strconv.Itoa(Map.Respawns)}
				LobbyPipe <- resp // Отправляет сообщение в тред
			}
		}

		if msg.Event == "GameView" {
			games := DB_info.GetLobbyGames()
			for _, game := range games {
				var resp = LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), NameGame: game.Name, NameMap: game.Map, Creator: game.Creator}
				LobbyPipe <- resp
			}
		}

		if msg.Event == "DontEndGamesList" {
			games := DB_info.GetDontEndGames(LoginWs(ws, &usersLobbyWs))
			for _, game := range games {
				var resp = LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), NameGame: game.Name, IdGame: game.Id, PhaseGame: game.Phase, StepGame: game.Step, Ready: game.Ready}
				LobbyPipe <- resp
			}
		}

		if msg.Event == "JoinToLobbyGame" {
			var resp LobbyResponse
			playerList := DB_info.GetUserList(msg.GameName)
			success, err := DB_info.JoinToLobbyGame(msg.GameName, LoginWs(ws, &usersLobbyWs))
			//все кто в лоби получают сообщение о том что подключился новйы игрок
			resp = LobbyResponse{Event: "initLobbyGame", UserName: LoginWs(ws, &usersLobbyWs), NameGame:msg.GameName, Error:err}
			LobbyPipe <- resp

			if success {
				for user := range playerList {
					if user != LoginWs(ws, &usersLobbyWs) {
						resp = LobbyResponse{Event: "NewUser", UserName: user, NewUser: LoginWs(ws, &usersLobbyWs)}
						LobbyPipe <- resp
					}
				}

				// игрок получает список всех игроков в лоби
				for user, ready := range playerList {
					if user != LoginWs(ws, &usersLobbyWs) {
						resp = LobbyResponse{Event: "JoinToLobby", UserName: LoginWs(ws, &usersLobbyWs), GameUser: user, Ready: strconv.FormatBool(ready)}
						LobbyPipe <- resp
					}
				}
			}
		}

		if msg.Event == "CreateLobbyGame" {
			DB_info.CreateNewLobbyGame(msg.GameName, msg.MapName, LoginWs(ws, &usersLobbyWs))
			var resp= LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs)}
			LobbyPipe <- resp
		}

		if msg.Event == "Ready" {
			DB_info.UserReady(msg.GameName, LoginWs(ws, &usersLobbyWs))
			playerList := DB_info.GetUserList(msg.GameName)
			for user := range playerList {
				if user != LoginWs(ws, &usersLobbyWs) {
					var resp = LobbyResponse{Event: msg.Event, UserName: user, GameUser: LoginWs(ws, &usersLobbyWs), Ready: strconv.FormatBool(playerList[LoginWs(ws, &usersLobbyWs)])}
					LobbyPipe <- resp
				}
			}
		}

		if msg.Event == "StartNewGame" {
			playerList := DB_info.GetUserList(msg.GameName) // список игроков которым надо разослать данные взятые из обьекта игры
			if len(playerList) > 1 {
				var readyAll = true
				for _, ready := range playerList {
					if !ready {
						readyAll = false
						break
					}
				}
				if readyAll {
					id, success := DB_info.StartNewGame(msg.GameName)
					if success {
						DB_info.DelLobbyGame(LoginWs(ws, &usersLobbyWs)) // удаляем обьект игры из лоби, ищем его по имени создателя ¯\_(ツ)_/¯
						for user := range playerList {
							var resp = LobbyResponse{Event: msg.Event, UserName: user, IdGame: id}
							LobbyPipe <- resp
						}
					} else {
						var resp= LobbyResponse{Event: msg.Event, UserName: LoginWs(ws, &usersLobbyWs), Error: "error ad to DB"}
						LobbyPipe <- resp
					}
				} else {
					var resp = LobbyResponse{Event: msg.Event, UserName:LoginWs(ws, &usersLobbyWs),  Error: "PlayerNotReady"}
					LobbyPipe <- resp
				}
			} else {
				var resp = LobbyResponse{Event: msg.Event, UserName:LoginWs(ws, &usersLobbyWs),  Error: "Players < 2"}
				LobbyPipe <- resp
			}
		}
	}
}

func LobbyReposeSender() {
	for {
		resp := <-LobbyPipe
		for client := range usersLobbyWs {
			if client.login == resp.UserName {
				err := client.ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					DB_info.DelLobbyGame(client.login)
					client.ws.Close()
					delete(usersLobbyWs, client)
				}
			}
		}
	}
}


