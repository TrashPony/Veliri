package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
	"time"
)

func evacuationUnit(unit *unit.Unit) {

	unit.HighGravity = move.GetGravity(unit.X, unit.Y, unit.MapID)
	user := globalGame.Clients.GetUserByUnitId(unit.ID)

	if unit.HighGravity {
		go SendMessage(Message{Event: "Error", Error: "High Gravity", IDUserSend: unit.OwnerID, IDMap: unit.MapID, Bot: user.Bot})
		return
	}

	mp, find := maps.Maps.GetByID(unit.MapID)

	if find && !unit.Evacuation {

		stopMove(unit, true)

		path, baseID, transport, err := move.LaunchEvacuation(unit, mp)
		defer func() {
			unit.ForceEvacuation = false
			unit.Evacuation = false
			unit.InSky = false
			if transport != nil {
				transport.Job = false
			}
		}()

		if err != nil {
			go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: unit.OwnerID, IDMap: unit.MapID, Bot: user.Bot})
			return
		}

		if len(path) == 0 {
			return
		}

		// начали эвакуацию, ставим флаг
		unit.Evacuation = true

		go SendMessage(Message{Event: "startMoveEvacuation", ShortUnit: unit.GetShortInfo(),
			PathUnit: path[0], BaseID: baseID, TransportID: transport.ID, IDMap: unit.MapID})

		for _, pathUnit := range path {

			if unit.HP <= 0 {
				// игрок умер, больше нечего телепортировать)
				return
			}

			go SendMessage(Message{Event: "MoveEvacuation", PathUnit: pathUnit, BaseID: baseID,
				TransportID: transport.ID, IDMap: unit.MapID})

			transport.X = pathUnit.X
			transport.Y = pathUnit.Y
			transport.Rotate = pathUnit.Rotate

			time.Sleep(100 * time.Millisecond)
		}

		go SendMessage(Message{Event: "placeEvacuation", ShortUnit: unit.GetShortInfo(), BaseID: baseID,
			TransportID: transport.ID, IDMap: unit.MapID})
		time.Sleep(2 * time.Second) // задержка что бы проиграть анимацию забора мс

		if unit.HP > 0 {
			unit.InSky = true
		} else {
			return
		}

		path = move.ReturnEvacuation(unit, mp, baseID, transport)

		for _, pathUnit := range path {

			if unit.HP <= 0 {
				// игрок умер, больше нечего телепортировать)
				return
			}

			go SendMessage(Message{Event: "ReturnEvacuation", ShortUnit: unit.GetShortInfo(), PathUnit: pathUnit,
				BaseID: baseID, TransportID: transport.ID, IDMap: unit.MapID})

			transport.X = pathUnit.X
			transport.Y = pathUnit.Y
			transport.Rotate = pathUnit.Rotate

			unit.X = pathUnit.X
			unit.Y = pathUnit.Y

			time.Sleep(100 * time.Millisecond)
		}

		go SendMessage(Message{Event: "stopEvacuation", ShortUnit: unit.GetShortInfo(), BaseID: baseID,
			TransportID: transport.ID, IDMap: unit.MapID})
		time.Sleep(1 * time.Second) // задержка что бы опустить мс

		user := globalGame.Clients.GetById(unit.OwnerID)
		if unit.Body.MotherShip {
			user.InBaseID = baseID
		}

		if unit.HP > 0 {
			unit.X = 0
			unit.Y = 0
		} else {
			return
		}

		if !user.Bot {
			go SendMessage(Message{Event: "IntoToBase", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
			go update.Squad(user.GetSquad(), true)
			go bases.UserIntoBase(user.GetID(), baseID)
		}

		time.Sleep(1 * time.Second) // искуственная задержка что бы клиент на логин не выбивало
		go DisconnectUser(user, true)
	}
}
