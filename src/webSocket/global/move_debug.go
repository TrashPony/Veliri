package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"time"
)

func DebugMoveWorker(user *player.Player) {

	go DebugSender()

	if debug.Store.Collisions {
		go PolygonSender()
	}

	if debug.Store.WeaponFirePos {
		go WeaponFirePos()
	}

	if !debug.Store.MoveInit {
		return
	}

	mp, _ := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)
	for _, zoneX := range mp.GeoZones {
		for _, zone := range zoneX {
			if zone != nil {

				for _, region := range zone.Regions {

					if region == nil {
						continue
					}

					color := ""
					if region.Index == 0 {
						color = "black"
					}

					if region.Index == 1 {
						color = "green"
					}

					if region.Index == 2 {
						color = "orange"
					}

					for _, x := range region.Cells {
						for _, cell := range x {
							debug.Store.AddMessage("CreateRect", color, cell.X, cell.Y, 0, 0, game_math.CellSize, mp.Id, 0)
						}
					}

					//for _, link := range region.Links {
					//	CreateLine(color, link.FromX+game_math.CellSize/2, link.FromY+game_math.CellSize/2,
					//		link.ToX+game_math.CellSize/2, link.ToY+game_math.CellSize/2, game_math.CellSize, mp.Id, user, 0)
					//}
				}
			}
		}
	}

	for _, zoneX := range mp.GeoZones {
		for _, zone := range zoneX {
			if zone != nil {
				debug.Store.AddMessage("CreateRect", "blue",
					zone.DiscreteX*game_math.DiscreteSize, zone.DiscreteY*game_math.DiscreteSize,
					0, 0, zone.Size, mp.Id, 20)
			}
		}
	}
}

func DebugSender() {
	for {

		messages := debug.Store.GetAllMessages()

		for i := 0; i < len(messages); i++ {
			if messages[i].Type == "CreateRect" {
				CreateRect(messages[i].Color, messages[i].X, messages[i].Y, messages[i].Size, messages[i].MapID, messages[i].MS)
			}

			if messages[i].Type == "CreateLine" {
				CreateLine(messages[i].Color, messages[i].X, messages[i].Y, messages[i].ToX, messages[i].ToY, messages[i].Size, messages[i].MapID, messages[i].MS)
			}

			if messages[i].Type == "ClearPath" {

			}
		}

		time.Sleep(20 * time.Millisecond)
	}
}

func CreateRect(color string, startX, startY int, rectSize, mapID int, ms int64) {
	go SendMessage(Message{Event: "CreateRect", Color: color, RectSize: rectSize,
		X: int(startX), Y: int(startY), IDMap: mapID})

	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func WeaponFirePos() {
	for {
		//mps := maps.Maps.GetAllShortInfoMap()
		//
		//for _, mp := range mps {
		//	go SendMessage(Message{Event: "ClearPath", Color: "red", IDMap: mp.Id})
		//}

		users, lock := globalGame.Clients.GetAll()
		for _, user := range users {
			pos := user.GetSquad().MatherShip.GetWeaponFirePos()
			for _, wPos := range pos {
				CreateRect("white", wPos.X, wPos.Y, 2, user.GetSquad().MatherShip.MapID, 0)
			}

			weaponX, weaponY := user.GetSquad().MatherShip.GetWeaponPos()
			CreateRect("red", weaponX, weaponY, 2, user.GetSquad().MatherShip.MapID, 0)

		}

		lock.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func PolygonSender() {
	for {

		mps := maps.Maps.GetAllShortInfoMap()

		for _, mp := range mps {
			go SendMessage(Message{Event: "ClearPolygon", Color: "red", IDMap: mp.Id})
		}

		for _, mp := range mps {

			for _, box := range boxes.Boxes.GetAllBoxByMapID(mp.Id) {

				boxRect := collisions.GetRect(float64(box.X), float64(box.Y), float64(box.Height), float64(box.Width))
				boxRect.Rotate(box.Rotate)

				go SendMessage(Message{Event: "CreatePolygon", Color: "red", IDMap: mp.Id, Polygon: *boxRect})
			}

			for _, unit := range globalGame.Clients.GetAllShortUnits(mp.Id, true) {
				unitRect := collisions.GetBodyRect(unit.Body, float64(unit.X), float64(unit.Y), unit.Rotate, false, false)

				go SendMessage(Message{Event: "CreatePolygon", Color: "red", IDMap: mp.Id, Polygon: *unitRect})
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func ClearVisiblePath(mapID int) {
	go SendMessage(Message{Event: "ClearPath", IDMap: mapID})
}

func CreateLine(color string, X, Y, ToX, ToY int, rectSize, mapID int, ms int64) {
	go SendMessage(Message{Event: "CreateLine", Color: color, RectSize: rectSize,
		X: int(X), Y: int(Y), ToX: float64(ToX), ToY: float64(ToY), IDMap: mapID})

	time.Sleep(time.Duration(ms) * time.Millisecond)
}
