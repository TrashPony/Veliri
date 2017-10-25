package webSocket

import (
	"websocket-master"
	"net/http"
	"log"
	"./lobby"
	"./field"
)

var upgrader = websocket.Upgrader{} // методами приема обычного HTTP-соединения и обновления его до WebSocket


func ReadSocket(login string, id int, w http.ResponseWriter, r *http.Request, pool string)  {

	ws, err := upgrader.Upgrade(w, r, nil) // запрос GET для перехода на протокол websocket
	if err != nil {
		log.Fatal(err)
	}

	if pool == "/wsLobby" {
		lobby.AddNewUser(ws, login, id)
	}
	if pool == "/wsField" {
		field.AddNewUser(ws, login, id)
	}
}