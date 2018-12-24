package globalGame

import (
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
	"../../mechanics/factories/maps"
	"time"
)

func evacuationSquad(ws *websocket.Conn, msg Message, stopMove chan bool, moveChecker *bool, evacuation *bool) {

	mp, find := maps.Maps.GetByID(usersGlobalWs[ws].GetSquad().MapID)
	user := usersGlobalWs[ws]

	if find {
		*evacuation = true

		if *moveChecker {
			stopMove <- true // останавливаем движение
		}

		path, baseID := globalGame.LaunchEvacuation(user, mp)

		for ws := range usersGlobalWs {
			ws.WriteJSON(Message{Event: "startMoveEvacuation", PathUnit: path[0], BaseID: baseID})
		}

		time.Sleep(2 * time.Second) // задержка что бы проиграть анимацию взлета)

		for _, pathUnit := range path {
			for ws := range usersGlobalWs {
				ws.WriteJSON(Message{Event: "MoveEvacuation", PathUnit: pathUnit, BaseID: baseID})
				time.Sleep(100 * time.Millisecond)
			}
		}

		for ws := range usersGlobalWs {
			ws.WriteJSON(Message{Event: "placeEvacuation", OtherUser: GetShortUserInfo(user), BaseID: baseID})
		}

		time.Sleep(2 * time.Second) // задержка что бы проиграть анимацию забора мс

		path = globalGame.ReturnEvacuation(user, mp, baseID)
		for _, pathUnit := range path {
			for ws := range usersGlobalWs {
				// todo подствалять данные эвакуатора мсу
				ws.WriteJSON(Message{Event: "ReturnEvacuation", OtherUser: GetShortUserInfo(user), PathUnit: pathUnit, BaseID: baseID})
				time.Sleep(100 * time.Millisecond)
			}
		}

		for ws := range usersGlobalWs {
			ws.WriteJSON(Message{Event: "stopEvacuation", OtherUser: GetShortUserInfo(user), BaseID: baseID})
		}

		time.Sleep(2 * time.Second) // задержка что бы опустить мс


		// todo intobase
	}
}
