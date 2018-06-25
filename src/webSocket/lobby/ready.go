package lobby

import (
	"github.com/gorilla/websocket"
	"strconv"
)

func Ready(msg Message, ws *websocket.Conn) {

	game, ok := openGames[usersLobbyWs[ws].GetGameID()]
	user := usersLobbyWs[ws]

	if ok && user.GetReady() {

		respawn, err := game.SetRespawnUser(user, msg.RespawnID)

		if err == nil {
			game.UserReady(user, respawn)

			for _, gameUser := range game.Users {
				resp := Response{Event: msg.Event, UserName: gameUser.GetLogin(), GameUser: user.GetLogin(), Ready: strconv.FormatBool(user.GetReady()), Respawn: user.GetRespawn()}
				lobbyPipe <- resp
			}
		} else {
			resp := Response{Event: msg.Event, UserName: user.GetLogin(), GameUser: user.GetLogin(), Error: err.Error()}
			ws.WriteJSON(resp)
		}
	} else {
		if !user.GetReady() {
			resp := Response{Event: msg.Event, UserName: user.GetLogin(), NameGame: msg.GameName, Error: "не выбран или не настроен отряд"}
			ws.WriteJSON(resp)
		} else {
			resp := Response{Event: msg.Event, UserName: user.GetLogin(), NameGame: msg.GameName, Error: "Game not find"}
			ws.WriteJSON(resp)
		}
	}
}
