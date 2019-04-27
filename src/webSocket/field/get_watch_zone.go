package field

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/games"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/watchZone"
)

func GetAmmoZone(msg Message, client *player.Player) {
	activeGame, findGame := games.Games.Get(client.GetGameID())

	if findGame && client.GetSquad() != nil {

		gameUnit := client.GetSquad().MatherShip

		allCoordinate, _, err := watchZone.Watch(gameUnit, client.GetLogin(), activeGame)
		if err != nil {
			SendMessage(ErrorMessage{Event: msg.Event, Error: "not found unit"}, client.GetID(), activeGame.Id)
			return
		}

		watch := make(map[string]map[string]*coordinate.Coordinate)
		for _, coor := range allCoordinate {
			Phases.AddCoordinate(watch, coor)
		}

		SendMessage(
			TargetCoordinate{
				Event:   "GetAmmoZone",
				Unit:    gameUnit,
				Targets: watch,
			},
			client.GetID(),
			activeGame.Id,
		)

	} else {
		SendMessage(ErrorMessage{Event: msg.Event, Error: "not found unit"}, client.GetID(), activeGame.Id)
	}
}
