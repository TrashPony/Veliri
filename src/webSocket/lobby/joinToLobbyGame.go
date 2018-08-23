package lobby

import (
	"github.com/gorilla/websocket"
)

func JoinToLobbyGame(msg Message, ws *websocket.Conn) {

	game, ok := openGames[msg.GameID]

	if ok {
		err := game.JoinToLobbyGame(usersLobbyWs[ws])
		if err != nil {

			resp := Response{Event: "initLobbyGame", UserName: usersLobbyWs[ws].GetLogin(), NameGame: game.Name, Error: err.Error()}
			ws.WriteJSON(resp)

		} else {
			// добавляем в игрока имя игры
			usersLobbyWs[ws].SetGameID(game.ID)
			// игрок инициализирует лобби меню на клиенте и создает там игроков
			resp := Response{Event: "initLobbyGame", UserName: usersLobbyWs[ws].GetLogin(), User: usersLobbyWs[ws], GameUsers: game.Users}
			ws.WriteJSON(resp)

			//все кто в лоби получают сообщение о том что подключился новый игрок
			for _, user := range game.Users {
				resp = Response{Event: "NewUser", UserName: user.GetLogin(), GameUser: usersLobbyWs[ws].GetLogin(), Ready: usersLobbyWs[ws].GetLobbyReady()}
				lobbyPipe <- resp
			}

			// обновляет кол-во игроков и их характиристики в неигровом лоби
			RefreshLobbyGames(usersLobbyWs[ws])
		}
	} else {
		resp := Response{Event: msg.Event, UserName: usersLobbyWs[ws].GetLogin(), NameGame: msg.GameName, Error: "Game not find"}
		ws.WriteJSON(resp)
	}
}
