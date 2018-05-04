package lobby

import (
	"../../lobby"
	"github.com/gorilla/websocket"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*lobby.User) {
	for ws, client := range *usersWs {
		if client.Name == login {
			ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func NewLobbyUser(login string, usersWs *map[*websocket.Conn]*lobby.User) {
	for _, client := range *usersWs {
		var resp = Response{Event: "NewLobbyUser", UserName: client.Name, GameUser: login}
		lobbyPipe <- resp
	}
}

func SentOnlineUser(login string, usersWs *map[*websocket.Conn]*lobby.User) {
	for _, client := range *usersWs {
		if login != client.Name {
			var resp = Response{Event: "NewLobbyUser", UserName: login, GameUser: client.Name}
			lobbyPipe <- resp
		}
	}
}

func RefreshLobbyGames(user *lobby.User) {
	for _, client := range usersLobbyWs {
		if client.Name != user.Name {

			var refresh = Response{Event: "GameRefresh", UserName: client.Name}
			lobbyPipe <- refresh

			for _, game := range openGames {

				var resp = Response{Event: "GameView", UserName: client.Name, Game:game}
				lobbyPipe <- resp

			}
		}
	}
}
