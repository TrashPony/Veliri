package chat

import (
	"github.com/gorilla/websocket"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*Clients) {
	for ws, client := range *usersWs {
		if client.Login == login {
			ws.Close()
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*Clients, err error) {
	if (*usersWs)[ws] != nil{
		delete(*usersWs, ws) // удаляем его из активных подключений
	}
}