package webSocket

import (
	"websocket-master"
	"net/http"
	"log"
	"strconv"
	"../game/initGame"
)

var upgrader = websocket.Upgrader{} // методами приема обычного HTTP-соединения и обновления его до WebSocket
var usersLobbyWs = make(map[*websocket.Conn]*Clients) // тут будут храниться наши подключения
var usersFieldWs = make(map[*websocket.Conn]*Clients) // тут будут храниться наши подключения
var LobbyPipe = make(chan LobbyResponse)
var FieldPipe = make(chan FieldResponse)

func ReadSocket(login string, id int, w http.ResponseWriter, r *http.Request, pool string)  {

	ws, err := upgrader.Upgrade(w, r, nil) // запрос GET для перехода на протокол websocket
	if err != nil {
		log.Fatal(err)
	}

	if pool == "/wsLobby" {
		CheckDoubleLogin(login, &usersLobbyWs)

		usersLobbyWs[ws] = &Clients{login:login, id:id} // Регистрируем нового Клиента //TODO: map[Clients]bool --> map[ws]Clients
		print("WS lobby Сессия: ") // просто смотрим новое подключение
		print(ws)
		println(" login: " + login + " id: " + strconv.Itoa(id))
		defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
		LobbyReader(ws)
	}
	if pool == "/wsField" {
		CheckDoubleLogin(login, &usersFieldWs)
		usersFieldWs[ws] = &Clients{login:login, id:id} // Регистрируем нового Клиента
		print("WS field Сессия: ") // просто смотрим новое подключение
		print(ws)
		println(" login: " + login + " id: " + strconv.Itoa(id))
		defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
		FieldReader(ws)
	}
}

type Clients struct { // структура описывающая клиента ws соеденение
	login string
	id int
	permittedCoordinates []initGame.Coordinate
	Units []initGame.Unit
	Respawn initGame.Respawn
	CreateZone []initGame.Coordinate
}
