package webSocket

import (
	"github.com/TrashPony/Veliri/src/webSocket/editors/dialogEditor"
	"github.com/TrashPony/Veliri/src/webSocket/editors/mapEditor"
	"github.com/TrashPony/Veliri/src/webSocket/global"
	"github.com/TrashPony/Veliri/src/webSocket/inventory"
	"github.com/TrashPony/Veliri/src/webSocket/lobby"
	"github.com/TrashPony/Veliri/src/webSocket/market"
	"github.com/TrashPony/Veliri/src/webSocket/other"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ReadSocket(login string, id int, w http.ResponseWriter, r *http.Request, pool string) {

	ws, err := upgrader.Upgrade(w, r, nil) // запрос GET для перехода на протокол websocket
	if err != nil {
		log.Fatal(err)
	}

	ws.SetReadLimit(10485760)

	// TODO функция AddNewUser везде по сути одинаковая, возможно стоить вынести реализацию из всех соедений в общую функцию
	if pool == "/wsLobby" {
		lobby.AddNewUser(ws, login, id)
	}

	if pool == "/wsChat" {
		other.AddNewUser(ws, login, id)
	}

	if pool == "/wsInventory" {
		inventory.AddNewUser(ws, login, id)
	}

	if pool == "/wsMarket" {
		market.AddNewUser(ws, login, id)
	}

	if pool == "/wsGlobal" {
		global.AddNewUser(ws, login, id)
	}

	if pool == "/wsMapEditor" {
		mapEditor.AddNewUser(ws, login, id)
	}

	if pool == "/wsDialogEditor" {
		dialogEditor.AddNewUser(ws, login, id)
	}
}
