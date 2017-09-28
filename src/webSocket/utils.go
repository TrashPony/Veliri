package webSocket

import (
	"websocket-master"
	"log"
)


func LoginWs(ws *websocket.Conn, usersWs *map[Clients]bool) (string)  {
	for client := range *usersWs { // ходим по всем подключениям
		if(client.ws == ws) {
			return client.login
		}
	}
	return ""
}

func DelConn(ws *websocket.Conn, usersWs *map[Clients]bool, err error)  {
	log.Printf("error: %v", err)
	for client := range *usersWs { // ходим по всем подключениям
		if(client.ws == ws) { // находим подключение с ошибкой
			delete(*usersWs, client) // удаляем его из активных подключений
			break
		}
	}
}

type Clients struct { // структура описывающая клиента ws соеденение
	ws *websocket.Conn
	login string
	id int
}