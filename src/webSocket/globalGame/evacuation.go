package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
	"time"
)

func evacuationSquad(ws *websocket.Conn) {
	user := globalGame.Clients.GetByWs(ws)

	if user == nil {
		return
	}

	if user.GetSquad().HighGravity {
		go SendMessage(Message{Event: "Error", Error: "High Gravity", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
		return
	}

	mp, find := maps.Maps.GetByID(user.GetSquad().MapID)

	if find && !user.GetSquad().Evacuation && user.InBaseID == 0 {

		stopMove(user, true)

		path, baseID, transport, err := globalGame.LaunchEvacuation(user, mp)
		defer func() {
			if user.GetSquad() != nil {
				user.GetSquad().ForceEvacuation = false
				user.GetSquad().Evacuation = false
				user.GetSquad().InSky = false
			}
			transport.Job = false
		}()

		if err != nil {
			go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
			return
		}

		if len(path) == 0 {
			return
		}

		// начали эвакуацию, ставим флаг
		if user.GetSquad() != nil {
			user.GetSquad().Evacuation = true
		} else {
			return
		}

		go SendMessage(Message{Event: "startMoveEvacuation", OtherUser: user.GetShortUserInfo(true, false),
			PathUnit: path[0], BaseID: baseID, TransportID: transport.ID, IDMap: user.GetSquad().MapID})
		time.Sleep(2 * time.Second) // задержка что бы проиграть анимацию взлета)

		for _, pathUnit := range path {

			if user.GetSquad() == nil {
				// игрок умер, больше нечего телепортировать)
				return
			}

			go SendMessage(Message{Event: "MoveEvacuation", PathUnit: pathUnit, BaseID: baseID,
				TransportID: transport.ID, IDMap: user.GetSquad().MapID})

			transport.X = pathUnit.X
			transport.Y = pathUnit.Y

			time.Sleep(100 * time.Millisecond)
		}

		go SendMessage(Message{Event: "placeEvacuation", OtherUser: user.GetShortUserInfo(true, false), BaseID: baseID,
			TransportID: transport.ID, IDMap: user.GetSquad().MapID})
		time.Sleep(2 * time.Second) // задержка что бы проиграть анимацию забора мс

		if user.GetSquad() != nil {
			user.GetSquad().InSky = true
		} else {
			return
		}
		path = globalGame.ReturnEvacuation(user, mp, baseID)

		for _, pathUnit := range path {

			if user.GetSquad() == nil {
				// игрок умер, больше нечего телепортировать)
				return
			}

			go SendMessage(Message{Event: "ReturnEvacuation", OtherUser: user.GetShortUserInfo(true, false), PathUnit: pathUnit,
				BaseID: baseID, TransportID: transport.ID, IDMap: user.GetSquad().MapID})

			transport.X = pathUnit.X
			transport.Y = pathUnit.Y
			user.GetSquad().GlobalX = pathUnit.X
			user.GetSquad().GlobalY = pathUnit.Y

			time.Sleep(100 * time.Millisecond)
		}

		go SendMessage(Message{Event: "stopEvacuation", OtherUser: user.GetShortUserInfo(true, false), BaseID: baseID,
			TransportID: transport.ID, IDMap: user.GetSquad().MapID})
		time.Sleep(1 * time.Second) // задержка что бы опустить мс

		user.InBaseID = baseID

		if user.GetSquad() != nil {
			user.GetSquad().GlobalX = 0
			user.GetSquad().GlobalY = 0
		} else {
			return
		}

		if !user.Bot {
			go SendMessage(Message{Event: "IntoToBase", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
			go update.Squad(user.GetSquad(), true)
			go bases.UserIntoBase(user.GetID(), baseID)
		}

		go DisconnectUser(user, ws, true)
	}
}
