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

		if msg.Event == "DontEndGamesList"{
			// запрашивает списко незавершенных игор
			gameName, ids := DB_info.DontEndGames(LoginWs(ws, &usersLobbyWs))
			var resp = LobbyResponse{"DontEndGamesList", LoginWs(ws, &usersLobbyWs), gameName, ids, "", ""}
			LobbyPipe <- resp
		}

		if msg.Event == "JoinToLobbyGame"{
			//тут написана хуйня с расчетом на то что в будущем будет возможна игра больше 2х игроков одновременно
			playerList := DB_info.GetUserList(msg.GameName)
			creator := DB_info.JoinToLobbyGame(msg.GameName, LoginWs(ws, &usersLobbyWs))
			//все кто в лоби получают сообщение о том что подключился новйы игрок
			// for blabla список юзеров каждому отправить месагу
			var resp = LobbyResponse{"JoinToLobbyGame", creator, "", "", LoginWs(ws, &usersLobbyWs), ""}
			LobbyPipe <- resp

			// игрок получает список всех игроков в лоби (нет, пока только создателя так как 1 на 1)
			for i := 0; i < len(playerList); i++ {
				resp = LobbyResponse{"Joiner", LoginWs(ws, &usersLobbyWs), "", "", "", playerList[0]} //- костылька
				LobbyPipe <- resp
			}
		}

		if msg.Event == "CreateLobbyGame"{
			DB_info.CreateNewLobbyGame(msg.GameName, msg.MapName, LoginWs(ws, &usersLobbyWs))
			var resp = LobbyResponse{"CreateLobbyGame", LoginWs(ws, &usersLobbyWs), "", "", "", ""}
			LobbyPipe <- resp
		}

		if msg.Event == "StartNewGame"{
			id, success := DB_info.StartNewGame(msg.GameName)
			if success {
				//тут написана хуйня с расчетом на то что в будущем будет возможна игра больше 2х игроков одновременно
				playerList := DB_info.GetUserList(msg.GameName) // список игроков которым надо разослать данные взятые из обьекта игры
				DB_info.DelLobbyGame(LoginWs(ws, &usersLobbyWs)) // удаляем обьект игры из лоби, ищем его по имени создателя ¯\_(ツ)_/¯

				for i := 0; i < len(playerList); i++ {
					var resp = LobbyResponse{"StartNewGame", playerList[i], strconv.FormatBool(success), id, "", ""}
					LobbyPipe <- resp
				}
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


