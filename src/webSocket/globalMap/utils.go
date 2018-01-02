package globalMap

import (
	"github.com/gorilla/websocket"
	"log"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*Clients) {
	for ws, client := range *usersWs {
		if client.Login == login {
			ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*Clients, err error) {
	log.Printf("error: %v", err)
	if (*usersWs)[ws] != nil {
		delete(*usersWs, ws)         // удаляем его из активных подключений
	}
}


