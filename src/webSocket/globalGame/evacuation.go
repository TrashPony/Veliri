package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
	"time"
)

func evacuationSquad(ws *websocket.Conn) {
	user := Clients.GetByWs(ws)

	if user == nil {
		return
	}

	if user.GetSquad().HighGravity {
		globalPipe <- Message{Event: "Error", Error: "High Gravity", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
		return
	}

	mp, find := maps.Maps.GetByID(user.GetSquad().MapID)

	if find && !user.GetSquad().Evacuation && user.InBaseID == 0 {

		user.GetSquad().Evacuation = true

		stopMove(ws, true)

		path, baseID, transport, err := globalGame.LaunchEvacuation(user, mp)
		if err != nil {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
			return
		}

		if len(path) == 0 {
			return
		}

		globalPipe <- Message{Event: "startMoveEvacuation", OtherUser: GetShortUserInfo(user),
			PathUnit: path[0], BaseID: baseID, TransportID: transport.ID, idMap: user.GetSquad().MapID}
		time.Sleep(2 * time.Second) // задержка что бы проиграть анимацию взлета)

		for _, pathUnit := range path {
			globalPipe <- Message{Event: "MoveEvacuation", PathUnit: pathUnit, BaseID: baseID,
				TransportID: transport.ID, idMap: user.GetSquad().MapID}

			transport.X = pathUnit.X
			transport.Y = pathUnit.Y

			time.Sleep(100 * time.Millisecond)
		}

		globalPipe <- Message{Event: "placeEvacuation", OtherUser: GetShortUserInfo(user), BaseID: baseID,
			TransportID: transport.ID, idMap: user.GetSquad().MapID}
		time.Sleep(2 * time.Second) // задержка что бы проиграть анимацию забора мс

		user.GetSquad().InSky = true
		path = globalGame.ReturnEvacuation(user, mp, baseID)

		for _, pathUnit := range path {
			globalPipe <- Message{Event: "ReturnEvacuation", OtherUser: GetShortUserInfo(user), PathUnit: pathUnit,
				BaseID: baseID, TransportID: transport.ID, idMap: user.GetSquad().MapID}

			transport.X = pathUnit.X
			transport.Y = pathUnit.Y
			user.GetSquad().GlobalX = pathUnit.X
			user.GetSquad().GlobalY = pathUnit.Y

			time.Sleep(100 * time.Millisecond)
		}

		globalPipe <- Message{Event: "stopEvacuation", OtherUser: GetShortUserInfo(user), BaseID: baseID,
			TransportID: transport.ID, idMap: user.GetSquad().MapID}
		time.Sleep(1 * time.Second) // задержка что бы опустить мс

		user.InBaseID = baseID
		user.GetSquad().GlobalX = 0
		user.GetSquad().GlobalY = 0

		if !user.Bot {
			globalPipe <- Message{Event: "IntoToBase", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
			go update.Squad(user.GetSquad(), true)
		}

		DisconnectUser(user)

		user.GetSquad().ForceEvacuation = false
		user.GetSquad().Evacuation = false
		user.GetSquad().InSky = false
		transport.Job = false

	}
}
