package targetPhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../Phases"
)

func GetWeaponTargetCoordinate(gameUnit *unit.Unit, activeGame *localGame.Game) map[string]map[string]*coordinate.Coordinate {

	targetCoordinate := make(map[string]map[string]*coordinate.Coordinate)

	weaponRange := 0

	for _, weaponSlot := range gameUnit.Body.Weapons {
		if weaponSlot.Weapon != nil {
			weaponRange = weaponSlot.Weapon.Range
		}
	}

	unitCoordinate, find := activeGame.Map.GetCoordinate(gameUnit.GetQ(), gameUnit.GetR())

	if find {
		RadiusCoordinates := coordinate.GetCoordinatesRadius(unitCoordinate, weaponRange)
		zone := filter(gameUnit, RadiusCoordinates, activeGame)

		for _, gameCoordinate := range zone {
			if !(gameCoordinate.X == gameUnit.Q && gameCoordinate.Y == gameUnit.R) {
				Phases.AddCoordinate(targetCoordinate, gameCoordinate)
			}
		}
	}

	return targetCoordinate
}
