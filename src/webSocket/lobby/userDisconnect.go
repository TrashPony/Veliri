package lobby

import (
	"github.com/gorilla/websocket"
	"../../mechanics/lobby"
	"../../mechanics/player"
)

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*player.Player, err error) {
	if (*usersWs)[ws] != nil {
		user := (*usersWs)[ws]


		delete(*usersWs, ws)             // удаляем его из активных подключений
		DelLobbyUser(user.GetLogin(), usersWs) // удаляем из общего списка игроков

		game := openGames[user.GetGameID()] // получаем игру в которой он был

		if game != nil { // если такая игра есть оповещаем других игроков о том что он вышел
			DelUserInLobby(game, user.GetLogin())

			if user.GetLogin() == game.Creator {
				// если это создатель говорим другим игрокам выйти в общее лоби
				DisconnectLobby(game.Users)
				delete(openGames, game.ID)
			} else {
				game.RemoveUser(user)
			}
		}

		RefreshLobbyGames(user)
	}
}

func DelLobbyUser(login string, usersWs *map[*websocket.Conn]*player.Player) {
	for _, client := range *usersWs {
		var resp = Response{Event: "DelLobbyUser", UserName: client.GetLogin(), GameUser: login}
		lobbyPipe <- resp
	}
}

func DelUserInLobby(game *lobby.Game, delLogin string) {
	for _, user := range game.Users {
		if user != nil {
			var message = Response{Event: "DelUser", UserName: user.GetLogin(), GameUser: delLogin}
			lobbyPipe <- message
		}
	}
}

func DisconnectLobby(users map[string]*player.Player) {
	for _, user := range users {
		if user != nil {
			var refresh = Response{Event: "DisconnectLobby", UserName: user.GetLogin()}
			lobbyPipe <- refresh
		}
	}
}
