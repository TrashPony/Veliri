package globalGame

import (
	"../../mechanics/factories/maps"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
	"time"
)

func evacuationSquad(ws *websocket.Conn, msg Message, stopMove chan bool, moveChecker *bool, evacuation *bool) {
	user := usersGlobalWs[ws]

	if user == nil {
		return
	}

	mp, find := maps.Maps.GetByID(user.GetSquad().MapID)

	if find {
		*evacuation = true

		if *moveChecker {
			stopMove <- true // останавливаем движение
		}

		path, baseID, transport, err := globalGame.LaunchEvacuation(user, mp)
		if err != nil {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID()}
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
		user.GetSquad().Evacuation = true

		path = globalGame.ReturnEvacuation(user, mp, baseID)

		for _, pathUnit := range path {
			globalPipe <- Message{Event: "ReturnEvacuation", OtherUser: GetShortUserInfo(user), PathUnit: pathUnit,
				BaseID: baseID, TransportID: transport.ID}

			transport.X = pathUnit.X
			transport.Y = pathUnit.Y

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

		user.GetSquad().Evacuation = false
		*evacuation = false
		transport.Job = false
	}
}
