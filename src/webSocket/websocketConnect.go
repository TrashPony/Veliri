package webSocket

import (
	"./chat"
	"./field"
	"./inventory"
	"./lobby"
	"./mapEditor"
	"./market"
	"./storage"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // методами приема обычного HTTP-соединения и обновления его до WebSocket

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

	if pool == "/wsMapEditor" {
		mapEditor.AddNewUser(ws, login, id)
	}

	if pool == "/wsMarket" {
		market.AddNewUser(ws, login, id)
	}

	if pool == "/wsStorage" {
		storage.AddNewUser(ws, login, id)
	}
}
