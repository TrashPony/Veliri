package lobby

import (
	"github.com/gorilla/websocket"
	"errors"
	"../../mechanics/lobby"
	"../../mechanics/db/insert"
)

func StartNewGame(msg Message, ws *websocket.Conn) {

	game, ok := openGames[usersLobbyWs[ws].GetGameID()]

	if ok {
		// список игроков которым надо разослать данные взятые из обьекта игры
		if len(game.Users) > 1 {

			allReady := CheckReady(game)

			if allReady {

				id, success := insert.StartNewGame(game)

				if success {

					delete(openGames, game.ID) // удаляем обьект игры из лоби ¯\_(ツ)_/¯
					for _, user := range game.Users {
						var resp = Response{Event: msg.Event, UserName: user.GetLogin(), IdGame: id}
						lobbyPipe <- resp
					}

				} else {
					var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].GetLogin(), Error: errors.New("error ad to DB").Error()}
					ws.WriteJSON(resp)
				}
			} else {
				var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].GetLogin(), Error: errors.New("PlayerNotReady").Error()}
				ws.WriteJSON(resp)
			}
		} else {
			var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].GetLogin(), Error: errors.New("players < 2").Error()}
			ws.WriteJSON(resp)
		}
	} else {
		var resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].GetLogin(), Error: errors.New("game not found").Error()}
		ws.WriteJSON(resp)
	}
}

func CheckReady(game *lobby.Game) bool  {
	for _, user := range game.Users {
		if !user.GetLobbyReady() {
			return false
			break
		}
	}
	return true
}
