package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"time"
)

func HandlerParse(user *player.Player, coor *coordinate.Coordinate) {
	if coor.Handler == "base" {
		IntoToBase(user, coor.ToBaseID)
	}

	if coor.Handler == "sector" {
		ChangeSector(user, coor.ToMapID, coor)
	}

	if !user.Bot {
		go update.Squad(user.GetSquad(), true)
	}
}

func ChangeSector(user *player.Player, mapID int, coor *coordinate.Coordinate) {

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

	stopMove(user.GetSquad().MatherShip, true)

	go SendMessage(Message{Event: "changeSector", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID, Bot: user.Bot})
	DisconnectUser(user, true) // если только сообщение то можно не горутиной

	user.GetSquad().MatherShip.MapID = mapID
	user.GetSquad().MatherShip.X = toPosition.X
	user.GetSquad().MatherShip.Y = toPosition.Y
	user.GetSquad().MatherShip.Rotate = toPosition.RespRotate

	user.GetSquad().MatherShip.X = 0
	user.GetSquad().MatherShip.Y = 0
	user.GetSquad().MatherShip.PointsPath = nil

	if user.Bot {
		LoadGame(user, Message{Event: "InitGame"})
	}
}

func IntoToBase(user *player.Player, baseID int) {
	if !user.Bot {
		bases.UserIntoBase(user.GetID(), baseID)
	}

	go SendMessage(Message{Event: "IntoToBase", IDUserSend: user.GetID(), Bot: user.Bot})
	go DisconnectUser(user, true)

	user.InBaseID = baseID

	if user.GetSquad() != nil {
		user.GetSquad().MatherShip.X = 0
		user.GetSquad().MatherShip.Y = 0
	}
}
