package globalGame

import (
	"../../mechanics/factories/bases"
	"../../mechanics/player"
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/db/squad/update"
)

func HandlerParse(user *player.Player, coor *coordinate.Coordinate) {
	if coor.Handler == "base" {
		intoToBase(user, coor.ToBaseID)
	}

	if coor.Handler == "sector" {
		changeSector(user, coor.ToMapID, coor.ToQ, coor.ToR)
	}

	go update.Squad(user.GetSquad(), true)
}

func changeSector(user *player.Player, mapID, q, r int) {

	user.GetSquad().MapID = mapID

	user.GetSquad().Q = q
	user.GetSquad().R = r

	user.GetSquad().GlobalX = 0
	user.GetSquad().GlobalY = 0

	globalPipe <- Message{Event: "changeSector", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
	DisconnectUser(user)
}

func intoToBase(user *player.Player, baseID int) {
	user.InBaseID = baseID
	bases.UserIntoBase(user.GetID(), baseID)

	user.GetSquad().GlobalX = 0
	user.GetSquad().GlobalY = 0

	globalPipe <- Message{Event: "IntoToBase", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
	DisconnectUser(user)
}
