package webSocket

import (
	"log"
	"websocket-master"
	"strconv"
	"../DB_info"
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

		if msg.Event == "MapView"{
			// запрашивает список доступных карт
			var maps = DB_info.MapList()
			var resp = LobbyResponse{"MapView", LoginWs(ws, &usersLobbyWs), "", maps, "", ""}
			LobbyPipe <- resp // Отправляет сообщение в тред
		}

		if msg.Event == "GameView"{
			// запрашивает список созданых игор
			var games []string = DB_info.OpenLobbyGameList()
			var resp = LobbyResponse{"GameView", LoginWs(ws, &usersLobbyWs), games[0], games[1], games[2], ""}
			LobbyPipe <- resp
		}

		if msg.Event == "DontEndGames"{
			// запрашивает списко незавершенных игор
			var gameName string = DB_info.DontEndGames(LoginWs(ws, &usersLobbyWs))
			var resp = LobbyResponse{"DontEndGames", LoginWs(ws, &usersLobbyWs), gameName, "", "", ""}
			LobbyPipe <- resp
		}

		if msg.Event == "ConnectGame"{
			// подключается к игре
			user2 , success := DB_info.ConnectGame(msg.GameName, LoginWs(ws, &usersLobbyWs))
			var resp = LobbyResponse{"GameView", LoginWs(ws, &usersLobbyWs), strconv.FormatBool(success), "", "" , user2}
			LobbyPipe <- resp
		}

		if msg.Event == "CreateNewGame"{
			// создает новую игру
			DB_info.CreateNewLobbyGame(msg.GameName, msg.MapName, LoginWs(ws, &usersLobbyWs))
			var resp = LobbyResponse{"CreateNewGame", LoginWs(ws, &usersLobbyWs), "", "", "", ""}
			LobbyPipe <- resp
		}

		if msg.Event == "StartNewGame"{

			// а вот этот метод должен после того как все подтвердили создавать новую игру в бд и перекидывать на новую страницу
			//lobby.StartNewGame(msg.MapName, msg.UserName)
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


