package lobby

import (
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"../../lobby"
)

func Ready(msg Message, ws *websocket.Conn)  {
	var resp Response
	game, errGetName := lobby.GetGame(msg.GameName)
	if errGetName != nil {
		log.Panic(errGetName)
	}

	respName, err := lobby.SetRespawnUser(msg.GameName, usersLobbyWs[ws].Login, msg.Respawn)

	if err == nil {
		lobby.UserReady(msg.GameName, usersLobbyWs[ws].Login)

		for user := range game.Users {
			resp = Response{Event: msg.Event, UserName: user, GameUser: usersLobbyWs[ws].Login, Ready: strconv.FormatBool(game.Users[usersLobbyWs[ws].Login]), RespawnName: respName}
			lobbyPipe <- resp
		}

	} else {
		resp = Response{Event: msg.Event, UserName: usersLobbyWs[ws].Login, GameUser: usersLobbyWs[ws].Login, Error: err.Error()}
		lobbyPipe <- resp
	}
}
