package webSocket

import (
	"websocket-master"
	"log"
	"../DB_info"
	"strconv"
)


func LoginWs(ws *websocket.Conn, usersWs *map[Clients]bool) (string)  {
	for client := range *usersWs { // ходим по всем подключениям
		if(client.ws == ws) {
			return client.login
		}
	}
	return ""
}

func IdWs(ws *websocket.Conn, usersWs *map[Clients]bool) (int)  {
	for client := range *usersWs { // ходим по всем подключениям
		if(client.ws == ws) {
			return client.id
		}
	}
	return 0
}

func CheckDoubleLogin(login string, usersWs *map[Clients]bool)  {
	for client := range *usersWs {
		if client.login == login {
			client.ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[Clients]bool, err error)  {
	log.Printf("error: %v", err)
	for client := range *usersWs { // ходим по всем подключениям
		if(client.ws == ws) { // находим подключение с ошибкой
			delGame, users := DB_info.DelLobbyGame(client.login)
			diconnect, nameGame := DB_info.DisconnectLobbyGame(client.login)
			if delGame || diconnect {
				if delGame {
					DiconnectLobby(users)
				}
				if diconnect {
					RefreshUsersList(nameGame)
				}
				RefreshLobbyGames(ws)
			}
			delete(*usersWs, client) // удаляем его из активных подключений
			break
		}
	}
}

type Clients struct { // структура описывающая клиента ws соеденение
	ws *websocket.Conn
	login string
	id int
}

func RefreshUsersList(nameGame string)  {
	games := DB_info.GetLobbyGames()
	for _, game := range games {
		if game.Name == nameGame {
			for player := range game.Users {
				var refresh= LobbyResponse{Event: "DelUser", UserName: player}
				LobbyPipe <- refresh
				for players, ready := range game.Users {
					if player != players {
						var refresh= LobbyResponse{Event: "UserRefresh", UserName: player, GameUser: players, Ready: strconv.FormatBool(ready)}
						LobbyPipe <- refresh
					}
				}
			}
			break
		}
	}
}

func DiconnectLobby(users map[string]bool)  {
	for client := range users {
		var refresh = LobbyResponse{Event: "DisconnectLobby",  UserName: client}
		LobbyPipe <- refresh
	}
}
func RefreshLobbyGames(ws *websocket.Conn)  {
	games := DB_info.GetLobbyGames()
	for client := range usersLobbyWs { // TODO: сильно затратная операция, над сделать что бы отсылалась только новая игра а не обновляляся весь список заного
		if client.login != LoginWs(ws, &usersLobbyWs) {
			var refresh = LobbyResponse{Event: "GameRefresh",  UserName: client.login}
			LobbyPipe <- refresh
			for _, game := range games {
				var resp = LobbyResponse{Event: "GameView", UserName: client.login, NameGame: game.Name, NameMap: game.Map, Creator: game.Creator,
					Players: strconv.Itoa(len(game.Users)), NumOfPlayers: strconv.Itoa(len(game.Respawns))}
				LobbyPipe <- resp
			}
		}
	}
}