package chat

import (
	"../../mechanics/player"
	"github.com/gorilla/websocket"
)

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*player.Player, err error) {
	if (*usersWs)[ws] != nil {
		delete(*usersWs, ws) // удаляем его из активных подключений
	}
}
