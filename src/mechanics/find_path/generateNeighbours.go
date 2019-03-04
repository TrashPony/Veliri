package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

// TODO возможно есть способ это все упаковать в минимальное количества кода т.к. он тут ппц повтяющиеся
func generateNeighboursCoordinate(client *player.Player, curr *coordinate.Coordinate, gameMap *_map.Map, gameUnit *unit.Unit) (res map[string]map[string]*coordinate.Coordinate) {
	/*
	   соседи гексов беруться по разному в зависимости от четности строки
	   // even {Q,R}

	      {0,-1}  {+1,-1}
	   {-1,0} {0,0} {+1,0}
	      {0,+1}  {+1,+1}

	   // odd
	     {-1,-1}  {0,-1}
	   {-1,0} {0,0} {+1,0}
	     {-1,+1}  {0,+1}
	*/

	curr, find := gameMap.GetCoordinate(curr.Q, curr.R)
	if !find {
		return
	}

	res = make(map[string]map[string]*coordinate.Coordinate)

	//left
	checkNeighbour(curr.Q-1, curr.R, client, curr, gameMap, res, gameUnit)
	//right
	checkNeighbour(curr.Q+1, curr.R, client, curr, gameMap, res, gameUnit)

	if curr.R%2 != 0 {
		// topLeft
		checkNeighbour(curr.Q, curr.R-1, client, curr, gameMap, res, gameUnit)
		// topRight
		checkNeighbour(curr.Q+1, curr.R-1, client, curr, gameMap, res, gameUnit)
		// botLeft
		checkNeighbour(curr.Q, curr.R+1, client, curr, gameMap, res, gameUnit)
		// botRight
		checkNeighbour(curr.Q+1, curr.R+1, client, curr, gameMap, res, gameUnit)
	} else {
		// topLeft
		checkNeighbour(curr.Q-1, curr.R-1, client, curr, gameMap, res, gameUnit)
		// topRight
		checkNeighbour(curr.Q, curr.R-1, client, curr, gameMap, res, gameUnit)
		// botLeft
		checkNeighbour(curr.Q-1, curr.R+1, client, curr, gameMap, res, gameUnit)
		// botRight
		checkNeighbour(curr.Q, curr.R+1, client, curr, gameMap, res, gameUnit)
	}

	return
}

func checkNeighbour(q, r int, client *player.Player, curr *coordinate.Coordinate, gameMap *_map.Map, res map[string]map[string]*coordinate.Coordinate, gameUnit *unit.Unit) {
	neighbour, find := checkValidForMoveCoordinate(client, gameMap, q, r)

	x, y := globalGame.GetXYCenterHex(q, r)
	possible, _, _, _ := globalGame.CheckCollisionsOnStaticMap(x, y, gameUnit.Rotate, gameMap, gameUnit.Body)

	if find && possible {
		AddCoordinate(res, neighbour)
	}
}
