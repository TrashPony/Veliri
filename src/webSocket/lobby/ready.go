package lobby

import (
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"../../lobby"
)

func Ready(msg Message, ws *websocket.Conn)  {

	game, errGetName := lobby.GetGame(msg.GameName)
	if errGetName != nil {
		log.Panic(errGetName)
	}

	resawn, err := lobby.SetRespawnUser(msg.GameName, usersLobbyWs[ws].Login, msg.RespawnID)

	if err == nil {
		lobby.UserReady(msg.GameName, usersLobbyWs[ws].Login)

		for user := range game.Users {
			resp := Response{Event: msg.Event, UserName: user, GameUser: usersLobbyWs[ws].Login, Ready: strconv.FormatBool(game.Users[usersLobbyWs[ws].Login]), Respawn: resawn}
			lobbyPipe <- resp
		}
	} else {
		resp := Response{Event: msg.Event, UserName: usersLobbyWs[ws].Login, GameUser: usersLobbyWs[ws].Login, Error: err.Error()}
		ws.WriteJSON(resp)
	}
}
