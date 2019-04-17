package watchZone

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"strconv"
)

type Watcher interface {
	GetQ() int
	GetR() int
	GetY() int
	GetWatchZone() int
	GetOwnerUser() string
	GetWallHack() bool
}

func watch(gameObject Watcher, login string, game *localGame.Game) (allCoordinate map[string]*coordinate.Coordinate, unitsCoordinate map[int]map[int]*unit.Unit, Err error) {

	allCoordinate = make(map[string]*coordinate.Coordinate)
	unitsCoordinate = make(map[int]map[int]*unit.Unit)

	// если игрок является владельцем юнита или состоит в пакте с тем игроком который владеет юнитом
	owner := game.GetUserByName(gameObject.GetOwnerUser())
	client := game.GetUserByName(login)
	// TODO ошибка!
	if owner != nil && client != nil && (login == gameObject.GetOwnerUser() || game.CheckPacts(owner.GetID(), client.GetID())) {

		centerCoordinate, _ := game.Map.GetCoordinate(gameObject.GetQ(), gameObject.GetR())

		RadiusCoordinates := coordinate.GetCoordinatesRadius(centerCoordinate, gameObject.GetWatchZone())
		PermCoordinates := filter(gameObject, RadiusCoordinates, game, gameObject.GetWallHack())

		for _, gameCoordinate := range PermCoordinates {
			unitInMap, ok := game.GetUnit(gameCoordinate.Q, gameCoordinate.R)

			newCoordinate, find := game.Map.GetCoordinate(gameCoordinate.Q, gameCoordinate.R)
			if find {
				allCoordinate[strconv.Itoa(gameCoordinate.Q)+":"+strconv.Itoa(gameCoordinate.R)] = newCoordinate
			}

			if ok {
				if unitsCoordinate[gameCoordinate.Q] != nil {
					unitsCoordinate[gameCoordinate.Q][gameCoordinate.R] = unitInMap
				} else {
					unitsCoordinate[gameCoordinate.Q] = make(map[int]*unit.Unit)
					unitsCoordinate[gameCoordinate.Q][gameCoordinate.R] = unitInMap
				}
			}
		}
	} else {
		return allCoordinate, unitsCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitsCoordinate, nil
}
