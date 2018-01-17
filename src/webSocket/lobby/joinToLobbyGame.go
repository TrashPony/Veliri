package lobby

import (
	"log"
	"strconv"
	"../../lobby"
	"github.com/gorilla/websocket"
)

func JoinToLobbyGame(msg Message, ws *websocket.Conn)  {
	var resp Response
	game, errGetName := lobby.GetGame(msg.GameName)
	if errGetName != nil {
		log.Panic(errGetName)
	}
	err := lobby.JoinToLobbyGame(msg.GameName, usersLobbyWs[ws].Login)
	if err != nil {
		resp = Response{Event: "initLobbyGame", UserName: usersLobbyWs[ws].Login, NameGame: msg.GameName, Error: err.Error()}
		lobbyPipe <- resp
	} else {
		// игрок инициализирует лобби меню на клиенте
		resp = Response{Event: "initLobbyGame", UserName: usersLobbyWs[ws].Login, NameGame: msg.GameName}
		lobbyPipe <- resp

		//все кто в лоби получают сообщение о том что подключился новйы игрок
		for user := range game.Users {
			if user != usersLobbyWs[ws].Login {
				resp = Response{Event: "NewUser", UserName: user, NewUser: usersLobbyWs[ws].Login}
				lobbyPipe <- resp
			}
		}

		// игрок получает список всех игроков в лоби и их респауны
		for user, ready := range game.Users {
			if user != usersLobbyWs[ws].Login {
				var respown string
				if ready {
					for respawns := range game.Respawns {
						if game.Respawns[respawns] == user {
							respown = respawns.Name
						}
					}
				}
				resp = Response{Event: "JoinToLobby", UserName: usersLobbyWs[ws].Login, GameUser: user, Ready: strconv.FormatBool(ready), RespawnName: respown}
				lobbyPipe <- resp
			}
		}
		RefreshLobbyGames(usersLobbyWs[ws].Login) // обновляет кол-во игроков и их характиристики в неигровом лоби
	}
}
