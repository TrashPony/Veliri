package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/gorilla/websocket"
)

func HandlerParse(user *player.Player, ws *websocket.Conn, coor *coordinate.Coordinate) {
	if coor.Handler == "base" {
		IntoToBase(user, coor.ToBaseID, ws)
	}

	if coor.Handler == "sector" {
		ChangeSector(user, coor.ToMapID, coor.ToQ, coor.ToR, ws)
	}

	if !user.Bot {
		go update.Squad(user.GetSquad(), true)
	}
}

func ChangeSector(user *player.Player, mapID, q, r int, ws *websocket.Conn) {
	stopMove(user, true)

	go SendMessage(Message{Event: "changeSector", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
	DisconnectUser(user, ws, true) // если только сообщение то можно не горутиной

	user.GetSquad().MapID = mapID
	user.GetSquad().Q = q
	user.GetSquad().R = r

	user.GetSquad().GlobalX = 0
	user.GetSquad().GlobalY = 0

	if user.Bot {
		LoadGame(ws, Message{Event: "InitGame"})
	}
}

func IntoToBase(user *player.Player, baseID int, ws *websocket.Conn) {
	if !user.Bot {
		bases.UserIntoBase(user.GetID(), baseID)
	}

	go SendMessage(Message{Event: "IntoToBase", IDUserSend: user.GetID(), Bot: user.Bot})
	go DisconnectUser(user, ws, true)

	user.InBaseID = baseID

	if user.GetSquad() != nil {
		user.GetSquad().GlobalX = 0
		user.GetSquad().GlobalY = 0
	}
}
