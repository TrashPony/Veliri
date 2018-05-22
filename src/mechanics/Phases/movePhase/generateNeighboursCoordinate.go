package movePhase

import (
	"../../coordinate"
	"../../gameMap"
	"../../player"
)

func generateNeighboursCoordinate(client *player.Player, curr *coordinate.Coordinate, gameMap *gameMap.Map) (res map[string]map[string]*coordinate.Coordinate) {
	// берет все соседние клетки от текущей
	res = make(map[string]map[string]*coordinate.Coordinate)

	//строго лево
	leftCoordinate, left := checkValidForMoveCoordinate(client, gameMap, curr.X-1, curr.Y)
	if left {
		addCoordinate(res, leftCoordinate)
	}

	//строго право
	rightCoordinate, right := checkValidForMoveCoordinate(client, gameMap, curr.X+1, curr.Y)
	if right {
		addCoordinate(res, rightCoordinate)
	}

	//верх центр
	topCoordinate, top := checkValidForMoveCoordinate(client, gameMap, curr.X, curr.Y-1)
	if top {
		addCoordinate(res, topCoordinate)
	}

	//низ центр
	bottomCoordinate, bottom := checkValidForMoveCoordinate(client, gameMap, curr.X, curr.Y+1)
	if bottom {
		addCoordinate(res, bottomCoordinate)
	}

	//верх лево
	if !(left || top) {
		gameCoordinate, find := checkValidForMoveCoordinate(client, gameMap, curr.X-1, curr.Y-1)
		if find {
			addCoordinate(res, gameCoordinate)
		}
	}

	//верх право
	if !(right || top) {
		gameCoordinate, find := checkValidForMoveCoordinate(client, gameMap, curr.X+1, curr.Y-1)
		if find {
			addCoordinate(res, gameCoordinate)
		}
	}

	//низ лево
	if !(left || bottom) {
		gameCoordinate, find := checkValidForMoveCoordinate(client, gameMap, curr.X-1, curr.Y+1)
		if find {
			addCoordinate(res, gameCoordinate)
		}
	}

	//низ право
	if !(right || bottom) {
		gameCoordinate, find := checkValidForMoveCoordinate(client, gameMap, curr.X+1, curr.Y+1)
		if find {
			addCoordinate(res, gameCoordinate)
		}
	}

	return
}
