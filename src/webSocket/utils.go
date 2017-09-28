package webSocket

import (
	"websocket-master"
	"log"
	"net/http"
	"strconv"
)
var upgrader = websocket.Upgrader{} // методами приема обычного HTTP-соединения и обновления его до WebSocket
var usersLobbyWs = make(map[Clients]bool) // тут будут храниться наши подключения
var usersFieldWs = make(map[Clients]bool) // тут будут храниться наши подключения
var LobbyPipe = make(chan LobbyResponse)
var FieldPipe = make(chan FieldResponse)

func ReadSocket(login string, id int, w http.ResponseWriter, r *http.Request, pool string)  {

	ws, err := upgrader.Upgrade(w, r, nil) // запрос GET для перехода на протокол websocket
	if err != nil {
		log.Fatal(err)
	}

	if pool == "/wsLobby" {
		usersLobbyWs[Clients{ws, login, id}] = true // Регистрируем нового Клиента
		print("WS lobby Сессия: ") // просто смотрим новое подключение
		print(ws)
		println(" login: " + login + " id: " + strconv.Itoa(id))
		defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
		LobbyReader(ws)
	}
	if pool == "/wsField" {
		usersFieldWs[Clients{ws, login, id}] = true // Регистрируем нового Клиента
		print("WS field Сессия: ") // просто смотрим новое подключение
		print(ws)
		println(" login: " + login + " id: " + strconv.Itoa(id))
		defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
		FieldReader(ws)
	}

}

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