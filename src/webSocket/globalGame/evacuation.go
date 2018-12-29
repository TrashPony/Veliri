package globalGame

import (
	"../../mechanics/factories/maps"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
	"time"
)

func evacuationSquad(ws *websocket.Conn) {
	user := Clients.GetByWs(ws)

	if user == nil {
		return
	}

	if user.GetSquad().HighGravity {
		globalPipe <- Message{Event: "Error", Error: "High Gravity", idUserSend: user.GetID()}
		return
	}

	mp, find := maps.Maps.GetByID(user.GetSquad().MapID)

	if find && !user.GetSquad().Evacuation && user.InBaseID == 0 {

		user.GetSquad().Evacuation = true

		if user.GetSquad().MoveChecker {
			user.GetSquad().GetMove() <- true // останавливаем движение
		}

		path, baseID, transport, err := globalGame.LaunchEvacuation(user, mp)
		if err != nil {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID()}
			return
		}

		if len(path) == 0 {
			return
		}

		globalPipe <- Message{Event: "startMoveEvacuation", OtherUser: GetShortUserInfo(user),
			PathUnit: path[0], BaseID: baseID, TransportID: transport.ID}
		time.Sleep(2 * time.Second) // задержка что бы проиграть анимацию взлета)

		for _, pathUnit := range path {
			globalPipe <- Message{Event: "MoveEvacuation", PathUnit: pathUnit, BaseID: baseID,
				TransportID: transport.ID}

			transport.X = pathUnit.X
			transport.Y = pathUnit.Y

			time.Sleep(100 * time.Millisecond)
		}

		globalPipe <- Message{Event: "placeEvacuation", OtherUser: GetShortUserInfo(user), BaseID: baseID,
			TransportID: transport.ID}
		time.Sleep(2 * time.Second) // задержка что бы проиграть анимацию забора мс

		user.GetSquad().InSky = true
		path = globalGame.ReturnEvacuation(user, mp, baseID)

		for _, pathUnit := range path {
			globalPipe <- Message{Event: "ReturnEvacuation", OtherUser: GetShortUserInfo(user), PathUnit: pathUnit,
				BaseID: baseID, TransportID: transport.ID}

			transport.X = pathUnit.X
			transport.Y = pathUnit.Y
			user.GetSquad().GlobalX = pathUnit.X
			user.GetSquad().GlobalY = pathUnit.Y

			time.Sleep(100 * time.Millisecond)
		}

		globalPipe <- Message{Event: "stopEvacuation", OtherUser: GetShortUserInfo(user), BaseID: baseID,
			TransportID: transport.ID}
		time.Sleep(1 * time.Second) // задержка что бы опустить мс

		user.InBaseID = baseID
		user.GetSquad().GlobalX = 0
		user.GetSquad().GlobalY = 0

		globalPipe <- Message{Event: "IntoToBase", idUserSend: user.GetID()}

		DisconnectUser(user)

		user.GetSquad().ForceEvacuation = false
		user.GetSquad().Evacuation = false
		user.GetSquad().InSky = false
		transport.Job = false
	}
}
