package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
	"time"
)

func HandlerParse(user *player.Player, ws *websocket.Conn, coor *coordinate.Coordinate) {
	if coor.Handler == "base" {
		IntoToBase(user, coor.ToBaseID, ws)
	}

	if coor.Handler == "sector" {
		ChangeSector(user, coor.ToMapID, ws, coor)
	}

	if !user.Bot {
		go update.Squad(user.GetSquad(), true)
	}
}

func ChangeSector(user *player.Player, mapID int, ws *websocket.Conn, coor *coordinate.Coordinate) {

	var toPosition *coordinate.Coordinate
	for {

		users, rLock := globalGame.Clients.GetAll()
		toPosition = globalGame.CheckHandlerCoordinate(coor, users)
		rLock.Unlock()

		// ждем пока не получит свободную позицию
		if toPosition != nil {
			break
		}

		time.Sleep(300 * time.Millisecond)
	}

	stopMove(user, true)

	go SendMessage(Message{Event: "changeSector", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
	DisconnectUser(user, ws, true) // если только сообщение то можно не горутиной

	user.GetSquad().MapID = mapID
	user.GetSquad().Q = toPosition.Q
	user.GetSquad().R = toPosition.R

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
