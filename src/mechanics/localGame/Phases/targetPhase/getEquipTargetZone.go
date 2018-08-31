package targetPhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/equip"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../player"
)

func GetEquipAllTargetZone(gameUnit *unit.Unit, equip *equip.Equip, activeGame *localGame.Game) map[string]map[string]*coordinate.Coordinate {
	targetCoordinate := make(map[string]map[string]*coordinate.Coordinate)

	/*RadiusCoordinates := coordinate.GetCoordinatesRadius(gameUnit.GetQ(), gameUnit.GetR(), equip.Radius)
	zone := filter(gameUnit, RadiusCoordinates, activeGame)

	for _, gameCoordinate := range zone {
		if !(gameCoordinate.X == gameUnit.Q && gameCoordinate.Y == gameUnit.R) {
			Phases.AddCoordinate(targetCoordinate, gameCoordinate)
		}
	}*/

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
