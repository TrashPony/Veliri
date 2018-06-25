package inventory

import (
	"github.com/gorilla/websocket"
	"../../inventory"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*inventory.User) {
	for ws, client := range *usersWs {
		if client.Name == login {
			ws.Close()
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*inventory.User, err error) {
	if (*usersWs)[ws] != nil{
		delete(*usersWs, ws) // удаляем его из активных подключений
	}
}