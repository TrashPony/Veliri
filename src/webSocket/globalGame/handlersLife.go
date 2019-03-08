package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"time"
)

func HandlersLife() {
	allMaps := maps.Maps.GetAllMap()

	for _, mp := range allMaps {
		for _, coor := range mp.HandlersCoordinates {
			if !coor.Transport {
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

		if coor.Handler == "base" {
			continue
		}

		xOut, yOut := globalGame.GetXYCenterHex(coor.ToQ, coor.ToR)
		if checkHandlerCoordinate(xOut, yOut, coor.ToMapID) {
			// отключение телепорта
			globalPipe <- Message{Event: "handlerClose", idMap: mp.Id, Q: coor.Q, R: coor.R}
			coor.HandlerOpen = false
		} else {
			// включение телепорт
			globalPipe <- Message{Event: "handlerOpen", idMap: mp.Id, Q: coor.Q, R: coor.R}
			coor.HandlerOpen = true
		}
	}
}

func checkTransitionUser(x, y, mapID int, coor *coordinate.Coordinate) {
	for _, user := range globalGame.Clients.GetAll() {
		if mapID == user.GetSquad().MapID {
			dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
			if dist < 100 && !user.GetSquad().SoftTransition && coor.HandlerOpen {
				go softTransition(user, x, y, coor)
			}
		}
	}
}

func softTransition(user *player.Player, x, y int, coor *coordinate.Coordinate) {
	countTime := 0
	globalPipe <- Message{Event: "softTransition", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Seconds: 3, Bot: user.Bot}

	user.GetSquad().SoftTransition = true
	defer func() {
		user.GetSquad().SoftTransition = false
	}()

	for {
		time.Sleep(100 * time.Millisecond)
		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
		if dist < 120 && countTime > 50 {
			if coor.Handler == "base" {
				go intoToBase(user, coor.ToBaseID)
			}
			if coor.Handler == "sector" {
				go changeSector(user, coor.ToMapID, coor.ToQ, coor.ToR)
			}
			return
		} else {
			if dist > 120 {
				globalPipe <- Message{Event: "removeSoftTransition", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot}
				return
			}
		}
		countTime++
	}
}

func checkHandlerCoordinate(x, y, mapID int) bool {
	// true занята
	// false свободна
	for _, user := range globalGame.Clients.GetAll() {
		if mapID == user.GetSquad().MapID {
			dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
			if dist < 75 {
				return true
			}
		}
	}
	return false
}
