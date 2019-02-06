package targetPhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func GetEquipAllTargetZone(gameUnit *unit.Unit, equip *equip.Equip, activeGame *localGame.Game) map[string]map[string]*coordinate.Coordinate {
	targetCoordinate := make(map[string]map[string]*coordinate.Coordinate)

	unitCoordinate, find := activeGame.Map.GetCoordinate(gameUnit.GetQ(), gameUnit.GetR())

	if find {
		RadiusCoordinates := coordinate.GetCoordinatesRadius(unitCoordinate, equip.Radius)
		zone := filter(gameUnit, RadiusCoordinates, activeGame, false) // еквип не арта поэтому всегда false

		for _, gameCoordinate := range zone {
			if !(gameCoordinate.X == gameUnit.Q && gameCoordinate.Y == gameUnit.R) {
				Phases.AddCoordinate(targetCoordinate, gameCoordinate)
			}
		}
	}

	return targetCoordinate
}

func GetEquipMyUnitsTarget(gameUnit *unit.Unit, equip *equip.Equip, activeGame *localGame.Game, client *player.Player) []*unit.Unit {
	targetZone := GetEquipAllTargetZone(gameUnit, equip, activeGame)
	units := make([]*unit.Unit, 0)
	units = append(units, gameUnit) // кладем того кто использует что бы он мог кинуть на себя

	for _, xLine := range targetZone {
		for _, gameCoordinate := range xLine {
			gameUnit, ok := client.GetUnit(gameCoordinate.X, gameCoordinate.Y)
			if ok {
				units = append(units, gameUnit)
			}
		}
	}

	return units
}

func GetEquipHostileUnitsTarget(gameUnit *unit.Unit, equip *equip.Equip, activeGame *localGame.Game, client *player.Player) []*unit.Unit {
	targetZone := GetEquipAllTargetZone(gameUnit, equip, activeGame)
	units := make([]*unit.Unit, 0)
	for _, xLine := range targetZone {
		for _, gameCoordinate := range xLine {
			gameUnit, ok := client.GetHostileUnit(gameCoordinate.X, gameCoordinate.Y)
			if ok {
				units = append(units, gameUnit)
			}
		}
	}

	return units
}
