package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func HandlerParse(user *player.Player, coor *coordinate.Coordinate) {
	if coor.Handler == "base" {
		intoToBase(user, coor.ToBaseID)
	}

	if coor.Handler == "sector" {
		println("dfd")
		changeSector(user, coor.ToMapID, coor.ToQ, coor.ToR)
	}

	if !user.Bot {
		go update.Squad(user.GetSquad(), true)
	}
}

func changeSector(user *player.Player, mapID, q, r int) {

	if user.GetSquad().MoveChecker {
		user.GetSquad().MoveChecker = false
		user.GetSquad().ActualPath = nil
		user.GetSquad().CurrentSpeed = 0
	}

	globalPipe <- Message{Event: "changeSector", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot}
	DisconnectUser(user)

	user.GetSquad().MapID = mapID

	user.GetSquad().Q = q
	user.GetSquad().R = r

	user.GetSquad().GlobalX = 0
	user.GetSquad().GlobalY = 0
}

func intoToBase(user *player.Player, baseID int) {
	if !user.Bot {
		bases.UserIntoBase(user.GetID(), baseID)
	}

	globalPipe <- Message{Event: "IntoToBase", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot}
	DisconnectUser(user)

	user.InBaseID = baseID

	user.GetSquad().GlobalX = 0
	user.GetSquad().GlobalY = 0
}
