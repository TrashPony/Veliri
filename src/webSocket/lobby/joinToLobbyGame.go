package lobby

import (
	"github.com/gorilla/websocket"
)

func JoinToLobbyGame(msg Message, ws *websocket.Conn) {

	game, ok := openGames[msg.GameName]

	if ok {
		err := game.JoinToLobbyGame(usersLobbyWs[ws])
		if err != nil {

			resp := Response{Event: "initLobbyGame", UserName: usersLobbyWs[ws].Name, NameGame: game.Name, Error: err.Error()}
			ws.WriteJSON(resp)

		} else {
			// добавляем в игрока имя игры
			usersLobbyWs[ws].Game = game.Name
			// игрок инициализирует лобби меню на клиенте и создает там игроков
			resp := Response{Event: "initLobbyGame", UserName: usersLobbyWs[ws].Name, User: usersLobbyWs[ws], GameUsers: game.Users}
			ws.WriteJSON(resp)

			//все кто в лоби получают сообщение о том что подключился новый игрок
			for _, user := range game.Users {
				resp = Response{Event: "NewUser", UserName: user.Name, User: usersLobbyWs[ws]}
				lobbyPipe <- resp
			}

			// обновляет кол-во игроков и их характиристики в неигровом лоби
			RefreshLobbyGames(usersLobbyWs[ws])
		}
	} else {
		resp := Response{Event: msg.Event, UserName: usersLobbyWs[ws].Name, NameGame: msg.GameName, Error: "Game not find"}
		ws.WriteJSON(resp)
	}
}
