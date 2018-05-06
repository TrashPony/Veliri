package lobby

import (
	"github.com/gorilla/websocket"
	"errors"
	"../../lobby"
)

func StartNewGame(msg Message, ws *websocket.Conn) {

	game, ok := openGames[usersLobbyWs[ws].Game]

	if ok {
		// список игроков которым надо разослать данные взятые из обьекта игры
		if len(game.Users) > 1 {

			allReady := CheckReady(game)

			if allReady {

				id, success := lobby.StartNewGame(game)

				if success {

					delete(openGames, game.Name) // удаляем обьект игры из лоби ¯\_(ツ)_/¯
					for _, user := range game.Users {
						var resp = Response{Event: msg.Event, UserName: user.Name, IdGame: id}
						lobbyPipe <- resp
					}

				} else {
					var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Name, Error: errors.New("error ad to DB").Error()}
					ws.WriteJSON(resp)
				}
			} else {
				var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Name, Error: errors.New("PlayerNotReady").Error()}
				ws.WriteJSON(resp)
			}
		} else {
			var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Name, Error: errors.New("Players < 2").Error()}
			ws.WriteJSON(resp)
		}
	} else {
		var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Name, Error: errors.New("game not found").Error()}
		ws.WriteJSON(resp)
	}
}

func CheckReady(game *lobby.LobbyGames) bool  {
	for _, user := range game.Users {
		if !user.Ready {
			return false
			break
		}
	}
	return true
}
