package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
	"math"
)

func placeUnit(user *player.Player, msg Message) {
	if user.GetSquad() != nil && user.GetSquad().MatherShip != nil {
		outUnit := user.GetSquad().GetUnitByID(msg.UnitID)

		if outUnit == nil {
			go SendMessage(Message{Event: "Error", Error: "unit not find", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
			return
		}

		if outUnit.OnMap {
			// возвращаем

			outUnit.FollowUnitID = user.GetSquad().MatherShip.ID
			outUnit.Return = true

			Move(
				user,
				Message{
					ToX:     float64(user.GetSquad().MatherShip.X),
					ToY:     float64(user.GetSquad().MatherShip.Y),
					UnitsID: []int{outUnit.ID},
				},
				false,
			)

		} else {

			// берем координату позади отряда todo смотрим что бы она была пустая
			radRotate := float64(user.GetSquad().MatherShip.Rotate) * math.Pi / 180

			stopX := float64(90) * math.Cos(radRotate) // идем по вектору движения корпуса
			stopY := float64(90) * math.Sin(radRotate)

			placeX := float64(user.GetSquad().MatherShip.X) - stopX // - т.к. нам нужна точка позади
			placeY := float64(user.GetSquad().MatherShip.Y) - stopY

			outUnit.X = int(placeX)
			outUnit.Y = int(placeY)
			outUnit.MapID = user.GetSquad().MatherShip.MapID

			if user.GetSquad().MatherShip.Rotate > 180 {
				outUnit.Rotate = user.GetSquad().MatherShip.Rotate - 180
			} else {
				outUnit.Rotate = user.GetSquad().MatherShip.Rotate + 180
			}

			units := globalGame.Clients.GetAllShortUnits(user.GetSquad().MatherShip.MapID, true)
			mp, _ := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)

			okUnits, _ := collisions.CheckCollisionsPlayers(outUnit, outUnit.X, outUnit.Y, outUnit.Rotate, units)
			okMap, _, _, _ := collisions.CheckCollisionsOnStaticMap(outUnit.X, outUnit.Y, outUnit.Rotate, mp, outUnit.Body)

			if okUnits && okMap {

				// выводим
				outUnit.OnMap = true
				outUnit.Afterburner = false
				outUnit.MoveChecker = false
				outUnit.ActualPath = nil
				outUnit.HighGravity = move.GetGravity(outUnit.X, outUnit.Y, outUnit.MapID)

				outUnit.CalculateParams()

				go globalGame.Clients.PlaceUnit(outUnit)
				go update.Squad(user.GetSquad(), true)
				go SendMessage(Message{Event: "PlaceUnit", ShortUnit: outUnit.GetShortInfo(), IDMap: outUnit.MapID})
			} else {
				go SendMessage(Message{Event: "Error", Error: "place is busy", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
			}
		}
	}
}

func ReturnUnit(user *player.Player, moveUnit *unit.Unit) {
	moveUnit.OnMap = false

	go SendMessage(Message{
		Event:     "RemoveUnit",
		ShortUnit: moveUnit.GetShortInfo(),
		IDMap:     moveUnit.MapID,
	})
	go globalGame.Clients.RemoveUnitByID(moveUnit.ID)
	go update.Squad(user.GetSquad(), true)
}
