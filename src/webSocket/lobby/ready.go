package lobby

import (
	"github.com/gorilla/websocket"
	"strconv"
)

func Ready(msg Message, ws *websocket.Conn) {

	game, ok := openGames[usersLobbyWs[ws].Game]
	user := usersLobbyWs[ws]

	if ok && user.SetReady() {

		respawn, err := game.SetRespawnUser(user, msg.RespawnID)

		if err == nil {
			game.UserReady(user, respawn)

			for _, gameUser := range game.Users {
				resp := Response{Event: msg.Event, UserName: gameUser.Name, GameUser: user.Name, Ready: strconv.FormatBool(user.Ready), Respawn: user.Respawn}
				lobbyPipe <- resp
			}
		} else {
			resp := Response{Event: msg.Event, UserName: user.Name, GameUser: user.Name, Error: err.Error()}
			ws.WriteJSON(resp)
		}
	} else {
		if !user.SetReady() {
			resp := Response{Event: msg.Event, UserName: user.Name, NameGame: msg.GameName, Error: "не выбран или не настроен отряд"}
			ws.WriteJSON(resp)
		} else {
			resp := Response{Event: msg.Event, UserName: user.Name, NameGame: msg.GameName, Error: "Game not find"}
			ws.WriteJSON(resp)
		}
	}
}
