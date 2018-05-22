package field

import (
	"github.com/gorilla/websocket"
	"log"
	"../../mechanics/player"
)

func CheckDoubleLogin(login string, usersWs *map[*websocket.Conn]*player.Player) {
	for ws, client := range *usersWs {
		if client.GetLogin() == login {
			ws.Close()
			println(login + " Уже был в соеденениях")
		}
	}
}

func DelConn(ws *websocket.Conn, usersWs *map[*websocket.Conn]*player.Player, err error) {
	log.Printf("error: %v", err)
	delete(*usersWs, ws) // удаляем его из активных подключений
}
