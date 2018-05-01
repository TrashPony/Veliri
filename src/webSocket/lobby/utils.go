package lobby

import (
	"../../lobby"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*Clients) {
	for ws, client := range *usersWs {
		if client.Login == login {
			ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func NewLobbyUser(login string, usersWs *map[*websocket.Conn]*Clients) {
	for _, client := range *usersWs {
		var resp = Response{Event: "NewLobbyUser", UserName: client.Login, GameUser: login}
		lobbyPipe <- resp
	}
}

func SentOnlineUser(login string, usersWs *map[*websocket.Conn]*Clients) {
	for _, client := range *usersWs {
		if login != client.Login {
			var resp = Response{Event: "NewLobbyUser", UserName: login, GameUser: client.Login}
			lobbyPipe <- resp
		}
	}
}

func DelLobbyUser(login string, usersWs *map[*websocket.Conn]*Clients) {
	for _, client := range *usersWs {
		var resp = Response{Event: "DelLobbyUser", UserName: client.Login, GameUser: login}
		lobbyPipe <- resp
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*Clients, err error) {
	log.Printf("error: %v", err)
	if (*usersWs)[ws] != nil {
		login := (*usersWs)[ws].Login

		delete(*usersWs, ws)         // удаляем его из активных подключений
		DelLobbyUser(login, usersWs) // удаляем из общего списка игроков

		game := lobby.DisconnectLobbyGame(login) // получаем игру в которой он был

		if game != nil { // если такая игра есть оповещаем других игроков о том что он вышел
			DelUserInLobby(game, login)
		}

		delGame := lobby.DelLobbyGame(login)

		if delGame != nil {
			DiconnectLobby(delGame.Users)
			RefreshLobbyGames(login)
		}
	}
}

func DelUserInLobby(game *lobby.LobbyGames, delLogin string) {
	for user := range game.Users {
		var message = Response{Event: "DelUser", UserName: user, GameUser: delLogin}
		lobbyPipe <- message
	}
}

func DiconnectLobby(users map[string]bool) {
	for client := range users {
		var refresh = Response{Event: "DisconnectLobby", UserName: client}
		lobbyPipe <- refresh
	}
}

func RefreshLobbyGames(login string) {
	games := lobby.GetLobbyGames()
	for _, client := range usersLobbyWs {
		if client.Login != login {
			var refresh = Response{Event: "GameRefresh", UserName: client.Login}
			lobbyPipe <- refresh
			for _, game := range games {
				var resp = Response{Event: "GameView", UserName: client.Login, NameGame: game.Name, Map: game.Map, Creator: game.Creator,
					Players: strconv.Itoa(len(game.Users)), NumOfPlayers: strconv.Itoa(len(game.Respawns))}
				lobbyPipe <- resp
			}
		}
	}
}
