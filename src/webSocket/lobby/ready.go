package lobby

import (
	"github.com/gorilla/websocket"
)

func Ready(msg Message, ws *websocket.Conn) {

	game, ok := openGames[usersLobbyWs[ws].GetGameID()]
	user := usersLobbyWs[ws]

	if ok && user.GetSquad().MatherShip.Body != nil {

		respawn, err := game.SetRespawnUser(user, msg.RespawnID)

		if err == nil || user.Ready {
			game.UserReady(user, respawn)

			for _, gameUser := range game.Users {
				resp := Response{Event: msg.Event, UserName: gameUser.GetLogin(), GameUser: user.GetLogin(), Ready: user.GetReady(), Respawn: user.GetRespawn()}
				lobbyPipe <- resp
			}
		} else {
			resp := Response{Event: msg.Event, UserName: user.GetLogin(), GameUser: user.GetLogin(), Error: err.Error()}
			ws.WriteJSON(resp)
		}
	} else {
		if user.GetSquad().MatherShip.Body == nil {
			resp := Response{Event: msg.Event, UserName: user.GetLogin(), NameGame: msg.GameName, Error: "не выбран или не настроен отряд"}
			ws.WriteJSON(resp)
		} else {
			resp := Response{Event: msg.Event, UserName: user.GetLogin(), NameGame: msg.GameName, Error: "Game not find"}
			ws.WriteJSON(resp)
		}
	}
}
