package lobby

import (
	"github.com/gorilla/websocket"
	"../../lobby"
)

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*lobby.User, err error) {
	if (*usersWs)[ws] != nil {
		user := (*usersWs)[ws]


		delete(*usersWs, ws)             // удаляем его из активных подключений
		DelLobbyUser(user.Name, usersWs) // удаляем из общего списка игроков

		game := openGames[user.Game] // получаем игру в которой он был

		if game != nil { // если такая игра есть оповещаем других игроков о том что он вышел
			DelUserInLobby(game, user.Name)

			if user.Name == game.Creator.Name {
				// если это создатель говорим другим игрокам выйти в общее лоби
				DisconnectLobby(game.Users)
				delete(openGames, game.Name)
			} else {
				game.RemoveUser(user)
			}
		}

		RefreshLobbyGames(user)
	}
}

func DelLobbyUser(login string, usersWs *map[*websocket.Conn]*lobby.User) {
	for _, client := range *usersWs {
		var resp = Response{Event: "DelLobbyUser", UserName: client.Name, GameUser: login}
		lobbyPipe <- resp
	}
}

func DelUserInLobby(game *lobby.LobbyGames, delLogin string) {
	for _, user := range game.Users {
		if user != nil {
			var message = Response{Event: "DelUser", UserName: user.Name, GameUser: delLogin}
			lobbyPipe <- message
		}
	}
}

func DisconnectLobby(users []*lobby.User) {
	for _, user := range users {
		if user != nil {
			var refresh = Response{Event: "DisconnectLobby", UserName: user.Name}
			lobbyPipe <- refresh
		}
	}
}
