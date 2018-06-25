package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/player"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*player.Player) {
	for ws, client := range *usersWs {
		if client.GetLogin() == login {
			ws.Close()
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*player.Player, err error) {
	if (*usersWs)[ws] != nil{
		delete(*usersWs, ws) // удаляем его из активных подключений
	}
}