package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
)

func HandlerParse(user *player.Player, ws *websocket.Conn, coor *coordinate.Coordinate) {
	if coor.Handler == "base" {
		intoToBase(user, coor.ToBaseID, ws)
	}

	if coor.Handler == "sector" {
		changeSector(user, coor.ToMapID, coor.ToQ, coor.ToR, ws)
	}

	if !user.Bot {
		go update.Squad(user.GetSquad(), true)
	}
}

func changeSector(user *player.Player, mapID, q, r int, ws *websocket.Conn) {
	// TODO
	//if user.GetSquad().MoveChecker {
	//	stopMove(user, true)
	//}
	//
	//globalPipe <- Message{Event: "changeSector", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot}
	//DisconnectUser(user)
	//
	//user.GetSquad().MapID = mapID
	//
	//user.GetSquad().Q = q
	//user.GetSquad().R = r
	//
	//user.GetSquad().GlobalX = 0
	//user.GetSquad().GlobalY = 0
}

func intoToBase(user *player.Player, baseID int, ws *websocket.Conn) {
	if !user.Bot {
		bases.UserIntoBase(user.GetID(), baseID)
	}

	go sendMessage(Message{Event: "IntoToBase", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot})
	go DisconnectUser(user, ws)

	user.InBaseID = baseID

	user.GetSquad().GlobalX = 0
	user.GetSquad().GlobalY = 0
}
