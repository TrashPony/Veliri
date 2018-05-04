package lobby

import (
	"github.com/gorilla/websocket"
	"strconv"
)

func Ready(msg Message, ws *websocket.Conn) {

	game, ok := openGames[usersLobbyWs[ws].Game]

	if ok {

		respawn, err := game.SetRespawnUser(usersLobbyWs[ws], msg.RespawnID)

		if err == nil {
			game.UserReady(usersLobbyWs[ws], respawn)

			for _, user := range game.Users {
				resp := Response{Event: msg.Event, UserName: user.Name, GameUser: usersLobbyWs[ws].Name, Ready: strconv.FormatBool(usersLobbyWs[ws].Ready), Respawn: usersLobbyWs[ws].Respawn}
				lobbyPipe <- resp
			}
		} else {
			resp := Response{Event: msg.Event, UserName: usersLobbyWs[ws].Name, GameUser: usersLobbyWs[ws].Name, Error: err.Error()}
			ws.WriteJSON(resp)
		}
	} else {
		resp := Response{Event: msg.Event, UserName: usersLobbyWs[ws].Name, NameGame: msg.GameName, Error: "Game not find"}
		ws.WriteJSON(resp)
	}
}
