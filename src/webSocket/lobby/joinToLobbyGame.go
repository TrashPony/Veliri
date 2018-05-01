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

		//все кто в лоби получают сообщение о том что подключился новый игрок
		for user := range game.Users {
			if user != usersLobbyWs[ws].Login {
				resp = Response{Event: "NewUser", UserName: user, NewUser: usersLobbyWs[ws].Login}
				lobbyPipe <- resp
			}
		}

		// игрок получает список всех игроков в лоби и их респауны
		for user, ready := range game.Users {
			if user != usersLobbyWs[ws].Login {
				var respown lobby.Respawn
				if ready {
					for _, respawn := range game.Respawns {
						if respawn.UserName == user {
							respown = *respawn
						}
					}
				}
				resp = Response{Event: "JoinToLobby", UserName: usersLobbyWs[ws].Login, GameUser: user, Ready: strconv.FormatBool(ready), Respawn: &respown}
				lobbyPipe <- resp
			}
		}
		RefreshLobbyGames(usersLobbyWs[ws].Login) // обновляет кол-во игроков и их характиристики в неигровом лоби
	}
}
