package webSocket

import (
	"./field"
	"./lobby"
	"./chat"
	"./inventory"

	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{} // методами приема обычного HTTP-соединения и обновления его до WebSocket

func ReadSocket(login string, id int, w http.ResponseWriter, r *http.Request, pool string) {

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

	if pool == "/wsChat" {
		chat.AddNewUser(ws, login, id)
	}

	if pool == "/wsInventory" {
		inventory.AddNewUser(ws, login, id)
	}
}
