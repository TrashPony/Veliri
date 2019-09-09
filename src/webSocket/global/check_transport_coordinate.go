package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"time"
)

func OutBase(base *base.Base) *coordinate.Coordinate {

	if base == nil {
		return nil
	}

	base.RespawnLock.Lock()
	defer func() {
		time.Sleep(time.Second)
		base.RespawnLock.Unlock()
	}()

	findRespawn := false
	var respCoordinate *coordinate.Coordinate

	for !findRespawn && respCoordinate == nil {
		findRespawn, respCoordinate = CheckBaseRespawn(base)
		// запускаем механизм проверки и эвакуации игрока с респауна))))
		time.Sleep(time.Millisecond * 100)
	}

	return respCoordinate
}

func CheckBaseRespawn(base *base.Base) (bool, *coordinate.Coordinate) {
	for _, resp := range base.Respawns {
		if !CheckTransportCoordinate(resp.Q, resp.R, 15, 100, base.MapID) {
			return true, resp
		}
	}

	return false, nil
}

func CheckTransportCoordinate(q, r, seconds, distCheck, mapID int) bool { // заставляет игроков эвакуироватся с точки респауна базы

	x, y := globalGame.GetXYCenterHex(q, r)

	lock := false
	units := globalGame.Clients.GetAllShortUnits(mapID, true)

	for _, gameUnit := range units {

		dist := globalGame.GetBetweenDist(gameUnit.X, gameUnit.Y, x, y)

		if int(dist) < distCheck && mapID == gameUnit.MapID {
			dangerUnit := globalGame.Clients.GetUnitByID(gameUnit.ID)
			if !gameUnit.ForceEvacuation {
				user := globalGame.Clients.GetUserByUnitId(gameUnit.ID)

				if user != nil {
					go SendMessage(Message{Event: "setFreeCoordinate", IDUserSend: gameUnit.OwnerID, IDMap: gameUnit.MapID, Seconds: seconds, Bot: user.Bot})
				}

				go ForceEvacuation(dangerUnit, x, y, seconds, distCheck)
			}
			lock = true
			dangerUnit.ForceEvacuation = true
		}
	}

	return lock
}

func ForceEvacuation(unit *unit.Unit, x, y, seconds, distCheck int) {
	timeCount := 0
	for {
		timeCount++
		time.Sleep(100 * time.Millisecond)

		dist := globalGame.GetBetweenDist(unit.X, unit.Y, x, y)

		if int(dist) > distCheck {
			go SendMessage(Message{Event: "removeNoticeFreeCoordinate", IDUserSend: unit.OwnerID, IDMap: unit.MapID})
			unit.ForceEvacuation = false
			break
		} else {
			if timeCount > seconds*10 && !unit.Evacuation {
				go evacuationUnit(unit)
			}
		}
	}
}
