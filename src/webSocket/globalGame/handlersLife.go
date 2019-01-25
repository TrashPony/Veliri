package globalGame

import (
	"../../mechanics/factories/maps"
	"time"
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/globalGame"
	"../../mechanics/gameObjects/map"
	"../../mechanics/player"
)

func HandlersLife() {
	allMaps := maps.Maps.GetAllMap()

	for _, mp := range allMaps {
		for _, coor := range mp.HandlersCoordinates {
			if coor.Handler == "sector" {
				go entranceMonitor(coor, mp)
			}
		}
	}
}

func entranceMonitor(coor *coordinate.Coordinate, mp *_map.Map) {
	for {
		time.Sleep(100 * time.Millisecond)

		xEntry, yEntry := globalGame.GetXYCenterHex(coor.Q, coor.R)
		checkTransitionUser(xEntry, yEntry, mp.Id, coor)

		xOut, yOut := globalGame.GetXYCenterHex(coor.ToQ, coor.ToR)
		if checkHandlerCoordinate(xOut, yOut, coor.ToMapID) {
			// отключение телепорта
			coor.HandlerOpen = false
		} else {
			// включение телепорт
			coor.HandlerOpen = true
		}
	}
}

func checkTransitionUser(x, y, mapID int, coor *coordinate.Coordinate) {
	for _, user := range Clients.GetAll() {
		if mapID == user.GetSquad().MapID {
			dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
			if dist < 120 && !user.GetSquad().SoftTransition {
				go softTransition(user, x, y, coor)
			}
		}
	}
}

func softTransition(user *player.Player, x, y int, coor *coordinate.Coordinate) {
	countTime := 0
	globalPipe <- Message{Event: "softTransition", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Seconds: 3}

	user.GetSquad().SoftTransition = true
	defer func() {
		user.GetSquad().SoftTransition = false
	}()

	for {
		time.Sleep(100 * time.Millisecond)
		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
		if dist < 120 && countTime > 50 {
			go changeSector(user, coor.ToMapID, coor.ToQ, coor.ToR)
			break
		}
		countTime++
	}
}

func checkHandlerCoordinate(x, y, mapID int) bool {
	// true занята
	// false свободна
	for _, user := range Clients.GetAll() {
		if mapID == user.GetSquad().MapID {
			dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
			if dist < 75 {
				return true
			}
		}
	}
	return false
}
