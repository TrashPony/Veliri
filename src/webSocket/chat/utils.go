package chat

import (
	"github.com/gorilla/websocket"
	"../../mechanics/player"
)

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*player.Player, err error) {
	if (*usersWs)[ws] != nil{
		delete(*usersWs, ws) // удаляем его из активных подключений
	}
}