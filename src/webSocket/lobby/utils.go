package lobby

import (
	"github.com/gorilla/websocket"
	"../../mechanics/player"
)

func NewLobbyUser(login string, usersWs *map[*websocket.Conn]*player.Player) {
	for _, client := range *usersWs {
		var resp = Response{Event: "NewLobbyUser", UserName: client.GetLogin() , GameUser: login}
		lobbyPipe <- resp
	}
}

func SentOnlineUser(login string, usersWs *map[*websocket.Conn]*player.Player) {
	for _, client := range *usersWs {
		if login != client.GetLogin()  {
			var resp = Response{Event: "NewLobbyUser", UserName: login, GameUser: client.GetLogin() }
			lobbyPipe <- resp
		}
	}
}

func RefreshLobbyGames(user *player.Player) {
	for _, client := range usersLobbyWs {
		if client.GetLogin()  != user.GetLogin()  {

			var refresh = Response{Event: "GameRefresh", UserName: client.GetLogin() }
			lobbyPipe <- refresh

			for _, game := range openGames {

				var resp = Response{Event: "GameView", UserName: client.GetLogin() , Game:game}
				lobbyPipe <- resp

			}
		}
	}
}
